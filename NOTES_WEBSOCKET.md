# WebSocket – Terv

## Cél

- Az API WebSocket szerverként is működjön (`GET /ws` endpoint)
- A frontend az alkalmazás betöltésekor csatlakozzon
- **Kétféle push mechanizmus:**
  1. **User-scoped push**: notification érkezésekor az érintett userek kapcsolataira üzenet megy
  2. **Channel-based broadcast**: entity update (pl. issue módosítás) esetén a csatornára feliratkozott kapcsolatokra megy üzenet
- A kapcsolatokat in-memory tároljuk

---

## Backend

### 1. Dependency

```
go get github.com/gorilla/websocket
```

---

### 2. Client struktúra

A gorilla/websocket library nem thread-safe: egy conn-ra egyszerre csak egy writer lehet. Ezért minden WS kapcsolatot egy `Client` struktúrába csomagolunk, ami saját `send` csatornán keresztül fogad kiküldendő üzeneteket, és egy dedikált `writePump` goroutine-ban írja ki azokat.

**`apps/api/app/http/hub.go`**

```go
type Client struct {
    userID uuid.UUID
    conn   *websocket.Conn
    send   chan []byte       // buffered, kimenő üzenetek
    hub    *Hub
}

// writePump: saját goroutine-ban fut, hub.send csatornáról olvas és kiír
func (c *Client) writePump() {
    defer c.conn.Close()
    for msg := range c.send {
        c.conn.WriteMessage(websocket.TextMessage, msg)
    }
}

// readPump: a Connect handler-ben fut (blokkoló), subscribe/unsubscribe üzeneteket kezel,
// disconnect esetén cleanup
func (c *Client) readPump() {
    defer func() {
        c.hub.Unregister(c)
        close(c.send)
    }()
    for {
        _, msg, err := c.conn.ReadMessage()
        if err != nil { return }
        // inbound üzenet feldolgozása (subscribe/unsubscribe)
        c.hub.HandleClientMessage(c, msg)
    }
}
```

---

### 3. Hub

```go
type Hub struct {
    mu       sync.RWMutex
    users    map[uuid.UUID][]*Client   // userID → kapcsolatok
    channels map[string][]*Client      // "issue:uuid" → kapcsolatok
}

func NewHub() *Hub

// User-scoped push (notification)
func (h *Hub) PushToUser(userID uuid.UUID, msg []byte)

// Channel broadcast (entity update)
func (h *Hub) Subscribe(channel string, client *Client)
func (h *Hub) Unsubscribe(channel string, client *Client)
func (h *Hub) Broadcast(channel string, msg []byte)

// Belső lifecycle
func (h *Hub) Register(client *Client)
func (h *Hub) Unregister(client *Client)  // eltávolítja users-ből és minden channel-ből
func (h *Hub) HandleClientMessage(client *Client, raw []byte)
```

`HandleClientMessage` parse-olja az inbound JSON-t:
```json
{"type": "subscribe",   "channel": "issue:550e8400-..."}
{"type": "unsubscribe", "channel": "issue:550e8400-..."}
```

---

### 4. Üzenet formátumok

#### Server → Client: notification push
```json
{
  "type": "notification",
  "data": {
    "id": "...",
    "type": "issue_created",
    "entityType": "issue",
    "entityId": "...",
    "read": false,
    "createdAt": "2026-03-27T12:00:00Z"
  }
}
```

#### Server → Client: entity update broadcast
```json
{
  "type": "entity_update",
  "channel": "issue:550e8400-...",
  "entity": "issue",
  "id": "550e8400-...",
  "data": {
    /* teljes issue serializer output */
  }
}
```

#### Client → Server: subscription management
```json
{"type": "subscribe",   "channel": "issue:550e8400-..."}
{"type": "unsubscribe", "channel": "issue:550e8400-..."}
```

A `channel` formátuma: `"<entity>:<uuid>"` – ez ad alapot bővítésre (pl. `"project:uuid"`, `"comment:uuid"`).

---

### 5. WS Controller

**`apps/api/app/http/controller.ws.go`**

