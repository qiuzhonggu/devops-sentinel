# Backend

Status: **Planned**. This directory has no Go code yet.

The backend will be a Go API. It will authenticate each agent by client certificate, authenticate operators separately, redact ingest a second time, and store accepted data in PostgreSQL. The dashboard will use this API and will not talk to the database.

## Intended layout

```text
backend/
├── cmd/sentinel-api/
├── internal/http/
├── internal/identity/
├── internal/redact/
├── internal/store/
└── internal/alert/
```

Future module path: `github.com/qiuzhonggu/devops-sentinel/backend`.

`internal/alert/` and any NATS JetStream client belong to Phase 2. Phase 1 is ingest, identity, redaction, and PostgreSQL.

Draft routes are listed in [docs/ARCHITECTURE.md](../docs/ARCHITECTURE.md). They are not served by anything in this repository.

## Rules for the first implementation

- Take the account identifier from the authenticated credential.
- Reject unknown and revoked agent certificates.
- Redact again before insert.
- Keep user sessions and agent certificates as different credentials.
- Record certificate issuance and revocation.
