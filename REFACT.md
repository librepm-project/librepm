# LibrePM Refactoring Opportunities

## Overview
Analysis of today's commits (Mar 26, 2026) reveals several refactoring patterns that can improve code maintainability, reduce duplication, and enhance consistency across the codebase.

---

## Backend (Go) - HTTP Layer

### 1. **Error Handling & Response Patterns**
**Issue:** Repeated error handling boilerplate across all controllers.

**Current Pattern:**
```go
param, err := http_utils.GetParamUUID(r, "param_id")
if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
}
```

**Recommendation:**
- Extract into a middleware/helper that automatically handles parameter parsing and responds with consistent error format
- Create a `ControllerBase` with common methods like `GetParamUUID()`, `GetUserID()`, `RespondError()`

**Impact:** Reduces ~200 lines of boilerplate across 20+ controllers.

---

### 2. **Entity-Generic Methods (Attachment, Comment)**
**Issue:** `AllByEntity()` methods in AttachmentService and CommentService use string-based entity types ("issue", "project").

**Current Pattern:**
```go
attachments, err := c.AttachmentService.AllByEntity("issue", issueID)
```

**Recommendation:**
- Create an `EntityType` enum/const to replace magic strings
- Add type-safe wrappers: `AllByIssue()`, `AllByProject()` alongside generic method
- Example:
  ```go
  type EntityType string
  const (
    EntityIssue EntityType = "issue"
    EntityProject EntityType = "project"
  )
  ```

**Impact:** Prevents typos, improves IDE autocomplete, adds semantic meaning.

---

### 3. **Audit Log Dependency Injection Pattern**
**Issue:** Multiple controllers (Attachment, Worklog) manually inject and call `IssueAuditLogService`.

**Current Pattern:**
```go
type AttachmentController struct {
    AttachmentService    domain.AttachmentServiceInterface
    IssueAuditLogService domain.IssueAuditLogServiceInterface
}
```

**Recommendation:**
- Encapsulate audit logging logic inside service methods where data changes occur
- Move `LogAttachmentAdded()` logic into `AttachmentService.Create()`
- Controllers should NOT directly manage audit logs
- Benefits: Single responsibility, audit guarantee, easier testing

**Impact:** Simplifies controller logic by ~30%, ensures audit trails are never missed.

---

### 4. **Request Validation Consolidation**
**Issue:** JSON decoding happens inline in every controller without validation.

**Current Pattern:**
```go
var req CommentRequest
json.NewDecoder(r.Body).Decode(&req)
```

**Recommendation:**
- Create a `DecodeAndValidateJSON()` helper with built-in validation tags
- Use `encoding/json` struct tags for validation rules
- Return detailed validation errors instead of generic 400s

**Impact:** Better error feedback to clients, consistent validation across all endpoints.

---

## Backend (Go) - Domain Layer

### 1. **IssueAuditLogService Responsibilities**
**Issue:** Service has 9+ methods with similar patterns (`LogAttachmentAdded`, `LogWorklogAdded`, `LogRelatedIssueAdded`, etc.).

**Current Pattern:**
```go
func (s IssueAuditLogService) LogAttachmentAdded(userID uuid.UUID, attachment AttachmentModel) error
func (s IssueAuditLogService) LogWorklogAdded(userID uuid.UUID, worklog IssueWorklogModel) error
func (s IssueAuditLogService) LogWorklogUpdated(userID uuid.UUID, old, new IssueWorklogModel) error
```

**Recommendation:**
- Create a generic `LogAction()` method with action type and metadata
- Use a struct-based approach for different action types
- Example:
  ```go
  type AuditAction struct {
    Action string
    Entity string
    Metadata map[string]interface{}
  }
  func (s IssueAuditLogService) Log(userID, issueID uuid.UUID, action AuditAction) error
  ```

**Impact:** Reduces ~100 lines of repetitive code, easier to add new audit event types.

---

### 2. **Repository Pattern Inconsistency**
**Issue:** Multiple repositories with nearly identical CRUD patterns (Attachment, Comment, Worklog).

