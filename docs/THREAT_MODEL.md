# VRP Threat Model

## Purpose

This document defines the current threat model assumptions for the public VRP runtime proof surface.

The goal is not to claim perfect security.

The goal is to define:

- what VRP attempts to protect
- what VRP assumes
- what VRP validates
- what VRP does not guarantee

The repository focuses on continuity-safe execution semantics under unstable and hostile runtime conditions.

---

# Core Security Principle

```text
Transport input is not trusted automatically.
```

Packet arrival alone must not imply canonical execution truth.

All execution input must pass through:

- authentication validation
- replay validation
- authority validation
- epoch validation
- recovery validation

before crossing the canonical execution boundary.

---

# Canonical Execution Boundary

VRP defines a deterministic execution admission boundary.

Correct flow:

```text
transport receive
→ validation pipeline
→ canonical admission decision
→ application mutation
```

Incorrect flow:

```text
transport receive
→ automatic execution acceptance
```

The runtime attempts to ensure that invalid input does not silently mutate canonical state.

---

# Security Goals

The current public proof surface attempts to validate:

- replay-safe execution
- stale epoch rejection
- stale authority rejection
- authenticated mutation boundaries
- deterministic recovery validation
- fail-closed runtime behavior
- continuity preservation across transport instability

---

# Threat Assumptions

The runtime assumes:

- unstable networking
- duplicated packets
- delayed delivery
- replay attempts
- tampered payloads
- stale execution branches
- hostile transport conditions
- conflicting recovery attempts

The runtime does not assume transport reliability.

---

# Replay Threats

Threat:

```text
captured mutation replayed later
```

Goal:

```text
same logical mutation
must not commit twice
```

Protection model:

- mutation identity tracking
- replay window validation
- deterministic duplicate rejection

Expected verdict:

```text
REPLAY_REJECTED
```

---

# MITM Tampering Threats

Threat:

```text
payload modified during transport
```

Goal:

```text
tampered execution input
must not become canonical truth
```

Protection model:

- authenticated mutation boundaries
- deterministic authentication verification

Expected verdict:

```text
AUTHENTICATION_FAILED
```

---

# Stale Epoch Threats

Threat:

```text
older execution branch attempts overwrite
```

Goal:

```text
older execution state
must not replace newer canonical state
```

Protection model:

- monotonic epoch validation
- deterministic stale rejection

Expected verdict:

```text
STALE_EPOCH_REJECTED
```

---

# Stale Authority Threats

Threat:

```text
previous authority attempts mutation admission
```

Goal:

```text
non-canonical authority
must not advance execution
```

Protection model:

- authority identity validation
- canonical authority enforcement

Expected verdict:

```text
NON_AUTHORITY_REJECTED
```

---

# Fake Recovery Threats

Threat:

```text
invalid recovery lineage injected
```

Goal:

```text
invalid recovery branch
must not become canonical
```

Protection model:

- snapshot validation
- recovery lineage validation
- deterministic recovery rejection

Expected verdict:

```text
RECOVERY_SNAPSHOT_REJECTED
```

---

# Transport Threats

VRP assumes transport instability is normal.

Examples:

- route mutation
- carrier replacement
- disconnect windows
- NAT rebinding
- relay migration
- packet duplication
- packet reordering

The runtime attempts to preserve execution correctness independently of transport stability.

---

# Fail-Closed Model

If correctness cannot be validated:

```text
mutation rejected
```

is preferred over:

```text
unsafe mutation admission
```

The runtime intentionally prioritizes:

```text
execution correctness
```

over:

```text
availability optimism
```

---

# Scope Limitations

This repository does not claim:

- formal verification
- production hardening
- cryptographic perfection
- nation-state resistance
- operational deployment security
- external audit certification
- full production runtime guarantees

The repository demonstrates a deterministic public proof surface.

---

# Public Proof Goal

The purpose of the repository is:

```text
one command
→ observable runtime behavior
→ deterministic rejection verdicts
```

under hostile and unstable execution conditions.

---

# Architectural Goal

The long-term architectural objective is:

```text
transport instability
≠
canonical execution corruption
```

and:

```text
hostile input
must not silently cross
the canonical execution boundary
```