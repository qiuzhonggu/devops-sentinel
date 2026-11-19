# Contributing to DevOps Sentinel

Thanks for considering a contribution. The project is in **Planning / Initial Development**. The useful work right now is precise documentation, clear issues, and, once a phase starts, small pull requests that match the roadmap.

## Ground rules

- Be specific and respectful in issues and reviews.
- Search existing issues and pull requests before opening a new one.
- Keep one pull request focused on one change.
- Describe behavior that exists today. Mark unfinished behavior **Planned**.
- Do not commit secrets, private keys, certificates, tokens, or logs from a real system.
- Security reports go through a [private advisory](https://github.com/qiuzhonggu/devops-sentinel/security/advisories/new), not a public issue. See [SECURITY.md](SECURITY.md).

## Issues

1. Look through the [issue list](https://github.com/qiuzhonggu/devops-sentinel/issues).
2. Open a new issue from the [templates](https://github.com/qiuzhonggu/devops-sentinel/issues/new/choose).
3. Use a bug report only for something already in the repository. A missing planned feature is a feature request.
4. For a change that affects architecture, security, or phase order, wait for a maintainer comment before writing a large patch.

A good issue states the problem, the expected outcome, and what is out of scope.

## Pull requests

1. Fork the repository and create a branch from `main`.
2. Prefer a short branch name tied to the issue, such as `docs/phase-1-exit-criteria` or `agent/journald-offset`.
3. Make the change. Update the README feature table in the same pull request when a feature moves from **Planned** to implemented.
4. Run `bash scripts/check-docs.sh`.
5. Open a pull request. The template is applied automatically.
6. Link the related issue. Fill in the test plan.
7. A maintainer reviews the pull request. Address comments with new commits on the same branch.
8. The maintainer merges. Do not merge your own pull request unless you are the maintainer.

### Review expectations

- Documentation must match the tree. A README claim without code is a defect.
- Authentication, TLS, redaction, retention, and account isolation need maintainer review.
- Generated files, vendored dependencies, and unrelated formatting changes do not belong in a feature pull request.
- AI-assisted patches, including patches from Devin, use the same review path as any other contribution. See the README section "Development with Devin Bot".

## Development setup

No language toolchain is required for the current tree.

```bash
git clone https://github.com/qiuzhonggu/devops-sentinel.git
cd devops-sentinel
bash scripts/check-docs.sh
```

When Go or Vue code is added, the pull request that adds it will also document the exact toolchain versions and test commands.

## Commit messages

Write the subject in the imperative mood and say why the change exists.

```text
Document Phase 1 exit criteria for ingest.

Reviewers need a testable finish line before agent code starts.
```

Keep the subject under about 72 characters. Add a body when the subject is not enough.

## License

By submitting a contribution, you agree that it is licensed under the Apache License, Version 2.0, the same license as the rest of the repository. See [LICENSE](LICENSE).
