# Quarantine: edits from autopilot session 2026-04

## Why this branch exists

These commits landed on a PC-1 working tree but were almost certainly
made **accidentally**. The user (Jason White) was working on autopilot
firmware (gus-px4 / gus-airship-configs) at the time and believes one
of the agents in the session changed directory into this repo without
explicit instruction and made these edits.

The user has **not** verified that the changes are useful for the
function of this repo. They may be:

- Internally consistent and useful → integrate via PR
- Internally consistent but irrelevant to this repo's actual mission → discard
- Stale or partial work that was meant for a different repo → port elsewhere

## What's in here (vs `main` of `Constellation-Overwatch/aero-arc-relay2constellation`)

The 9 commits on this branch (most recent first):

| Commit | Subject |
|---|---|
| 9f090b2 | Merge branch 'main' |
| 7ca27a0 | feat(config): add PX4-compatible MAVLink dialects (all, px4, development) |
| c9c5f96 | docs: update |
| 1774346 | config: enhanced multi-sample |
| e75fe0d | feat: nats_auth, isolated multi-bird stream confirmed |
| c99f235 | feat: nats kv to config |
| aa86946 | Merge pull request #1 from Constellation-Overwatch/nats-jetstream-support |
| e779075 | feat: rip & replace kafka for nats only |
| 0367cfd | deps: add nats |

Net change vs origin/main covers: NATS replacement of Kafka, NATS auth +
multi-bird stream, NATS KV-to-config, PX4 MAVLink dialects (all/px4/development),
Taskfile.yml additions, README expansion, multi-sample config example.

## Recommended next step

Tobalo (or whoever owns this repo's roadmap) should review whether any of
these edits should be cherry-picked onto main. If none of it is wanted,
the branch can be deleted.

Local PC-1 main has been reset to origin/main as of the date of this
branch creation, so PC-1 will not regenerate this divergence on the next
session.

— created by Claude Code on PC-1, 2026-04-27
