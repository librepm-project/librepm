# LibrePM AI Agents & Automation Architecture

This document outlines the planned structure and guidelines for AI-based developments within the LibrePM project. The goal is to build a scalable, modular system that supports intelligent project management features.

## Architecture Overview

The system consists of three main layers:
1.  **AI Orchestrator (App):** The central control unit managing the lifecycle of agents and communication with external LLMs.
2.  **AI Shared Libraries (Libs):** Common functions for prompt management, embeddings, and tools.
3.  **Domain-Specific Tools:** Interfaces connected to the existing API, allowing the AI to interact with and modify data.

## Planned Directory Structure

```text
/
├── apps/
│   └── ai-orchestrator/           # Central AI service (Go or Node.js)
│       ├── internal/
│       │   ├── agent/             # Agent definitions (TriageAgent, SummaryAgent)
│       │   └── chain/             # Orchestration logic or LangChain flows
│       └── project.json
│
├── libs/
│   ├── ai-core/                   # Shared AI interfaces and LLM clients
│   │   ├── llm/                   # Adapters for OpenAI, Anthropic, Ollama, etc.
│   │   └── prompt-templates/      # Versioned prompt collections (.md files)
│   │
│   ├── ai-tools/                  # Capabilities "Tools" agents can invoke
│   │   ├── issue-manager/         # Reading/writing issues from an AI perspective
│   │   └── project-analyzer/      # Preparing statistical data for LLMs
│   │
│   └── vector-db/                 # Vector database client (e.g., Qdrant, PGVector)
│
└── apps/frontend/src/features/ai/ # UI components for AI interactions
```

## Naming Conventions (Terminology)

To maintain consistency across the monorepo, the following naming patterns must be followed:

### 1. Agents
- **Pattern:** `[Feature]Agent`
- **Description:** Autonomous entities with a specific goal and persona.
- **Examples:** `TriageAgent`, `EstimateAgent`, `SupportAgent`.

### 2. Skills & Tools
- **Pattern:** `[Action][Entity]Skill` or `[Entity]Tool`
- **Description:** Atomic capabilities that agents can call.
- **Examples:** `UpdateIssueSkill`, `SearchDocumentationTool`, `GenerateReportSkill`.

### 3. Prompts
- **Pattern:** `prompt.[agent-name].[action].md`
- **Description:** Stored as Markdown files for better readability and version control.
- **Examples:** `prompt.triage.classify.md`, `prompt.summary.sprint.md`.

### 4. Services & Components
- **Orchestrator:** The main service managing multiple agents.
- **Provider:** Wrappers for external LLM APIs (e.g., `OpenAIProvider`).
- **Chain:** A sequence of predefined AI steps.

## Layering Model

### 1. Backend Layering (Go / Python)
*   **API / Transport Layer:** REST/gRPC endpoints with SSE (Server-Sent Events) support for streaming.
*   **Orchestration Layer:** The "brain" coordinating multi-step workflows.
*   **Agent / Skill Layer:** Specific AI logic, prompt assembly, and output validation.
*   **Tool / Adapter Layer:** Interfaces to existing Domain Services (e.g., `IssueService`). AI never writes to the DB directly.
*   **Prompt Management Layer:** External template management (non-hardcoded strings).

### 2. Frontend Layering (Vue 3 / TypeScript with Vuetify)
*   **Presentation Layer (AI-Augmented UI):** Components like `ChatWindow`, `ThinkingIndicator`, `StreamingResponse`, and `AI-Suggestion-Badge`.
*   **Interaction / Hook Layer:** `useAIStream` for streaming and `useAgentStatus` for showing "thinking" states.
*   **Context Aware Layer:** Providers that collect UI state (open issue, selected text) as context for the AI.
*   **Service Layer (AI Client):** Client-side handling of orchestrator calls and retry logic.

## Key Functional Areas

1.  **Smart Triage Agent:** Automatic categorization, priority setting, and assignment of new issues.
2.  **Sprint Summarizer:** Generating weekly summaries and sprint reports based on activity logs.
3.  **Documentation Assistant:** Natural language search over project docs and issues (RAG).

## Tech Stack Recommendations

- **Backend:** Go (primary) or Python/FastAPI (for ML-heavy tasks).
- **LLM Connectivity:** OpenAI/Anthropic SDKs or LangChainGo.
- **Vector DB:** PGVector or Qdrant.
- **Frontend:** Vue 3 with Vuetify and specialized streaming hooks.

## Security & Cost Management

1.  **Rate Limiting:** Strict quotas per user/agent.
2.  **Data Privacy:** Masking sensitive data (PII) before sending it to external LLMs.
3.  **Audit Log:** All AI actions must be flagged as "bot-generated" and logged for review.

---

## UI Design System & Guidelines

LibrePM implements a modern, professional design system using **Vuetify 4** as the CSS framework. The following guidelines ensure consistency across AI-augmented features and the entire application.

