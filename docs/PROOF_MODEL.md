# VRP Proof Model

## Purpose

The VRP Runtime Proof repository exists to provide a deterministic and reproducible public validation surface for VRP runtime behavior.

The goal is not to expose the private production runtime.

The goal is to demonstrate the execution-correctness model behind continuity-safe runtime semantics.

---

# Core Invariant

```text
Transport may fail.
Execution correctness must not.
```

VRP treats transport as a replaceable carrier rather than the source of execution truth.

Canonical execution admission depends on:

- authority validation
- replay protection
- monotonic epoch progression
- authenticated mutation boundaries
- deterministic recovery validation

---

# Runtime Proof

The runtime proof validates:

- replay-safe mutation admission
- stale epoch rejection
- authority enforcement
- deterministic commit boundaries
- transport reattach continuity
- snapshot recovery preservation
- fail-closed rejection behavior

Command:

```bash
go run ./cmd/vrp_proof
```

Expected verdict:

```text
VERDICT=VRP_RUNTIME_CONTINUITY_PROOF_VALID
```

---

# Attack Proof

The attack proof validates hostile-input rejection behavior.

Command:

```bash
go run ./cmd/vrp_attack_proof
```

Expected verdict:

```text
VERDICT=VRP_ATTACK_SURFACE_PROOF_VALID
```

---

# Deterministic Runtime Behavior

The repository intentionally produces deterministic runtime verdicts.

The goal is reproducibility across:

- Linux
- macOS
- Windows
- Termux
- virtualized environments
- CI runners

---

# Public Proof Surface

This repository intentionally exposes only a minimal public runtime surface.

It does not include:

- private runtime internals
- relay infrastructure
- production deployment logic
- operational security layers
- commercial integration code

---

# Architectural Goal

The repository demonstrates a runtime model where:

```text
packet arrival
≠
automatic execution truth
```

Instead:

```text
packet arrival
→ validation pipeline
→ canonical admission decision
→ execution boundary
```

This is the core continuity model behind VRP.