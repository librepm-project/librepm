# Project Structure & Conventions

This document defines the structural rules, naming conventions, and architectural patterns used across the LibrePM codebase. All new code must follow these rules.

---

## Table of Contents

1. [Repository Layout](#repository-layout)
2. [Backend — Go](#backend--go)
   - [Package Structure](#package-structure)
   - [File Naming](#backend-file-naming)
   - [Layers](#backend-layers)
   - [Naming Conventions](#backend-naming-conventions)
   - [DTOs & Serializers](#dtos--serializers)
   - [Routing](#routing)
   - [Dependency Injection](#dependency-injection)
   - [Database & Migrations](#database--migrations)
   - [Seed Data](#seed-data)
3. [Frontend — Vue / TypeScript](#frontend--vue--typescript)
   - [Folder Structure](#folder-structure)
   - [File Naming](#frontend-file-naming)
   - [Layers](#frontend-layers)
   - [Naming Conventions](#frontend-naming-conventions)
   - [Stores](#stores)
   - [API Layer](#api-layer)
   - [Routing](#frontend-routing)
   - [Interfaces](#interfaces)
   - [i18n](#i18n)

---

## Repository Layout

```
/
├── apps/
│   ├── api/          # Go backend
│   └── frontend/     # Vue 3 frontend
├── libs/             # Shared Go libraries (http_utils, jwt_utils, mysql_utils)
├── STRUCTURE.md
└── docker-compose.yml
```

---

## Backend — Go

### Package Structure

The API application has three packages under `apps/api/app/`:

| Package | Path | Purpose |
|---------|------|---------|
| `domain` | `app/domain/` | Business logic: models, repositories, services, domain wiring |
| `http` | `app/http/` | HTTP layer: controllers, serializers, router, server setup |
| `seed` | `app/seed/` | Database seeding: data structures, seed services |

### Backend File Naming

Files use a `{type}.{entity}.go` pattern, lowercase, dot-separated.

| Layer | Pattern | Example |
|-------|---------|---------|
| Model | `model.{entity}.go` | `model.project.go` |
| Repository | `repository.{entity}.go` | `repository.project.go` |
| Service | `service.{entity}.go` | `service.project.go` |
| Controller | `controller.{entity}.go` | `controller.project.go` |
| Serializer | `serializer.{entity}.go` | `serializer.project.go` |
| Seed service | `service.seed.{entity}.go` | `service.seed.project.go` |
| Seed data | `data.{entity}.go` | `data.project.go` |
| Singleton files | exact name | `router.go`, `domain.go`, `http.go` |

Multi-word entities use underscores: `model.project_tracker.go`, `repository.project_tracker_status.go`.

### Backend Layers

#### Model (`domain` package)

Represents a database table. Contains GORM struct tags, relations, and lifecycle hooks.

```go
type ProjectModel struct {
    ID        uuid.UUID  `gorm:"type:char(36);primary_key;"`
    Name      string     `gorm:"type:varchar(100);not null;unique;"`
    // ...
    Issues    []IssueModel `gorm:"foreignKey:ProjectID;"`
}

func (p ProjectModel) TableName() string { return "project" }

func (p *ProjectModel) BeforeCreate(tx *gorm.DB) (err error) {
    p.ID = uuid.New()
    return
}
```

Rules:
- Struct name: `{Entity}Model`
- `TableName()` returns lowercase, underscore-separated: `"project_tracker"`
- `BeforeCreate` assigns UUID
- Relations declared as typed slices or pointers on the struct

#### Repository (`domain` package)

One interface + one struct per entity. Handles all database access.

```go
type ProjectRepositoryInterface interface {
    All() (*[]ProjectModel, error)
    FindByID(project_id uuid.UUID) (*ProjectModel, error)
    FindByName(name string) (*ProjectModel, error)
    Create(project *ProjectModel) error
    Update(project_id uuid.UUID, project *ProjectModel) error
    Destroy(project_id uuid.UUID) error
}

type ProjectRepository struct {
    DB *gorm.DB
}
```

Interface is always defined first, above the struct.

Standard method signatures:

| Method | Signature | Purpose |
|--------|-----------|---------|
| `All` | `() (*[]EntityModel, error)` | Fetch all records |
| `FindByID` | `(id uuid.UUID) (*EntityModel, error)` | Fetch single by primary key |
| `FindBy{Field}` | `(value T) (*EntityModel, error)` | Fetch single by field |
| `Find{Entities}By{Parent}ID` | `(parentID uuid.UUID) (*[]EntityModel, error)` | Fetch related records |
| `Create` | `(model *EntityModel) error` | Insert record |
| `Update` | `(id uuid.UUID, model *EntityModel) error` | Update record |
| `Destroy` | `(id uuid.UUID) error` | Delete record |

Error handling: always assign to the named return, never shadow with `:=`:

```go
// Correct
var err error
if err = query.Find(&items).Error; err != nil {
    slog.Error("...", "error", err)
}
return &items, err

// Wrong — err is shadowed and never assigned
if err := query.Find(&items).Error; err != nil { ... }
return &items, err  // always nil
```

#### Service (`domain` package)

One interface + one struct per entity. Contains business logic and delegates to repositories. Does not interact with HTTP.

```go
type ProjectServiceInterface interface {
    All() (*[]ProjectModel, error)
    Show(project_id uuid.UUID) (*ProjectModel, error)
    Create(project *ProjectModel) error
    Update(project_id uuid.UUID, project *ProjectModel) error
    Destroy(project_id uuid.UUID) error
}

type ProjectService struct {
    ProjectRepository        ProjectRepositoryInterface
    ProjectTrackerRepository ProjectTrackerRepositoryInterface
}
```

Method naming:

| Method | Maps to |
|--------|---------|
| `All()` | List all |
| `Show(id)` | Single item (REST semantics) |
| `Create(model)` | Create |
| `Update(id, model)` | Update |
| `Destroy(id)` | Delete |
| Domain methods | Descriptive: `TrackersByProject(id)`, `ShowIssuePropertyById(id)` |

#### Controller (`http` package)

One interface + one struct per entity. Handles HTTP request parsing and response writing. Delegates all logic to the service.

```go
type ProjectControllerInterface interface {
    Index(w http.ResponseWriter, r *http.Request)
    Show(w http.ResponseWriter, r *http.Request)
    Create(w http.ResponseWriter, r *http.Request)
    Update(w http.ResponseWriter, r *http.Request)
    Destroy(w http.ResponseWriter, r *http.Request)
}

type ProjectController struct {
    ProjectService domain.ProjectServiceInterface
}
```

Standard action names:

| Action | HTTP | Purpose |
|--------|------|---------|
| `Index` | GET `/entity` | List all |
| `Show` | GET `/entity/:id` | Single item |
| `ShowBy{Field}` | GET `/entity/{field}/:value` | Single item by alternate key |
| `Create` | POST `/entity` | Create |
| `Update` | PUT `/entity/:id` | Update |
| `Destroy` | DELETE `/entity/:id` | Delete |
| `Index{SubResource}` | GET `/entity/:id/sub` | List sub-resource |

Helper functions used in all controllers (from `controller_helpers.go`):
- `GetParamUUID(r, "param_name")` — extract UUID route param
- `DecodeJSON(r, &target)` — decode request body
- `RespondJSON(w, status, data)` — write JSON response
- `RespondBadRequest(w)`, `RespondNoContent(w)`, `RespondInternalError(w)`

### Backend Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Struct names | PascalCase + suffix | `ProjectModel`, `ProjectService` |
| Interface names | PascalCase + `Interface` suffix | `ProjectServiceInterface` |
| Method names | PascalCase | `FindByID`, `ModelToResponse` |
| Local variables | snake_case | `project_id`, `project_request` |
| Function parameters | snake_case | `project_id uuid.UUID` |
| Struct fields | PascalCase | `DefaultStatusID`, `CreatedAt` |
| JSON tags | camelCase | `json:"defaultStatusId"` |
| Constants | PascalCase | `EntityTypeIssue` |
| Method receivers | value receiver, single letter | `(s ProjectSerializer)`, `(c ProjectController)` |

**Method names must not repeat the struct name.** A method on `ProjectService` is not named `CreateProject` — it is `Create`.

### DTOs & Serializers

Request and Response types live in the `http` package inside `serializer.{entity}.go`.

```go
type ProjectRequest struct {
    Name             string `json:"name"`
    CodeName         string `json:"codeName"`
    DefaultStatusID  string `json:"defaultStatusId"`
    DefaultTrackerID string `json:"defaultTrackerId"`
}

type ProjectResponse struct {
    ID               uuid.UUID  `json:"id"`
    Name             string     `json:"name"`
    CodeName         string     `json:"codeName"`
    DefaultStatusID  *uuid.UUID `json:"defaultStatusId"`
    DefaultTrackerID *uuid.UUID `json:"defaultTrackerId"`
}

type ProjectSerializer struct{}

func (s ProjectSerializer) RequestToModel(r ProjectRequest) domain.ProjectModel
func (s ProjectSerializer) ModelToResponse(m domain.ProjectModel) ProjectResponse
func (s ProjectSerializer) ModelToResponseMany(ms []domain.ProjectModel) []ProjectResponse
```

Rules:
- `{Entity}Request` — input DTO, unmarshalled from HTTP body
- `{Entity}Response` — output DTO, marshalled to HTTP response
- `{Entity}Serializer` — stateless struct with conversion methods
- Serializer methods: `RequestToModel`, `ModelToResponse`, `ModelToResponseMany`
- The domain layer never imports `http` types
- Pointer fields (`*uuid.UUID`) in Response for nullable values; plain strings in Request

### Routing

All routes are registered in `router.go` on the `Router` struct.

```go
type Router struct {
    ProjectController ProjectControllerInterface
    IssueController   IssueControllerInterface
    // ...
}

func (r Router) Init() *chi.Mux {
    router := chi.NewRouter()

    router.Get("/project", r.ProjectController.Index)
    router.Post("/project", r.ProjectController.Create)
    router.Get("/project/{project_id}", r.ProjectController.Show)
    router.Put("/project/{project_id}", r.ProjectController.Update)
    router.Delete("/project/{project_id}", r.ProjectController.Destroy)

    // Sub-resources
    router.Get("/project/{project_id}/tracker", r.ProjectController.IndexTrackers)

    return router
}
```

URL conventions:
- Lowercase, hyphen-separated path segments: `/issue-property`
- Route params: `{snake_case_id}` — e.g. `{project_id}`, `{widget_id}`
- Nested resources: `/parent/{parent_id}/child`
- Non-CRUD actions as sub-paths: `/issue/{id}/worklog`, `/notification/{id}/read`

HTTP methods map to actions: `GET→Index/Show`, `POST→Create`, `PUT→Update`, `PATCH→partial Update`, `DELETE→Destroy`.

### Dependency Injection

`domain.go` is the composition root. It instantiates all repositories and services and returns a `Domain` struct.

```go
type Domain struct {
    ProjectService   ProjectServiceInterface
    IssueService     IssueServiceInterface
    // ... all services
    ProjectRepository ProjectRepositoryInterface
    // ... all repositories
}

func NewDomain() Domain {
    DB := mysql_utils.Init()
    projectRepository := ProjectRepository{DB: DB}
    // ...
    return Domain{
        ProjectService: ProjectService{
            ProjectRepository: projectRepository,
        },
        // ...
    }
}
```

`http.go` receives the `Domain` and constructs the `Router` with controllers:

```go
func StartHttpServer(d domain.Domain) {
    Router{
        ProjectController: ProjectController{ProjectService: d.ProjectService},
        // ...
    }.Init()
}
```

Controllers receive service interfaces, not concrete types. Services receive repository interfaces. Nothing bypasses a layer.

### Database & Migrations

`model.migrate.go` calls `DB.AutoMigrate` for all models and adds any explicit foreign key constraints. All models must be listed here.

### Seed Data

```
app/seed/
├── data.{entity}.go          # Plain Go structs describing seed records
├── service.seed.{entity}.go  # Methods that write seed data to DB
└── service.seed.go           # Orchestrates all seed calls
```

---

## Frontend — Vue / TypeScript

### Folder Structure

```
apps/frontend/src/
├── api/           # API client functions, one file per entity
├── component/     # Reusable Vue components (flat, no subdirectories)
├── i18n/          # Translation files
├── layout/        # App-level layout components
├── lib/
│   ├── formRule/  # Vuelidate / custom validation rules
│   └── interfaces/ # TypeScript interfaces, one file per entity
├── page/          # Page components, grouped by domain
│   ├── admin/
│   │   ├── board/
│   │   ├── priority/
│   │   ├── project/
│   │   ├── setting/
│   │   ├── status/
│   │   ├── tracker/
│   │   └── user/
│   ├── board/
│   ├── dashboard/
│   ├── filter/
│   ├── issue/
│   └── session/
├── router/        # Vue Router configuration, one file per domain
├── store/         # Pinia stores, one file per entity
│   └── utils/     # Store utilities (crudMixin, entity-store.factory)
├── theme/         # Vuetify theme configuration
└── main.ts
```

### Frontend File Naming

| Type | Pattern | Example |
|------|---------|---------|
| API | `{entity}.api.ts` | `project.api.ts`, `issue-audit-log.api.ts` |
| Store | `{entity}.store.ts` | `project.store.ts`, `related-issue.store.ts` |
| Interface | `{entity}.interface.ts` | `project.interface.ts`, `issue-audit-log.interface.ts` |
| Component | `{PascalCase}.vue` | `ProjectForm.vue`, `IssueList.vue` |
| Page | `{PascalCase}Page.vue` | `AdminProjectShowPage.vue`, `IssueIndexPage.vue` |
| Router | `{section}.router.ts` | `board.router.ts`, `admin.project.router.ts` |

Multi-word entities use kebab-case in file names: `issue-audit-log.api.ts`, `related-issue.store.ts`.

### Frontend Layers

```
Page
  └── Store (reads/writes state, calls API)
        └── API (HTTP calls)
              └── Backend
  └── Component (receives props, emits / calls callbacks)
```

Pages orchestrate stores and pass data to components. Components are stateless where possible. Stores own all shared state.

### Frontend Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Interface | PascalCase | `Project`, `Issue`, `FilterCondition` |
| Store export | `use{Entity}Store` | `useProjectStore`, `useIssueStore` |
| Component | PascalCase + type suffix | `ProjectForm`, `IssueList`, `StatusChip` |
| Page | PascalCase + `Page` suffix | `AdminProjectShowPage`, `IssueIndexPage` |
| Router export | camelCase + `Router` | `boardRouter`, `adminProjectRouter` |
| Route name | camelCase | `adminProjectIndex`, `boardShow` |
| URL path | lowercase, kebab-case | `/admin/project`, `/board/:boardId` |
| URL params | camelCase | `:projectId`, `:boardId` |
| Store state | `current` (single) / `index` (list) | `current: Project \| null`, `index: Project[]` |
| Props (callbacks) | `on{Action}` | `onSubmit`, `onDelete` |

**Method names must not repeat the store/struct name.** A method on `useProjectStore` is not named `getProject` — it is `get`. A method on `useSettingStore` is not named `fetchSettings` — it is `fetch`.

### Stores

Every entity has one Pinia store: `{entity}.store.ts`.

```typescript
// project.store.ts
export const useProjectStore = defineStore('project', {
    state: (): ProjectStore => ({
        current: null,
        index: [],
    }),
    actions: {
        ...createCrudActions<Project>(projectApi),

        async getIssueProperty(projectId: string): Promise<void> {
            this.currentIssueProperty = await projectIssuePropertyApi.index(projectId);
        },
    },
});
```

**State shape:**
- `current: Entity | null` — the currently selected/loaded entity
- `index: Entity[]` — the full list of entities
- Additional properties for domain-specific state (e.g. `currentIssueProperty`)

**`createCrudActions<T>(api)` provides:**

| Method | Description |
|--------|-------------|
| `get(id)` | Fetch single item, sets `this.current` |
| `getAll()` | Fetch all items, sets `this.index` |
| `post(item)` | Create item, pushes to `this.index` |
| `put(id, item)` | Update item, replaces in `this.index` |
| `delete(id)` | Delete item, removes from `this.index` |

For stores that cannot use `createCrudActions` (e.g. nested resources, special state), define methods directly — but follow the same naming: `get`, `getAll`, `post`, `put`, `delete` for CRUD, and semantic names (without the entity noun) for domain actions.

### API Layer

One file per entity: `{entity}.api.ts`. For simple CRUD entities, use `createCrudApi`:

```typescript
// tracker.api.ts
export default createCrudApi<Tracker>('tracker');
```

`createCrudApi<T>(endpoint)` provides: `index()`, `show(id)`, `create(item)`, `update(id, item)`, `destroy(id)`.

For entities with extra endpoints, spread the CRUD API and add custom methods:

```typescript
// project.api.ts
const crudApi = createCrudApi<Project>('project');

const trackers = async (projectId: string): Promise<Tracker[]> => {
    const response = await api.apiCall().get(`/project/${projectId}/tracker`);
    return response.data as Tracker[];
};

export default { ...crudApi, trackers };
```

For nested resources (e.g. worklogs under an issue), all methods receive the parent ID as first parameter:

```typescript
const index = async (issueId: string): Promise<Worklog[]> => { ... };
const create = async (issueId: string, payload: Partial<Worklog>): Promise<Worklog> => { ... };
```

API method naming follows the same pattern as backend controller actions: `index`, `show`, `create`, `update`, `destroy`, plus semantic names for custom endpoints.

### Frontend Routing

One router file per domain section: `{section}.router.ts`. Each file exports a `RouteRecordRaw[]` array.

```typescript
// admin.project.router.ts
export const adminProjectRouter: RouteRecordRaw[] = [
    {
        path: '/admin/project',
        name: 'adminProjectIndex',
        component: AdminProjectIndexPage,
    },
    {
        path: '/admin/project/create',
        name: 'adminProjectCreate',
        component: AdminProjectCreatePage,
    },
    {
        path: '/admin/project/:projectId',
        name: 'adminProjectShow',
        component: AdminProjectShowPage,
    },
];
```

All routers are combined in the root `index.router.ts`.

Page naming maps to route names:

| Page type | Route name | Path |
|-----------|------------|------|
| List | `{section}Index` | `/{section}` |
| Create | `{section}Create` | `/{section}/create` |
| Show/Edit | `{section}Show` | `/{section}/:entityId` |

### Interfaces

All TypeScript interfaces live in `lib/interfaces/{entity}.interface.ts`.

```typescript
// project.interface.ts
export interface Project {
    id?: string;
    name: string;
    codeName: string;
    defaultStatusId?: string | null;
    defaultTrackerId?: string | null;
}
```

Conventions:
- `id?: string` — always optional (not present on unsaved objects)
- Foreign keys: `{entity}Id?: string` — optional string
- Related objects: `{entity}?: Entity | null` — optional, explicitly nullable when the API may return null
- Multiple related types in one file when tightly coupled (e.g. `FilterCondition`, `Filter`, `FilterConditionOperator` in `filter.interface.ts`)

### i18n

Translations live in `i18n/en.ts` as a nested object.

```typescript
export const en = {
    global: {
        name: 'Name',
        create: 'Create',
        delete_confirm: 'Are you sure?',
    },
    project: {
        codeName: 'Code Name',
        default_tracker: 'Default Tracker',
    },
    issue: {
        summary: 'Summary',
        direction: {
            blocks: 'Blocks',
            blocked_by: 'Blocked By',
        },
    },
}
```

Key conventions:
- Top-level keys are domain names (`global`, `project`, `issue`) or page section (`title`)
- Second-level keys: camelCase for single words, snake_case for multi-word phrases
- Accessed in templates: `$t('project.default_tracker')`
- `global.*` keys are for labels reused across multiple domains
