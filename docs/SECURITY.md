# Security Design

Status: **design requirements**. Nothing in this document is implemented yet. The short version used by operators and contributors is the Security Design section of the [README](../README.md). Report a vulnerability through the process in the root [SECURITY.md](../SECURITY.md).

## Principles

- Collect the minimum that explains an operational failure.
- Authenticate both ends of every agent connection.
- Remove secrets before storage, twice.
- Put a finite bound on every retry, queue, and cache.
- Keep one account's data out of another account's queries.
- Prefer a visible failure over a silent fallback to a weaker mode.

## Least privilege

| Actor | Allowed | Denied |
| --- | --- | --- |
| Agent OS user | Read configured log paths, write its spool and offset files, present its own client certificate | Shell execution, reading other users' home directories, holding database or NATS credentials |
| Agent process | Ship redacted telemetry to the API | Accept remote commands, disable TLS, use a fleet-wide shared secret |
| Backend | Validate identity, redact, read and write that account's rows | Trust client-supplied account IDs without binding them to the credential |
| Dashboard | Call the API as the signed-in operator | Connect to PostgreSQL or NATS, download agent private keys |
| Docker collector | Read container inventory when an operator turns it on | Default installation. The Docker socket can control the host, so this collector stays off until configured |

Log lines are data. They are never turned into shell commands.

## Mutual TLS

Agent transport requirements:

- TLS 1.3 is the minimum protocol version.
- The agent verifies the API certificate against the project CA, or against a pinned certificate, and checks the server name.
- The agent presents a client certificate issued for that agent.
- The API maps the certificate to one agent record and one account.
- Certificate authentication and user-session authentication stay separate. An operator password does not authenticate an agent, and an agent certificate does not open the dashboard.

A failed verification closes the connection. The client will not retry on cleartext HTTP.

## Rotation and revocation

- Issue short-lived client certificates.
- During rotation, accept the outgoing and incoming certificate together for a limited overlap so the agent can keep shipping.
- Store enough certificate identity on the server to recognize the agent. The private key stays on the host.
- Maintain a revocation list. The API checks it when the connection is established and when a certificate is presented again.
- Revocation takes effect before expiry. Removing one agent does not require new certificates for the others.
- Record issuance and revocation in the audit trail.

A stolen certificate is contained by revocation plus a short lifetime, not by lifetime alone.

## Redaction

Redaction runs on the agent before an event reaches the spool or the network. The API runs the same classes of rules again before insert. Server-side redaction does not trust the agent.

Planned pattern classes:

- cloud and vendor API keys
- bearer tokens and authorization headers
- passwords in `key=value` assignments and in URLs
- PEM private-key blocks
- connection strings with embedded credentials

A match is replaced with a fixed marker that names the class, for example `[redacted:token]`. The original value is not encrypted for later recovery. Tests will use fixtures, never live credentials.

Redaction will miss novel formats. That is why the server repeats the pass, why collection is limited to configured paths, and why AI features are not allowed to receive a second, raw copy.

## Local cache limits

The on-disk spool has a configured quota. The in-memory queue has a separate cap.

When the spool is full:

- the agent stops taking new events
- it reports the overflow as its own health signal
- it keeps the events already stored
- it does not delete retained logs to free space for newer lines

Metrics may be downsampled before that stop. Security-relevant logs are not the first thing discarded. Operators can lower the collection rate or raise the quota. They cannot set the quota to unlimited.

Spool files are readable only by the agent user.

## Retries and resource exhaustion

Unbounded retry is treated as a defect.

- Exponential backoff with a maximum delay.
- Randomized jitter on each delay.
- A cap on attempts for a single payload, after which the payload is recorded as deferred or discarded by an explicit policy.
- A circuit breaker that opens after consecutive transport failures and allows one probe after a cool-down.
- A rate limit on backlog drain so recovery does not stampede the API.
- Process budgets for CPU, memory, disk, and file descriptors. Crossing a budget sheds new work.

Each limit will have a test that proves the agent stops. "Retry until success" is not an acceptable default.

## Account isolation

Isolation is an application invariant, not a convention in one query.

- An account identifier is stored on agent, metric, log, alert, and audit rows.
- The identifier comes from the authenticated agent or the authenticated user, not from an unchecked field in the body.
- List and search endpoints include that identifier in the query.
- Jobs published later to JetStream carry it, and consumers check it again before they touch stored rows.
- Certificate issuance is scoped to the account that owns the host.

There is no cross-account support view in this design.

## Threats the design is meant to address

| Threat | Intended response |
| --- | --- |
| Passive network observer | TLS 1.3 |
| Agent talking to an impersonated API | API certificate verification |
| Stolen or copied agent credential used from elsewhere | Per-agent certificate, revocation list, short lifetime |
| Secret written into an application log | Agent redaction, server redaction, no reversible archive |
| Modified agent that skips redaction | Server redaction is mandatory |
| Disk fill by the spool or by retries | Spool quota, backoff cap, circuit breaker |
| Fleet reconnecting at the same moment | Jitter and a catch-up rate limit |
| Query that omits the account filter | Account identifier bound to the credential and applied in one place |
| Dashboard or model prompt used to extract secrets | Redaction before storage, and no raw-log path for AI |
| Docker socket used as an accidental root backdoor | Collector off by default and documented as privileged |

## What this document does not do

It does not claim the controls are present in code. It does not set production certificate lifetimes or quota numbers. Those values belong with the implementation, next to the tests that enforce them, and they can be revised with a measured reason.

When a control is implemented, the pull request should link the test that demonstrates it and should update the feature status in the README.
