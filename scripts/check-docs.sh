#!/usr/bin/env bash
# Verify that the repository still contains its required planning documents.
# This does not build or test the agent, backend, or dashboard.
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${root}"

required=(
  README.md
  LICENSE
  CONTRIBUTING.md
  SECURITY.md
  docs/ARCHITECTURE.md
  docs/SECURITY.md
  docs/ROADMAP.md
  agent/README.md
  backend/README.md
  frontend/README.md
  deploy/README.md
)

missing=0
for file in "${required[@]}"; do
  if [[ ! -f "${file}" ]]; then
    printf 'missing: %s\n' "${file}"
    missing=1
  fi
done

if [[ "${missing}" -ne 0 ]]; then
  exit 1
fi

printf 'required documents present\n'
