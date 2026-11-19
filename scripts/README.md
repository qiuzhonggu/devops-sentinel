# Scripts

Status: the documentation check below is available. Product installers, certificate helpers, and load generators are **Planned**.

## Available

| Script | Purpose |
| --- | --- |
| `check-docs.sh` | Fail if a required planning document is missing. GitHub Actions runs it on every push to `main` and on pull requests. |

Run it locally:

```bash
bash scripts/check-docs.sh
```

## Planned

- A development CA helper for short-lived test certificates. Generated keys will stay out of git.
- Lint and test entrypoints, added in the same change that introduces the first Go packages.
- A local stack helper, added when Docker Compose can start real images.