**Recommendation:**
- Create a `BaseRepository<T>` generic with common CRUD operations
- Inherit from it in specific repositories (only override custom queries)
- Reduces ~150 lines of duplicated repository code

**Example Structure:**
```go
type BaseRepository[T any] struct {
    db *gorm.DB
}
func (r BaseRepository[T]) FindByID(id uuid.UUID) (*T, error) { ... }
func (r BaseRepository[T]) Create(entity *T) error { ... }
func (r BaseRepository[T]) Update(id uuid.UUID, updates *T) error { ... }
func (r BaseRepository[T]) Destroy(id uuid.UUID) error { ... }
```

---

## Frontend (Vue 3/TypeScript) - Stores

### 1. **Store Pattern Duplication**
**Issue:** `attachment.store.ts`, `comment.store.ts`, `worklog.store.ts` follow identical patterns.

**Current Pattern:**
```typescript
export const useAttachmentStore = defineStore('attachment', {
  state: () => ({ attachments: [] }),
  actions: {
    async getAttachments(issueId: string) { ... },
    async create(issueId: string, file: File) { ... },
    async destroy(attachmentId: string) { ... },
  },
});
```

**Recommendation:**
- Create a factory function `createEntityStore()` to generate store boilerplate
- Example:
  ```typescript
  const useAttachmentStore = createEntityStore('attachment', attachmentApi);
  const useCommentStore = createEntityStore('comment', commentApi);
  ```

**Impact:** Eliminates ~80 lines of store code, standardizes state management patterns.

---

### 2. **API Client Duplications**
**Issue:** Similar API patterns across `attachment.api.ts`, `comment.api.ts`, `worklog.api.ts`.

**Recommendation:**
- Create a generic `createEntityAPI()` factory for standard CRUD operations
- Define only custom endpoints per entity
- All three would reduce to ~10 lines each

---

## General Architecture

### 1. **Circular Audit Log Dependency**
**Issue:** Controllers depend on both service AND audit service. Services don't encapsulate audit logic.

**Recommendation:**
- Implement the Audit Log pattern inside domain services (service → repository → audit repository)
- Controllers call service, service handles audit as a side effect
- Cleaner dependency graph: Controller → Service → (Repository + Audit)

---

### 2. **Missing Error Categorization**
**Issue:** All errors respond with `StatusBadRequest` or `StatusInternalServerError` generically.

**Recommendation:**
- Create domain error types:
  ```go
  type DomainError struct {
    Type string
    Message string
    StatusCode int
  }
  // e.g., NotFoundError, ValidationError, ConflictError
  ```
- Map errors to HTTP status codes systematically

---

## Testing Implications

### 1. **Testability Improvements**
- Generic repository base will simplify mock creation
- Audit log consolidation reduces integration test complexity
- Store factory approach enables parameterized store tests

### 2. **Suggested Test Coverage**
- Unit test generic repository with type variations
- Test audit logging in service layer (not controller layer)
- Verify error mappings with error categorization

---

## Implementation Priority

### Phase 1 (High Impact, Low Risk)
1. ✅ Create error response helpers (HTTP layer)
2. ✅ Consolidate JSON decode/validate logic
3. ✅ Extract entity type constants

### Phase 2 (Medium Impact, Medium Effort)
1. Consolidate IssueAuditLogService methods
2. Create generic store factory (frontend)
3. Create generic API factory (frontend)

### Phase 3 (Architectural Refactoring)
1. Move audit logging into services (remove from controllers)
2. Implement BaseRepository generic pattern
3. Error categorization system

---

## Estimated Codebase Reduction
- **Backend:** ~400-500 lines eliminated (boilerplate & duplication)
- **Frontend:** ~100-150 lines eliminated (stores & APIs)
- **Total:** ~15-20% reduction in domain/store code

## Maintainability Improvements
- 🟢 Consistency: All similar operations follow one pattern
- 🟢 Testability: Generics + factories enable parameterized tests
- 🟢 Safety: Enum-based entity types prevent bugs
- 🟢 Audit Guarantee: Logging moved to service layer ensures coverage
