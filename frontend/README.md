# Frontend

Status: **Planned**. This directory has no application yet. There is no `package.json` and nothing to install.

The dashboard is scheduled for Phase 2. It will be a Vue 3 and TypeScript application. It will sign operators in, list the servers in their account, search logs, and show alerts by calling the backend API.

## Intended layout

```text
frontend/
├── index.html
├── package.json
├── vite.config.ts
└── src/
    ├── main.ts
    ├── router/
    ├── views/
    └── api/
```

## Rules for the first implementation

- Call the API only. Do not embed database or NATS credentials.
- Keep session handling on the operator path, separate from agent certificates.
- Render redacted log text as text, not as HTML.
- An account view shows that account's agents only.

Phase 3 diagnosis views, if they need a UI, will be added after the Phase 2 console is in place. See [docs/ROADMAP.md](../docs/ROADMAP.md).
