# Implementation Plan: Filters & Dashboard Widgets

## Jelenlegi állapot (problémák, amiket kezelni kell)

### Backend problémák
- `FilterConditionModel` egyetlen `Condition string` mezőt tartalmaz, de a frontendnek és a logikának `Field`, `Op`, `Value` szükséges
- `FilterSerializer` csak a `Name`-et kezeli – `Description` és `Conditions` hiányoznak
- `FilterRepository.All()` és `FindByID()` nem preloadolja a `FilterConditions`-t
- A `FilterService.Create/Update` nem kezeli a condition-öket
- Nincs endpoint az elérhető condition opciókhoz (mezők/operátorok)
- Az issue-k nem szűrhetők filter alapján
- `DashboardModel`-nek nincsenek widget-jei
- Nincs `DashboardWidget` model/repo/service/controller

### Frontend problémák
- `FilterForm.vue` JSON textarea-t használ condition-ökhöz – ember által szerkeszthető UI kell
- `DashboardShowPage.vue` üres – widget grid display szükséges
- Dashboard router csak `/` és `/dashboard/create` – nincs `/dashboard/:dashboardId` route
- `dashboard.store.ts` nem használja a `createCrudActions` mixint
- Nincs widget-hez interface, API, component

---

## Fázisok

### FÁZIS 1: Backend – FilterConditionModel átstrukturálás

**Érintett fájlok:**
- `apps/api/app/domain/model.filter_condition.go`

**Változtatások:**
- A `Condition string` mező helyett 3 mező: `Field string`, `Op string`, `Value string`
- Az `uniqueIndex` eltávolítása vagy frissítése (filter_id + field kombinációra)
- GORM AutoMigrate automatikusan elvégzi a schema változást

```go
type FilterConditionModel struct {
    ID       uuid.UUID `gorm:"type:char(36);primary_key;"`
    FilterID uuid.UUID `gorm:"type:char(36) references filter;not null;"`
    Field    string    `gorm:"type:varchar(50);not null;"`
    Op       string    `gorm:"type:varchar(20);not null;"`
    Value    string    `gorm:"type:varchar(255);not null;"`
    Filter   FilterModel
}
```

---

### FÁZIS 2: Backend – Filter serializer & repository javítás

**Érintett fájlok:**
- `apps/api/app/http/serializer.filter.go`
- `apps/api/app/domain/repository.filter.go`
- `apps/api/app/domain/service.filter.go`

**serializer.filter.go változtatások:**
```go
type FilterConditionRequest struct {
    Field string `json:"field"`
    Op    string `json:"op"`
    Value string `json:"value"`
}

type FilterConditionResponse struct {
    ID    uuid.UUID `json:"id"`
    Field string    `json:"field"`
    Op    string    `json:"op"`
    Value string    `json:"value"`
}

type FilterRequest struct {
    Name        string                   `json:"name"`
    Description string                   `json:"description"`
    Conditions  []FilterConditionRequest `json:"conditions"`
}

type FilterResponse struct {
    ID          uuid.UUID                 `json:"id"`
    Name        string                    `json:"name"`
    Description string                    `json:"description"`
    Conditions  []FilterConditionResponse `json:"conditions"`
}
```

**repository.filter.go változtatások:**
- `All()`: Preload `FilterConditions`
- `FindByID()`: Preload `FilterConditions`
- `Create()`: GORM association-t használ, condition-ök automatikusan létrejönnek
- `Update()`: Régi condition-öket törli, újakat vesz fel (replace strategy)

---

### FÁZIS 3: Backend – Condition Options endpoint

**Érintett fájlok:**
- `apps/api/app/http/controller.filter.go` (új `Options` metódus)
- `apps/api/app/http/router.go` (új route)

**Új endpoint:** `GET /filter/condition-options`

Ez statikus adatot ad vissza (nem DB alapú), az Issue model mezőiből fakad:

