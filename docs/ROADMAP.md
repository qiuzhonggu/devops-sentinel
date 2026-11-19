# Roadmap

Status of the product: **Planning / Initial Development**.

This document is the build order. The feature table in the [README](../README.md) stays at **Planned** until the exit criteria for that feature are met in the same pull request that changes the table.

A change that pulls work forward from a later phase updates this file first, in that pull request, with the reason.

## Phase 0 — Repository baseline

This is the current phase.

### Done in the initial commit

- Public project definition
- Target architecture
- Security design
- Contributor workflow, issue templates, and pull request template
- A GitHub Actions check that the required documents exist

### Not done

Phase 0 does not include a build, a release, or a runnable stack.

## Phase 1 — Agent, collection, transport, and ingest

Goal: one Linux host can ship redacted host metrics and logs to an API that stores them for that host's account.

### In scope

- Go agent binary
- CPU, memory, disk, and network collection
- journald collection and file collection
- offsets so collection resumes after a restart
- bounded memory queue and disk spool
- backoff, jitter, circuit breaker, and catch-up rate limit
- TLS 1.3 with mutual authentication
- per-agent client identity, including a way to revoke it
- secret redaction on the agent and again in the API
- backend ingest API
- PostgreSQL as the system of record for accepted ingest

### Out of scope

- Dashboard
- Log search UI
- Docker collection
- NATS JetStream consumers
- Alerts
- AI features
- Automated remediation

JetStream can be introduced earlier only if ingest cannot meet its exit criteria without it. That decision belongs in an issue before the code lands.

### Exit criteria

- `sentinel-agent` builds from this repository as a static Linux binary.
- On a Linux host it samples CPU, memory, disk, and network, and the samples show up through the ingest API.
- It reads journald and a configured file, stores an offset, and continues from that offset after a process restart.
- A fixture containing an API key, a bearer token, a password assignment, and a private-key block is stored only in redacted form. A test feeds the same fixture to the API as if the agent had skipped redaction, and the stored row is still redacted.
- With the API stopped, the spool grows only up to its quota and then reports overflow. With the API returned, the agent drains within its rate limit.
- An unknown client certificate is rejected. A revoked certificate is rejected before expiry.
- Rows written for one account are not returned by a query authenticated as another account.
- The README rows for these features are updated only in the pull request that adds the tests above.

## Phase 2 — Dashboard, fleet view, search, and alerts

Goal: an operator can watch more than one server, find a log line, and get notified about a condition without opening an SSH session.

### In scope

- Vue 3 and TypeScript dashboard
- sign-in and account-scoped sessions
- multi-server inventory
- log search
- alert rules and a delivery path for notifications
- NATS JetStream for asynchronous alert evaluation
- opt-in Docker container monitoring, documented as privileged because of the Docker socket
- Docker Compose for PostgreSQL, NATS, the API, and the dashboard

### Out of scope

- AI diagnosis
- automated remediation
- multi-region deployment

### Exit criteria

- An operator can sign in and see two agents that belong to that account, and cannot see an agent that belongs to a different account.
- Search returns a known ingested line and does not return a secret that the redaction tests cover.
- A configured condition creates one alert, and a notification is emitted once for that firing.
- Docker inventory runs only when enabled, and the default configuration leaves it off.
- `docker compose` starts the stack from the documented steps on a clean machine.

## Phase 3 — AI diagnosis and assisted remediation

Goal: help an operator explain a failure using data the platform already stored, without creating a new place where secrets accumulate.

### In scope

- analysis over redacted logs and metrics already in PostgreSQL
- anomaly detection on collected signals
- a root-cause report that cites the records it used
- suggested or automated actions that an operator has explicitly enabled

### Out of scope

- sending unredacted logs to a model
- a default setting that restarts services or deletes data
- hidden prompts that the operator cannot inspect

### Exit criteria

- A diagnosis response cites stored records by identifier.
- A test shows that a fixture secret never appears in the model request.
- Every automated action is opt-in, limited to a named operation, and written to the audit trail.
- The feature can be turned off without disabling collection.

## Working agreement

- One issue maps to one pull request whenever the work can be split that way.
- A pull request says which phase it belongs to.
- Devin sessions, when used, follow the same exit criteria as a human contributor. The README explains that workflow.
- New collectors, transports, and storage backends need a short note in [ARCHITECTURE.md](ARCHITECTURE.md) and, when they touch identity or data handling, in [SECURITY.md](SECURITY.md).
