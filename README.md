# AI-Ops Desktop

AI-Ops Desktop is an observability and operational automation tool designed to run as a desktop application with a Go backend and Svelte frontend. The goal is to deliver a complete workflow: host metrics collection, alert engine, safe auditable actions, and an AI assistant to explain alerts — all packaged as a desktop application.

**Status:** Sprint 4 Complete ✅ | Next: Sprint 5 (AI Explanation Flow)

## Why this project?
- To have a lightweight tool for monitoring and acting on local or remote machine issues.
- To prototype telemetry integration, alert rules, and automation with a focus on security.
- To package as a desktop app for easier use by operators and engineers.

## Key Technologies
- **Backend:** Go (slog for logging, gopsutil for metrics, SQLite for persistence)
- **Frontend:** Svelte (SPA connected to backend via HTTP + SSE)
- **Desktop:** Electron (to package frontend + start backend)
- **Persistence:** SQLite (Sprint 6)
- **AI:** Provider-agnostic abstraction for model calls (Sprint 5)

## Current Progress
- ✅ **Sprint 1:** Go Foundation (completed)
- ✅ **Sprint 2:** Live Metrics Pipeline (completed)
- ✅ **Sprint 3:** Alert Engine (completed)
- ✅ **Sprint 4:** Safe Automation (completed)
- ⏸️ **Sprints 5-8:** Planned

The frontend already includes container control surfaces and the backend now exposes guarded action and history flows implemented in Sprint 4. The AI explanation flow, persistence refinements, packaging, and polish remain in upcoming sprints.

## Design Decisions
- Minimalist interface without glow/shadow effects for better readability
- Flat colors and clean design focused on functionality
- Real-time updates via Server-Sent Events (SSE)
- Structured logging with `slog` in the backend

## Getting Started (Development)

### 1. Backend
Open a terminal in the `backend` folder:

```bash
cd backend
go run ./cmd/api
```

**Important Endpoints:**
- `GET /health` — Service health check
- `GET /metrics` — JSON metrics snapshot
- `GET /metrics/stream` — SSE of snapshots (Sprint 2)
- `GET /alerts` — Alert list (Sprint 3)
- `POST /alerts/:id/acknowledge` — Acknowledge alert (Sprint 3)
- `POST /alerts/:id/silence` — Silence alert (Sprint 3)
- `POST /ai/explain-alert` — Generate AI explanation (Sprint 5)
- Action and history endpoints — Implemented in Sprint 4 (see backend docs for exact routes and payloads)

### 2. Frontend
Open a terminal in the `frontend` folder:

```bash
cd frontend
npm install  # or pnpm/yarn
npm run dev
```

Configure the API base URL (centralized variable) to point to the local backend in development.

The app consumes `/health`, `/metrics` and subscribes to `/metrics/stream` (SSE).

**Environment Variables:**
```bash
VITE_API_BASE_URL=http://localhost:8080
```

### 3. Desktop Package (Development)
The Electron UI should start the backend locally and connect via a predictable local port or IPC (Sprint 7).

**Notes:**
- Enable CORS on the backend for local development.
- Structured logging with `slog` in the backend.
- Use SQLite for alert and action persistence when implementing Sprint 6.

## Roadmap & Sprints

Use 1-week sprints. Rule: each sprint ends with something visible in the UI and real backend progress.

**Sprint template (each sprint should contain):**
- Clear goal
- Demoable outcome
- Backend / frontend scope
- Acceptance criteria
- One explicit risk

### Planned Sprints:

#### Sprint 1 — Go Foundation ✅
**Goal:** Replace deleted Python scaffold with initial Go backend.
**Demo:** Svelte app loads data from a Go server instead of mocks.

**Backend:** `go.mod`, `cmd/api/main.go`, config, `slog`, `GET /health`, `GET /metrics` with typed mock data, CORS.

**Frontend:** Centralize base URL, fetch `/health` and `/metrics`, use typed responses.

**Acceptance:** Go server starts with 1 command; `/health` → 200; `/metrics` → valid JSON; dashboard renders from Go.

**Risk:** Discussing structure too much instead of delivering endpoints.

#### Sprint 2 — Live Metrics Pipeline ✅
**Goal:** Transform mock API into real telemetry service.
**Demo:** Cards update with real machine metrics in near real-time.

**Backend:** Normalized metrics models, `gopsutil` integration, collector loop with `context` and `time.Ticker`, in-memory latest snapshot store, `GET /metrics/stream` SSE.

**Frontend:** Subscribe to SSE, update dashboard cards without refresh, loading and disconnected states.

**Acceptance:** CPU, memory, disk, and network values come from the host machine; SSE reconnect behavior is acceptable in dev; dashboard updates without page reload.

**Risk:** Cross-platform metric differences and noisy data shaping.

#### Sprint 3 — Alert Engine ✅
**Goal:** Convert telemetry into stateful alerts.
**Demo:** When stressing the machine, alerts appear on the Alerts page.

**Backend:** Alert models, CPU/mem/disk rules, short evaluation windows, in-memory store, endpoints to list/ack/silence.

**Frontend:** Connect Alerts to backend, show severity/status/timestamps, acknowledge/silence actions.

**Acceptance:** At least 3 alert types; status persists in memory; alerts page uses real data.