```json
{
  "fields": [
    {
      "key": "project_id",
      "label": "Project",
      "valueType": "select",
      "valueEndpoint": "/project",
      "operators": [
        {"key": "eq", "label": "is"},
        {"key": "ne", "label": "is not"}
      ]
    },
    {
      "key": "tracker_id",
      "label": "Tracker",
      "valueType": "select",
      "valueEndpoint": "/tracker",
      "operators": [
        {"key": "eq", "label": "is"},
        {"key": "ne", "label": "is not"}
      ]
    },
    {
      "key": "status_id",
      "label": "Status",
      "valueType": "select",
      "valueEndpoint": "/status",
      "operators": [
        {"key": "eq", "label": "is"},
        {"key": "ne", "label": "is not"}
      ]
    },
    {
      "key": "summary",
      "label": "Summary",
      "valueType": "text",
      "operators": [
        {"key": "contains", "label": "contains"},
        {"key": "eq", "label": "equals"}
      ]
    }
  ]
}
```

**Router:**
```go
router.Get("/filter/condition-options", r.FilterController.Options)
```
> Fontos: ez a route a `{filter_id}` wildcard elé kerül!

---

### FÁZIS 4: Backend – Issue szűrés filter alapján

**Érintett fájlok:**
- `apps/api/app/domain/repository.issue.go`
- `apps/api/app/http/controller.issue.go`

**repository.issue.go:**
- Új `AllByFilter(filterConditions []FilterConditionModel) (*[]IssueModel, error)` metódus
- Dinamikusan épít WHERE clausulát a condition-ök alapján:
  - `field=project_id, op=eq` → `.Where("project_id = ?", value)`
  - `field=status_id, op=ne` → `.Where("status_id != ?", value)`
  - `field=summary, op=contains` → `.Where("summary LIKE ?", "%"+value+"%")`

**controller.issue.go:**
- Az `Index` metódus elfogad opcionális `?filterId=<uuid>` query parametert
- Ha meg van adva, betölti a filtert a condition-öivel, majd `AllByFilter`-t hív

---

### FÁZIS 5: Backend – Dashboard Widget rendszer

**Új fájlok:**
- `apps/api/app/domain/model.dashboard_widget.go`
- `apps/api/app/domain/repository.dashboard_widget.go`
- `apps/api/app/domain/service.dashboard_widget.go`
- `apps/api/app/http/controller.dashboard_widget.go`
- `apps/api/app/http/serializer.dashboard_widget.go`

**model.dashboard_widget.go:**
```go
type DashboardWidgetModel struct {
    ID          uuid.UUID  `gorm:"type:char(36);primary_key;"`
    DashboardID uuid.UUID  `gorm:"type:char(36) references dashboard;not null;"`
    FilterID    *uuid.UUID `gorm:"type:char(36) references filter;"`  // nullable
    Type        string     `gorm:"type:varchar(50);not null;"`        // "filter"
    Title       string     `gorm:"type:varchar(100);not null;"`
    CreatedAt   time.Time

    Dashboard DashboardModel
    Filter    *FilterModel
}
```

**Serializer:**
```go
type DashboardWidgetRequest struct {
    Type     string     `json:"type"`     // "filter"
    Title    string     `json:"title"`
    FilterID *uuid.UUID `json:"filterId"` // nullable
}

type DashboardWidgetResponse struct {
    ID          uuid.UUID        `json:"id"`
    DashboardID uuid.UUID        `json:"dashboardId"`
    Type        string           `json:"type"`
    Title       string           `json:"title"`
    Filter      *FilterResponse  `json:"filter"` // preloaded with conditions
}
```

**Routes:**
```go
router.Get("/dashboard/{dashboard_id}/widget", r.DashboardWidgetController.Index)
router.Post("/dashboard/{dashboard_id}/widget", r.DashboardWidgetController.Create)
router.Delete("/dashboard/{dashboard_id}/widget/{widget_id}", r.DashboardWidgetController.Destroy)
```

**domain.go frissítés:**
- `DashboardWidgetRepository` és `DashboardWidgetService` hozzáadása
- `model.migrate.go`-ban `DashboardWidgetModel{}` hozzáadása

---

### FÁZIS 6: Frontend – FilterConditionBuilder component

**Új fájl:** `apps/frontend/src/component/FilterConditionBuilder.vue`

