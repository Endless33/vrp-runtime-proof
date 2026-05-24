# VRP Attack Model

## Purpose

The VRP attack proof harness exists to validate deterministic rejection behavior under hostile execution input.

The goal is not to prove invulnerability.

The goal is to validate that hostile or invalid execution input does not silently cross the canonical execution boundary.

---

# Attack Surface Validation

The attack harness currently validates:

- MITM tampering rejection
- replay attack rejection
- stale epoch rejection
- stale authority rejection
- fake recovery snapshot rejection

Command:

```bash
go run ./cmd/vrp_attack_proof
```

Expected verdict:

```text
VERDICT=VRP_ATTACK_SURFACE_PROOF_VALID
```

---

# MITM Tampering

The runtime validates authenticated mutation boundaries.

If payload contents are modified after authentication:

```text
payload modified
→ authentication mismatch
→ canonical rejection
```

Expected verdict:

```text
AUTHENTICATION_FAILED
```

---

# Replay Attacks

Previously admitted mutations may not commit more than once.

Expected behavior:

```text
captured mutation replayed
→ replay rejected
```

Expected verdict:

```text
REPLAY_REJECTED
```

---

# Stale Epoch Injection

Older execution state may not overwrite newer canonical state.

Expected behavior:

```text
old epoch submitted
→ stale epoch rejected
```

Expected verdict:

```text
STALE_EPOCH_REJECTED
```

---

# Stale Authority Injection

Previous authorities may not regain canonical execution rights automatically.

Expected behavior:

```text
non-canonical authority submits mutation
→ rejected
```

Expected verdict:

```text
NON_AUTHORITY_REJECTED
```

---

# Fake Recovery Snapshot Injection

Recovery paths must validate canonical recovery lineage.

Expected behavior:

```text
invalid recovery snapshot submitted
→ recovery rejected
```

Expected verdict:

```text
RECOVERY_SNAPSHOT_REJECTED
```

---

# Important Scope Limitation

This repository is not a production security audit.

It is a deterministic runtime proof harness.

The repository does not claim:

- formal verification
- production hardening
- external pentest certification
- operational deployment security
- cryptographic perfection

The repository demonstrates observable runtime rejection behavior under controlled hostile input conditions.