```go
func (c *WebSocketController) Connect(w http.ResponseWriter, r *http.Request) {
    // 1. HTTP → WS upgrade (CheckOrigin: fejlesztésben true)
    conn, _ := c.upgrader.Upgrade(w, r, nil)

    // 2. JWT kiolvasása query paraméterből (?token=...)
    userID := extractUserIDFromToken(r.URL.Query().Get("token"))

    // 3. Client létrehozása, regisztrálás
    client := &Client{userID: userID, conn: conn, send: make(chan []byte, 64), hub: c.Hub}
    c.Hub.Register(client)

    // 4. writePump saját goroutine-ban indul
    go client.writePump()

    // 5. readPump blokkolja a handler-t (disconnect-ig fut)
    client.readPump()
}
```

---

### 6. Notification push bekötése (domain interface)

**`apps/api/app/domain/service.notification.go`** – interface definíció:

```go
type NotificationPusher interface {
    PushToUser(userID uuid.UUID, notification *NotificationModel)
}
```

`NotificationService` kap `Pusher NotificationPusher` mezőt. `Create()` végén:
```go
if s.Pusher != nil {
    s.Pusher.PushToUser(notification.UserID, notification)
}
```

Hub implementálja ezt:
```go
func (h *Hub) PushToUser(userID uuid.UUID, n *domain.NotificationModel) {
    msg, _ := json.Marshal(WsMessage{Type: "notification", Data: serializeNotification(n)})
    h.pushToUser(userID, msg)
}
```

---

### 7. Entity update broadcast bekötése

Az entity broadcast HTTP-réteg logika (nem domain), ezért **a controllerből** hívjuk, nem service-ből.

**`apps/api/app/http/controller.issue.go`** – `IssueController` kap `Hub *Hub` mezőt:

```go
// Update() végén, sikeres mentés után:
channel := "issue:" + issueID.String()
msg, _ := json.Marshal(WsMessage{
    Type:    "entity_update",
    Channel: channel,
    Entity:  "issue",
    ID:      issueID.String(),
    Data:    serializeIssue(updated),
})
c.Hub.Broadcast(channel, msg)
```

Ugyanezt lehet alkalmazni más entity-kre (comment, project, stb.) — a controller mindig tudja, melyik channel-re kell broadcastolni.

---

### 8. Bekötés (`http.go`)

```go
hub := NewHub()

// Notification service: user-scoped push
domain.NotificationService.Pusher = hub

// Controllers: Hub injektálás entity broadcast-hoz
issueController := &IssueController{..., Hub: hub}

// WS controller
wsController := &WebSocketController{Hub: hub, upgrader: ...}
```

**`router.go`**:
```go
router.Get("/ws", r.WebSocketController.Connect)
```

Auth middleware: accept `token` query param is (csak `/ws` route-on).

---

## Frontend

### 1. WebSocket service (bővített)

**`apps/frontend/src/lib/websocket.ts`**

```ts
type Handler = (data: unknown) => void

class WebSocketService {
  private socket: WebSocket | null = null
  private token: string | null = null
  private typeHandlers: Map<string, Handler[]> = new Map()
  private channelHandlers: Map<string, Handler[]> = new Map()

  connect(token: string) {
    this.token = token
    const url = `ws://${location.host}/api/ws?token=${token}`
    this.socket = new WebSocket(url)

    this.socket.onmessage = (event) => {
      const msg = JSON.parse(event.data)
      // type-alapú routing (pl. 'notification')
      this.typeHandlers.get(msg.type)?.forEach(h => h(msg.data))
      // channel-alapú routing (pl. 'entity_update' → 'issue:uuid')
      if (msg.channel) {
        this.channelHandlers.get(msg.channel)?.forEach(h => h(msg.data))
      }
    }

    this.socket.onclose = () => {
      setTimeout(() => { if (this.token) this.connect(this.token) }, 3000)
    }
  }

  // Type-alapú feliratkozás (App.vue szintű, hosszú életű)
  on(type: string, handler: Handler) {
    if (!this.typeHandlers.has(type)) this.typeHandlers.set(type, [])
    this.typeHandlers.get(type)!.push(handler)
  }

  off(type: string, handler: Handler) {
    const list = this.typeHandlers.get(type)
    if (list) this.typeHandlers.set(type, list.filter(h => h !== handler))
  }

  // Channel-alapú subscription (komponens lifecycle-hoz kötve)
  subscribe(channel: string, handler: Handler) {
    if (!this.channelHandlers.has(channel)) this.channelHandlers.set(channel, [])
    this.channelHandlers.get(channel)!.push(handler)
    this.socket?.send(JSON.stringify({ type: 'subscribe', channel }))
  }

  unsubscribe(channel: string, handler: Handler) {
    const list = this.channelHandlers.get(channel)
    if (list) {
      const remaining = list.filter(h => h !== handler)
      this.channelHandlers.set(channel, remaining)
      // Csak akkor küldjük az unsubscribe-t, ha nincs több handler erre a channel-re
      if (remaining.length === 0) {
        this.channelHandlers.delete(channel)
        this.socket?.send(JSON.stringify({ type: 'unsubscribe', channel }))
      }
    }
  }

  disconnect() {
    this.token = null
    this.socket?.close()
    this.socket = null
  }
}