Ez a komponens felváltja a JSON textarea-t a FilterForm-ban.

**Props:**
```typescript
modelValue: FilterCondition[]  // v-model
```

**Logika:**
- Betölti a condition-options endpointból az elérhető mezőket és operátorokat
- Minden condition sorban: Field dropdown, Op dropdown, Value input (text vagy v-select a field valueType alapján)
- Value input: ha `valueType === 'select'`, hívja a `valueEndpoint`-ot és dropdown-t jelenít meg; ha `text`, egyszerű szövegmező
- Plusz gomb új sor hozzáadásához, mínusz gomb sor törléséhez

**Viselkedés:**
- Az options betöltése `onMounted`-ban
- Select típusú értékek esetén a megfelelő resource-t fetch-eli (project, tracker, status) – ezekhez a meglévő store-okat használja (projectStore, trackerStore, statusStore)

---

### FÁZIS 7: Frontend – FilterForm frissítés

**Érintett fájl:** `apps/frontend/src/component/FilterForm.vue`

**Változtatások:**
- JSON textarea eltávolítása
- `FilterConditionBuilder` beillesztése `v-model="conditions"` bindinggal
- `conditions` ref típusa `FilterCondition[]`

---

### FÁZIS 8: Frontend – Filter API frissítés

**Érintett fájl:** `apps/frontend/src/api/filter.api.ts`

A `createCrudApi`-n túl egy extra metódus szükséges:

```typescript
import api from '@/api/api';
import { createCrudApi } from './baseApi';
import { Filter } from '@/lib/interfaces/filter.interface';

const crudApi = createCrudApi<Filter>('filter');

export default {
    ...crudApi,
    async conditionOptions(): Promise<FilterConditionOptions> {
        const response = await api.apiCall().get('/filter/condition-options');
        return response.data;
    }
};
```

**Új interface (`filter.interface.ts` kiegészítés):**
```typescript
export interface FilterConditionOperator {
    key: string;
    label: string;
}

export interface FilterConditionField {
    key: string;
    label: string;
    valueType: 'text' | 'select';
    valueEndpoint?: string;
    operators: FilterConditionOperator[];
}

export interface FilterConditionOptions {
    fields: FilterConditionField[];
}
```

---

### FÁZIS 9: Frontend – Dashboard interface & store frissítés

**Érintett fájl:** `apps/frontend/src/lib/interfaces/dashboard.interface.ts` (létrehozás)

A widget-ek polimorf típushierarchiával épülnek fel: minden konkrét widget type implementálja a base `DashboardWidget` interfészt. A `type` mező a diszkriminátor.

```typescript
import { Filter } from './filter.interface';

export interface Dashboard {
    id: string;
    name: string;
    description?: string;
}

// Base interface – minden widget type implementálja
export interface DashboardWidget {
    id: string;
    dashboardId: string;
    type: string;
    title: string;
}

// Konkrét filter widget – extends DashboardWidget
export interface DashboardFilterWidget extends DashboardWidget {
    type: 'filter';
    filterId: string;
    filter?: Filter;
}

// Union type – ez kerül a store-ba és a component prop-okba
// Jövőbeli widget type-ok ide kerülnek: | DashboardChartWidget | ...
export type AnyDashboardWidget = DashboardFilterWidget;
```

**Miért ez a struktúra:**
- `DashboardWidget` (base) definiálja a közös mezőket – minden widget compliant
- `DashboardFilterWidget` kiterjeszti a base-t és szűkíti a `type`-ot `'filter'`-re
- A `AnyDashboardWidget` union type a tényleges használati hely, bővítéssel csak ide kell új tagot hozzáadni
- TypeScript type guard segítségével a renderelő logika biztonságosan narrowolja a típust:

```typescript
// type guard
function isFilterWidget(w: AnyDashboardWidget): w is DashboardFilterWidget {
    return w.type === 'filter';
}
```

**Érintett fájl:** `apps/frontend/src/api/dashboardWidget.api.ts` (új)

