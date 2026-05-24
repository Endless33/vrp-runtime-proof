# VRP Runtime Proof

One-command public proof harness for VRP runtime continuity behavior.

This repository does not expose the private VRP core implementation.

It contains a minimal public proof surface that demonstrates the execution-correctness model behind VRP continuity semantics.

The goal is not to benchmark transport speed.

The goal is to show that hostile or invalid execution input does not cross the canonical execution boundary.

---

# Runtime Properties

The public proof harness demonstrates:

- replay-safe mutation admission
- stale epoch rejection
- authority enforcement
- deterministic commit admission
- session continuity preservation
- recovery snapshot validation
- fail-closed rejection behavior
- transport-independent execution continuity

---

# Quick Start

## Clone Repository

```bash
git clone https://github.com/Endless33/vrp-runtime-proof
```

---

## Enter Repository

```bash
cd vrp-runtime-proof
```

---

# Runtime Continuity Proof

## Run

```bash
go run ./cmd/vrp_proof
```

---

## Expected Verdict

```text
VERDICT=VRP_RUNTIME_CONTINUITY_PROOF_VALID
```

---

## Example Output

```text
=== VRP CONTINUITY RUNTIME PROOF ===

REPLAY TEST                  PASSED
STALE EPOCH TEST             PASSED
AUTHORITY TEST               PASSED
TRANSPORT REATTACH TEST      PASSED
SNAPSHOT RECOVERY TEST       PASSED
FAIL-CLOSED TEST             PASSED

VERDICT=VRP_RUNTIME_CONTINUITY_PROOF_VALID
```

---

# Attack Surface Proof

## Run

```bash
go run ./cmd/vrp_attack_proof
```

---

## Expected Verdict

```text
VERDICT=VRP_ATTACK_SURFACE_PROOF_VALID
```

---

## Example Output

```text
=== VRP ATTACK SURFACE PROOF ===

VALID AUTHENTIC FRAME             PASSED
MITM TAMPER ATTACK                PASSED
REPLAY ATTACK                     PASSED
STALE EPOCH ATTACK                PASSED
STALE AUTHORITY ATTACK            PASSED
FAKE RECOVERY SNAPSHOT ATTACK     PASSED

VERDICT=VRP_ATTACK_SURFACE_PROOF_VALID
```

---

# Core Claim

```text
Transport may fail.
Execution correctness must not.
```

---

# What This Repository Is

This repository is:

- a public runtime proof harness
- a deterministic execution validation surface
- a continuity correctness demonstration
- a replay/recovery validation environment

---

# What This Repository Is Not

This repository is not:

- a VPN benchmark
- a transport speed test
- a production relay network
- the private VRP runtime core
- the commercial integration layer

---

# Security Note

This repository intentionally exposes only a minimal public proof surface.

It does not include:

- private runtime internals
- production deployment logic
- relay infrastructure
- hardened operational security layers
- commercial integration code

---

# Architectural Model

VRP treats transport as a replaceable carrier rather than the source of execution truth.

Canonical execution admission depends on:

- authority validation
- monotonic epoch progression
- replay protection
- authenticated mutation boundaries
- deterministic recovery validation

rather than assuming transport stability.

---

# Repository Purpose

The purpose of this repository is to provide:

```text
one command
→ observable runtime behavior
→ deterministic proof verdicts
```

under unstable and hostile execution conditions.