export const wsService = new WebSocketService()
```

---

### 2. App.vue – csatlakozás + notification binding

```ts
onMounted(() => {
  const auth = useAuthStore()
  const notifStore = useNotificationStore()

  if (auth.token) {
    wsService.connect(auth.token)
    wsService.on('notification', (data) => notifStore.addFromWs(data as Notification))
  }
})

onUnmounted(() => wsService.disconnect())
```

---

### 3. IssueShowPage – issue live update

**`apps/frontend/src/page/IssueShowPage.vue`**

```ts
const channel = computed(() => `issue:${props.issueId}`)

const onIssueUpdate = (data: unknown) => {
  issue.value = data as Issue  // frissíti a lokális issue state-et
}

onMounted(() => {
  wsService.subscribe(channel.value, onIssueUpdate)
})

onUnmounted(() => {
  wsService.unsubscribe(channel.value, onIssueUpdate)
})
```

A handler referenciát meg kell tartani (nem inline lambda), hogy az `unsubscribe` meg tudja találni.

---

### 4. Notification store módosítás

```ts
addFromWs(notification: Notification) {
  this.items.unshift(notification)
}
```

---

## Összefoglaló – érintett fájlok

### Backend – új fájlok
| Fájl | Leírás |
|------|--------|
| `apps/api/app/http/hub.go` | `Client` struktúra, Hub (users + channels map) |
| `apps/api/app/http/controller.ws.go` | `/ws` upgrade, readPump/writePump indítás |

### Backend – módosított fájlok
| Fájl | Módosítás |
|------|-----------|
| `apps/api/app/domain/service.notification.go` | `NotificationPusher` interface + `Pusher` mező |
| `apps/api/app/http/http.go` | Hub init, injektálás notification service-be és issue controllerbe |
| `apps/api/app/http/router.go` | `GET /ws` route + token query param auth |
| `apps/api/app/http/controller.issue.go` | `Hub` mező + `Broadcast` hívás update után |

### Frontend – új fájlok
| Fájl | Leírás |
|------|--------|
| `apps/frontend/src/lib/websocket.ts` | WS service (type + channel routing, reconnect) |

### Frontend – módosított fájlok
| Fájl | Módosítás |
|------|-----------|
| `apps/frontend/src/App.vue` | connect + notification binding |
| `apps/frontend/src/store/notification.store.ts` | `addFromWs()` |
| `apps/frontend/src/page/IssueShowPage.vue` | subscribe/unsubscribe issue channel |

---

## Bővíthetőség

Új entity live update hozzáadásához:
1. **Backend**: az entity controller kap `Hub` mezőt, update/create után `hub.Broadcast("entity:id", msg)`
2. **Frontend**: az entity page `wsService.subscribe("entity:id", handler)` / `unsubscribe` on/unmount

A channel naming convention (`"<entity>:<uuid>"`) és az üzenet formátum egységes marad.

---

## Nem kezelt edge case-ek (scope-on kívül)

- Reconnect után a pending subscription-ök újraküldése (socket reconnect esetén a subscribe üzenetek elvesznek – egyszerű megoldás: `connect()` után újraküldi az összes aktív channel subscribe-ot)
- Több tab: mindkét kapcsolaton megérkezik az üzenet, store-ban ID alapján deduplication ha szükséges
- Token lejárat: szerver lezárja, frontend reconnect próbálkozik