### Design Philosophy

- **Modern & Clean:** Flat design with subtle gradients and smooth transitions
- **Accessible:** High contrast, semantic color usage, keyboard navigation support
- **Responsive:** Works seamlessly on mobile, tablet, and desktop
- **Component-Based:** Reusable, modular components following Vuetify patterns
- **Performance-Focused:** Lightweight, efficient animations and lazy loading

### Color Palette

| Element | Color | Hex | Usage |
|---------|-------|-----|-------|
| Primary | Deep Blue | `#3F51B5` | Main actions, links, highlights |
| Secondary | Light Blue | `#5C6BC0` | Supporting UI, hover states |
| Success | Green | `#4CAF50` | Positive actions, confirmations |
| Error | Red | `#F44336` | Destructive actions, errors |
| Warning | Amber | `#FF9800` | Warnings, alerts |
| Info | Cyan | `#00BCD4` | Informational messages |
| Surface | Light Gray | `#F5F7FA` | Background, cards, surfaces |
| Borders | Gray | `rgba(0,0,0,0.08)` | Subtle borders, dividers |
| Text Primary | Dark Gray | `#212121` | Main text |
| Text Secondary | Medium Gray | `#666666` | Secondary text |

### Typography

| Element | Style | Usage |
|---------|-------|-------|
| h1 | Roboto 32px bold | Page titles |
| h2 | Roboto 28px bold | Section headers |
| h3 | Roboto 24px bold | Subsection headers |
| h4 | Roboto 20px bold | Card titles |
| h5 | Roboto 18px bold | Form titles, content headers |
| h6 | Roboto 16px bold | Component titles, labels |
| body1 | Roboto 16px regular | Main text content |
| body2 | Roboto 14px regular | Secondary text, descriptions |
| subtitle1 | Roboto 16px medium | Section subtitles |
| subtitle2 | Roboto 14px medium | Field labels, hints |
| button | Roboto 14px medium | Button text |
| code | Courier 13px monospace | Code blocks, JSON editors |

### Spacing System

All spacing follows an 8px base grid:

```
xs: 4px    (0.5 unit)
sm: 8px    (1 unit)
md: 16px   (2 units)
lg: 24px   (3 units)
xl: 32px   (4 units)
2xl: 48px  (6 units)
```

**Usage in Vuetify utilities:**
- `pa-2` = 16px padding (all sides)
- `mb-4` = 32px margin (bottom)
- `gap-3` = 24px gap (flex layout)

### Layout Components

#### 1. **v-app-bar** (Header)
- **Background:** Primary color (`#3F51B5`)
- **Elevation:** 4
- **Height:** 64px
- **Features:**
  - Logo + branding on left
  - Navigation menu with hover dropdowns
  - Search bar (max-width: 300px)
  - White text for contrast

#### 2. **v-card** (Content Containers)
- **Elevation:** 0 (flat) on surfaces, 2 on hover
- **Border-radius:** `rounded-xl` (16px)
- **Background:** Gradient `linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%)`
- **Padding:** `pa-6` (24px)
- **Spacing:** `mb-6` between cards (32px)

#### 3. **v-data-table** (Lists/Tables)
- **Header:** Primary background with white bold text
- **Rows:** Transparent background with subtle hover effect
- **Hover:** `rgba(63, 81, 181, 0.05)` background
- **Borders:** Bottom border `1px solid rgba(0,0,0,0.08)`
- **Density:** Comfortable spacing for readability

#### 4. **v-text-field** (Forms)
- **Variant:** Outlined or Solo-filled
- **Background:** `rgba(255,255,255,0.8)` when filled
- **Border-radius:** `rounded` or custom via class
- **Density:** Compact for space efficiency
- **Icons:** Prepend/append supported

#### 5. **v-btn** (Actions)
- **Default Color:** Primary
- **Variants:** Filled, Outlined, Text
- **Border-radius:** `rounded-lg` (8px)
- **Icon Position:** Prepend or append
- **Sizes:** small, default, large
- **Elevation:** 1 for filled buttons

#### 6. **v-navigation-drawer** (Sidebar)
- **Width:** Responsive (200px on desktop)
- **Elevation:** 2
- **Background:** Gradient surface
- **Border:** Right border `1px solid rgba(0,0,0,0.05)`
- **Items:** Rounded corners with hover highlight

### AI Component Patterns

#### **ChatWindow Component**
```vue
<v-card elevation="2" class="rounded-xl">
  <v-card-title class="bg-primary text-white">
    <v-icon>mdi-chat</v-icon> AI Assistant
  </v-card-title>
  <v-card-text class="chat-content pa-6">
    <!-- Messages here -->
  </v-card-text>
  <v-card-actions class="gap-3">
    <v-text-field placeholder="Ask AI..." class="flex-grow-1" />
    <v-btn color="primary" icon="mdi-send" />
  </v-card-actions>
</v-card>
```