**Risk:** Evaluation rule becoming too complex too early.

#### Sprint 4 — Safe Automation ✅
**Goal:** Add guarded, auditable operational actions, including container control.
**Demo:** UI executes auditable actions with confirmation, including at least one container restart or scale flow.

**Backend:** Action models, history, safe executor, allowlisted actions such as `kill_process`, `restart_container`, and `scale_container`, target validation, feature flags.

**Frontend:** Confirmation modal, action states (pending/success/failure), toggle for risky actions, container control surface wired to real endpoints.

**Acceptance:** Every intent is recorded; dangerous actions require confirmation; no generic shell execution; at least one container control action works end-to-end from UI to backend.

**Risk:** Runtime-specific container behavior and permissions making the action layer too broad or unsafe.

#### Sprint 5 — AI Explanation Flow ✅
**Goal:** Explain alerts in structured operational language.
**Demo:** Select alert and see AI summary, probable cause, and next steps.

**Backend:** Provider abstraction, `POST /ai/explain-alert`, build prompt with context, normalize output, timeout/errors.

**Frontend:** Assistant connected to alert context, section rendering, fallback when AI unavailable.

**Acceptance:** One complete explanation route; AI failure doesn't break; structured response.

**Risk:** Spending too much time on prompt tuning.

#### Sprint 6 — Persistence And History ⏸️
**Goal:** Preserve operational history across restarts.
**Demo:** Alerts and actions remain after restart.

**Backend:** SQLite, persist alerts and actions, metrics rollup, history queries.

**Frontend:** History page, simple tables and charts, basic export.

**Acceptance:** Alerts and action history survive restart; history no longer depends on mocks.

**Risk:** Trying to persist all samples and becoming a retention problem.

#### Sprint 7 — Desktop Packaging ⏸️
**Goal:** Package as desktop app.
**Demo:** App opens in Electron window and starts backend.

**Backend:** Prod config, predictable port/IPC, clean startup/shutdown.

**Frontend:** Integrate Electron, desktop menus, configure API base in packaged mode.

**Acceptance:** Packaged app opens, backend starts, main flows work.

**Risk:** Packaging complexity masking unresolved problems.

#### Sprint 8 — Polish And Portfolio ⏸️
**Goal:** Make repository presentable.
**Demo:** New user understands the app; repo ready for demo.

**Backend:** Better logs, error responses, API docs, missing tests.
**Frontend:** Empty/loading/error states, copy, onboarding, screenshots.

**Acceptance:** Clear README; screenshots; app stable for presentation.

**Risk:** Focusing only on appearance and not reliability.

**Recommended start:** Sprint 1 → Sprint 2 → Sprint 3 → Sprint 4.

## Acceptance Criteria & Sprint Discipline

At the end of each sprint, review only:
- What is demoable
- What is blocked
- What changed in the API contract

If a sprint generates a lot of internal code without anything visible, re-scope.

## API (minimum expected contract initially)

### Health Check
```
GET /health
200 OK with { "status": "healthy" }
```

### Metrics
```
GET /metrics
JSON with typed shape (cpu, memory, disk, network)
```

### Metrics Stream
```
GET /metrics/stream
SSE delivering periodic snapshots
```

### Alerts
```
GET /alerts
List of alerts with id, severity, status, timestamps, description

POST /alerts/:id/acknowledge
POST /alerts/:id/silence
```

### AI Explanation
```
POST /ai/explain-alert
Receives { alert_id, context } and returns { summary, probable_cause, suggested_actions }
```

### Actions & History
Action and history endpoints — implemented in Sprint 4; these are guarded, auditable flows (see backend docs for exact routes and payloads).

Keep response forms consistent across sprints; document changes in README and changelog.

## Development & Best Practices

- Prefer small iterations with visible frontend results.
- Backend tests: `go test ./...` for alert/action logic.
- Structured logs and readable messages for local debugging.
- Avoid exposing shell/generic commands; allow a small set of authorized actions.
- Feature flags for dangerous actions.
- Timeouts and circuit-breakers for external AI provider calls.

## Contribution

- Work in sprint branches: `sprint/1-go-foundation`, `sprint/2-live-metrics`, etc.
- Each PR should have:
  - Short description
  - Sprint acceptance checklist
  - How to test locally (minimum steps)
- Small, focused commits with demonstrable value.

## Immediate Roadmap (first suggested steps)

1. **Sprint 1 Complete** ✅: Create minimal Go backend and `health`/`metrics` endpoints.
2. **Sprint 2 Complete** ✅: Configure frontend to consume `/health`, `/metrics`, and `/metrics/stream`.
3. **Sprint 3 Complete** ✅: Alert Engine with stateful alerts.
4. **Sprint 4 Complete** ✅: Safe, auditable actions and action history wired end-to-end.
5. **Sprint 5 Next** 🔄: AI Explanation Flow.

## Contact

- Maintainer: @momarinho
- Issues: Open issues in the repository for bugs and feature proposals.

## License

Define an appropriate license (for example MIT) by adding a LICENSE file to the repository.

---

**Last Updated:** 2026-04-17
**Current Sprint:** 4 (Complete)
**Next Sprint:** 5 (AI Explanation Flow)