```typescript
// Nested resource: /dashboard/:dashboardId/widget
// Nem használja a createCrudApi-t mert nested endpoint
import api from '@/api/api';
import { AnyDashboardWidget, DashboardWidget } from '@/lib/interfaces/dashboard.interface';

export default {
    async index(dashboardId: string): Promise<AnyDashboardWidget[]> {
        const r = await api.apiCall().get(`/dashboard/${dashboardId}/widget`);
        return r.data;
    },
    async create(dashboardId: string, widget: Omit<DashboardWidget, 'id' | 'dashboardId'>): Promise<AnyDashboardWidget> {
        const r = await api.apiCall().post(`/dashboard/${dashboardId}/widget`, widget);
        return r.data;
    },
    async destroy(dashboardId: string, widgetId: string): Promise<void> {
        await api.apiCall().delete(`/dashboard/${dashboardId}/widget/${widgetId}`);
    },
};
```

**Érintett fájl:** `apps/frontend/src/store/dashboard.store.ts`

- `createCrudActions` mixin bevezetése (mint a filter.store.ts)
- Az aktuális dashboard widget-jei külön `widgets: AnyDashboardWidget[]` state-ben
- Widget CRUD akciók: `loadWidgets(dashboardId)`, `addWidget(...)`, `removeWidget(...)`

---

### FÁZIS 10: Frontend – Dashboard router frissítés

**Érintett fájl:** `apps/frontend/src/router/dashboard.router.ts`

```typescript
{
    path: '/dashboard/:dashboardId',
    name: 'dashboardShow',
    component: DashboardShowPage,
},
{
    path: '/',
    redirect: () => { /* redirect első dashboardhoz vagy /dashboard/create-hoz */ }
}
```

---

### FÁZIS 11: Frontend – DashboardFilterWidget component

**Új fájl:** `apps/frontend/src/component/DashboardFilterWidget.vue`

Ez a komponens implementálja a `DashboardFilterWidget` interfészt (nem a base-t).

**Props:**
```typescript
widget: DashboardFilterWidget  // a szűkített típus, nem az általános DashboardWidget
```

**Logika:**
- `onMounted`: betölti az issue-kat `?filterId=<widget.filterId>` query paraméterrel
- Megjelenít: widget title, issue lista (key, summary, status chip)
- `@remove` event emit a parentnek (widgetId-vel)

**Jövőbeli widget típus hozzáadásakor** ugyanígy: `DashboardChartWidget.vue` props-a `DashboardChartWidget` típusú lesz.

---

### FÁZIS 12: Frontend – DashboardShowPage implementálás

**Érintett fájl:** `apps/frontend/src/page/dashboard/DashboardShowPage.vue`

**Funkciók:**
1. Route param alapján betölti az aktuális dashboardot és widget-jeit
2. Sidebar-ban felsorolja az összes dashboardot (mint jelenleg)
3. Widget grid megjelenítés (`v-row/v-col` alapú, nem drag-drop - ez majd később)
4. "Add Widget" gomb → dialog nyílik:
   - Widget type kiválasztása (egyelőre csak "Filter" opció)
   - Filter kiválasztása (v-select a filterek listájából)
   - Title mező
   - Mentés gomb
5. Minden widget-en X gomb a törléshez

**Dinamikus komponens renderelés `type` alapján:**

A show page nem hardcolja be a widget komponenseket – egy mapping objecten keresztül dinamikusan renderi őket:

```typescript
import DashboardFilterWidgetComponent from '@/component/DashboardFilterWidget.vue';
// import DashboardChartWidget from '@/component/DashboardChartWidget.vue';  // jövőbeli

const widgetComponents: Record<string, Component> = {
    'filter': DashboardFilterWidgetComponent,
    // 'chart': DashboardChartWidget,
};
```

```vue
<component
    v-for="widget in dashboardStore.widgets"
    :key="widget.id"
    :is="widgetComponents[widget.type]"
    :widget="widget"
    @remove="handleRemoveWidget"
/>
```

Ha egy ismeretlen `type` érkezik a backendről, a `widgetComponents[widget.type]` `undefined` lesz – ezt `v-if` guardolhatja.