#### **ThinkingIndicator Component**
```vue
<v-chip 
  label 
  color="info" 
  size="small"
  class="animated-pulse"
>
  <v-icon start>mdi-lightbulb</v-icon>
  Thinking...
</v-chip>
```

#### **StreamingResponse Component**
```vue
<v-card class="mt-4 pa-4 rounded-xl">
  <v-progress-linear 
    indeterminate 
    color="primary"
  />
  <p class="mt-2 text-body2">{{ streamingText }}</p>
</v-card>
```

#### **AI-Suggestion-Badge Component**
```vue
<v-chip
  size="small"
  color="success"
  variant="elevated"
  prepend-icon="mdi-lightbulb-outline"
>
  AI Suggestion
</v-chip>
```

### Form Patterns

#### **Standard Form Layout**
```vue
<v-form @submit.prevent="submit" ref="form">
  <v-card elevation="0" class="rounded-xl pa-6 mb-6">
    <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
      Form Title
    </v-card-title>

    <v-row>
      <v-col cols="12" md="6">
        <v-text-field
          v-model="field1"
          :rules="[requiredRule]"
          label="Field 1"
          outlined
          dense
        />
      </v-col>
      <v-col cols="12" md="6">
        <v-text-field
          v-model="field2"
          label="Field 2"
          outlined
          dense
        />
      </v-col>
    </v-row>

    <v-divider class="my-6" />

    <v-row class="mt-4">
      <v-col cols="12" class="d-flex gap-3">
        <v-btn
          type="submit"
          color="primary"
          size="large"
          prepend-icon="mdi-check"
          rounded="lg"
          class="font-weight-bold"
        >
          Submit
        </v-btn>
        <v-btn
          type="button"
          variant="outlined"
          color="default"
          size="large"
          rounded="lg"
          @click="$router.back()"
        >
          Cancel
        </v-btn>
      </v-col>
    </v-row>
  </v-card>
</v-form>
```

### Table Patterns

#### **Standard Data Table**
```vue
<v-card elevation="0" class="rounded-xl">
  <v-data-table
    :items="items"
    :headers="headers"
    density="comfortable"
    class="transparent-table"
    item-key="id"
  >
    <template #header>
      <thead class="bg-primary text-white">
        <tr>
          <th
            v-for="header in headers"
            :key="header.key"
            class="text-white font-weight-bold"
          >
            {{ header.title }}
          </th>
        </tr>
      </thead>
    </template>

    <template #item.actions="{ item }">
      <v-icon
        size="small"
        class="me-2 cursor-pointer"
        @click="handleEdit(item)"
      >
        mdi-pencil
      </v-icon>
      <v-icon
        size="small"
        color="error"
        class="cursor-pointer"
        @click="handleDelete(item)"
      >
        mdi-delete
      </v-icon>
    </template>
  </v-data-table>
</v-card>
```

### CSS Classes & Utilities

#### **Responsive Spacing**
```css
.pa-6    /* padding: 24px (all sides) */
.px-4    /* padding: 16px (horizontal) */
.py-2    /* padding: 8px (vertical) */
.mb-6    /* margin-bottom: 24px */
.mt-4    /* margin-top: 16px */
.gap-3   /* gap: 24px (flex/grid) */
```

#### **Responsive Layout**
```css
.d-none.d-md-flex  /* Hidden on mobile, flex on tablet+ */
cols="12" md="6"   /* Full width mobile, 50% on tablet+ */
cols="12" md="8" lg="10"  /* Responsive column widths */
```

#### **Utilities for AI Components**
```vue
<div class="animated-pulse">Animated pulse effect</div>
<div class="text-gradient">Gradient text effect</div>
<div class="shadow-elevated">Elevated shadow effect</div>
```

### Implementation Best Practices

1. **Always use Vuetify components** instead of raw HTML/CSS
2. **Maintain consistent spacing** using the 8px grid
3. **Apply rounded-xl** to all card-like containers
4. **Use color props** instead of hardcoded hex values
5. **Responsive-first** design: mobile → tablet → desktop
6. **Accessibility:** Always label form fields, use semantic HTML
7. **Performance:** Lazy-load heavy components, optimize re-renders
8. **Dark mode ready:** Use theme variables, not hardcoded colors

### AI-Specific UI Patterns

#### **Real-time Streaming**
- Show progress indicators during AI processing
- Display thinking states with animated spinners
- Stream responses character-by-character with v-textarea
- Allow cancellation of long-running operations

#### **Confidence & Reasoning**
- Display AI confidence levels as badges
- Show reasoning/explanation in collapsible sections
- Allow users to thumbs up/down responses

#### **Context Display**
- Show context used by AI in a collapsible panel
- Highlight relevant sections from documents
- Display source references with links

#### **Error Handling**
- Show AI errors in v-alert components
- Provide "Try Again" or "Refine" actions
- Log errors for debugging but show user-friendly messages
