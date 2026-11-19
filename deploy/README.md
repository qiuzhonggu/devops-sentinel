# Deployment

Status: **Planned**. There is no Docker Compose file and there are no images to run.

Docker Compose is the intended way to start a single-node control plane once the backend and dashboard exist. The planned services are:

| Service | Role | Phase |
| --- | --- | --- |
| `postgres` | System of record | 1 |
| `backend` | Go API | 1 |
| `nats` | JetStream bus for asynchronous work | 2 |
| `dashboard` | Vue 3 console | 2 |

The agent is installed on the monitored hosts. It is not a Compose service on those hosts, because the point of the agent is a single binary with a small footprint.

## Rules for the first Compose file

- It will be added in the pull request that can actually start the services it defines.
- It will not ship real certificates, passwords, or API keys. Local credentials will be generated and ignored by git.
- Published ports will be limited to the dashboard and the API.
- PostgreSQL and NATS will stay on the Compose network.
- The Docker log collector will stay disabled in the sample agent configuration.

Certificate material is ignored by the repository `.gitignore` (`*.pem`, `*.key`, `*.crt`, and related files). Test certificates should be generated when a test runs, not committed.

See [docs/ARCHITECTURE.md](../docs/ARCHITECTURE.md) and [docs/ROADMAP.md](../docs/ROADMAP.md).
