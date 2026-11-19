# Agent

Status: **Planned**. This directory has no Go code yet.

`sentinel-agent` will be a single static Go binary for Linux. It will collect host metrics and logs, redact secrets, keep a bounded local spool, and ship them to the backend over TLS 1.3 with a per-agent client certificate.

## Intended layout

```text
agent/
├── cmd/sentinel-agent/
├── internal/collect/
├── internal/docker/
├── internal/redact/
├── internal/buffer/
├── internal/transport/
└── internal/identity/
```

Future module path: `github.com/qiuzhonggu/devops-sentinel/agent`.

`internal/docker/` is reserved for the Phase 2 opt-in collector. Phase 1 does not include it.

## Rules for the first implementation

- Resume file and journald reads from a stored offset.
- Redact before writing the spool.
- Stop intake when the spool quota is full.
- Back off with jitter, open a circuit breaker, and rate-limit catch-up.
- Refuse a cleartext fallback when TLS or certificate checks fail.
- Run as an unprivileged user and stay inside an explicit CPU, memory, disk, and file-descriptor budget.

The behavior is specified in [docs/ARCHITECTURE.md](../docs/ARCHITECTURE.md), [docs/SECURITY.md](../docs/SECURITY.md), and [docs/ROADMAP.md](../docs/ROADMAP.md).
