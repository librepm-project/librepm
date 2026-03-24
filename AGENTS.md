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

### 2. Frontend Layering (React / TypeScript)
*   **Presentation Layer (AI-Augmented UI):** Components like `ChatWindow`, `DiffViewer`, and `AI-Suggestion-Badge`.
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
- **Frontend:** React with specialized streaming hooks.

## Security & Cost Management

1.  **Rate Limiting:** Strict quotas per user/agent.
2.  **Data Privacy:** Masking sensitive data (PII) before sending it to external LLMs.
3.  **Audit Log:** All AI actions must be flagged as "bot-generated" and logged for review.
