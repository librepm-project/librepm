type Handler = (data: unknown) => void

class WebSocketService {
  private socket: WebSocket | null = null
  private token: string | null = null
  private typeHandlers: Map<string, Handler[]> = new Map()
  private channelHandlers: Map<string, Handler[]> = new Map()

  connect(token: string) {
    this.token = token
    const protocol = location.protocol === 'https:' ? 'wss:' : 'ws:'
    const url = `${protocol}//${location.host}/api/ws?token=${encodeURIComponent(token)}`
    this.socket = new WebSocket(url)

    this.socket.onmessage = (event) => {
      const msg = JSON.parse(event.data as string)
      this.typeHandlers.get(msg.type)?.forEach((h) => h(msg.data))
      if (msg.channel) {
        this.channelHandlers.get(msg.channel)?.forEach((h) => h(msg.data))
      }
    }

    this.socket.onclose = () => {
      if (this.token) {
        setTimeout(() => this.connect(this.token!), 3000)
      }
    }
  }

  on(type: string, handler: Handler) {
    if (!this.typeHandlers.has(type)) this.typeHandlers.set(type, [])
    this.typeHandlers.get(type)!.push(handler)
  }

  off(type: string, handler: Handler) {
    const list = this.typeHandlers.get(type)
    if (list) this.typeHandlers.set(type, list.filter((h) => h !== handler))
  }

  subscribe(channel: string, handler: Handler) {
    if (!this.channelHandlers.has(channel)) this.channelHandlers.set(channel, [])
    this.channelHandlers.get(channel)!.push(handler)
    this.socket?.send(JSON.stringify({ type: 'subscribe', channel }))
  }

  unsubscribe(channel: string, handler: Handler) {
    const list = this.channelHandlers.get(channel)
    if (!list) return
    const remaining = list.filter((h) => h !== handler)
    this.channelHandlers.set(channel, remaining)
    if (remaining.length === 0) {
      this.channelHandlers.delete(channel)
      this.socket?.send(JSON.stringify({ type: 'unsubscribe', channel }))
    }
  }

  disconnect() {
    this.token = null
    this.socket?.close()
    this.socket = null
  }
}

export const wsService = new WebSocketService()