**UI mock:**
```
┌─────────────────────────────────────────────────┐
│ [Dashboard neve ▼]              [+ Add Widget]  │
├─────────────────┬───────────────────────────────┤
│ FilterWidget    │ FilterWidget                  │
│ "Open Issues"  ×│ "My Tasks"                  ×│
│ ─────────────── │ ──────────────────────────── │
│ BUG-1 Summary   │ FEAT-3 Summary               │
│ BUG-2 Summary   │ FEAT-7 Summary               │
│ ...             │ ...                           │
└─────────────────┴───────────────────────────────┘
```

---

## i18n kiegészítések (`en.ts`)

```typescript
filter: {
    name: 'Name',
    description: 'Description',
    conditions: 'Conditions',
    add_condition: 'Add Condition',
    field: 'Field',
    operator: 'Operator',
    value: 'Value',
},
board: {
    // ... meglévők
    boards: 'Boards',  // már megvan a BoardShowPage-ben de hiányzik a translations-ből
},
dashboard: {
    add_widget: 'Add Widget',
    widget_type: 'Widget Type',
    select_filter: 'Select Filter',
    widget_title: 'Widget Title',
    no_widgets: 'No widgets yet. Add one with the button above.',
},
```

---

## Összefoglaló – érintett fájlok

### Backend (új fájlok)
- `apps/api/app/domain/model.dashboard_widget.go`
- `apps/api/app/domain/repository.dashboard_widget.go`
- `apps/api/app/domain/service.dashboard_widget.go`
- `apps/api/app/http/controller.dashboard_widget.go`
- `apps/api/app/http/serializer.dashboard_widget.go`

### Backend (módosítások)
- `apps/api/app/domain/model.filter_condition.go` – Field/Op/Value mezők
- `apps/api/app/domain/repository.filter.go` – preload conditions, update strategy
- `apps/api/app/domain/service.filter.go` – nincs lényegi változás
- `apps/api/app/domain/repository.issue.go` – AllByFilter metódus
- `apps/api/app/domain/repository.issue_interface.go` (ha létezik) – interface frissítés
- `apps/api/app/domain/domain.go` – DashboardWidgetRepository/Service
- `apps/api/app/domain/model.migrate.go` – DashboardWidgetModel hozzáadás
- `apps/api/app/http/serializer.filter.go` – description + conditions
- `apps/api/app/http/controller.filter.go` – Options metódus
- `apps/api/app/http/controller.issue.go` – filterId query param
- `apps/api/app/http/router.go` – új route-ok

### Frontend (új fájlok)
- `apps/frontend/src/component/FilterConditionBuilder.vue`
- `apps/frontend/src/component/DashboardFilterWidget.vue`
- `apps/frontend/src/api/dashboardWidget.api.ts`
- `apps/frontend/src/lib/interfaces/dashboard.interface.ts`

### Frontend (módosítások)
- `apps/frontend/src/component/FilterForm.vue` – FilterConditionBuilder bevezetés
- `apps/frontend/src/api/filter.api.ts` – conditionOptions() hozzáadás
- `apps/frontend/src/lib/interfaces/filter.interface.ts` – FilterConditionOptions
- `apps/frontend/src/store/dashboard.store.ts` – crudMixin + widget kezelés
- `apps/frontend/src/page/dashboard/DashboardShowPage.vue` – widget grid + add dialog
- `apps/frontend/src/router/dashboard.router.ts` – `/dashboard/:dashboardId` route
- `apps/frontend/src/i18n/en.ts` – új kulcsok

---

## Implementációs sorrend (függőségek alapján)

1. Backend Fázis 1 (FilterConditionModel)
2. Backend Fázis 2 (Filter serializer/repo)
3. Backend Fázis 3 (Condition options endpoint)
4. Backend Fázis 4 (Issue filter by filterId)
5. Backend Fázis 5 (Dashboard Widget backend)
6. Frontend Fázis 6–8 (FilterConditionBuilder + FilterForm + filter.api.ts)
7. Frontend Fázis 9–10 (Dashboard interface/store + router)
8. Frontend Fázis 11–12 (DashboardFilterWidget + DashboardShowPage)
