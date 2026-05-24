# VRP Runtime Proof

One-command public proof harness for VRP runtime continuity behavior.

This repository does not expose the private VRP core implementation.

It contains a minimal public proof harness that demonstrates the core execution-correctness model:

- replay-safe mutation admission
- stale epoch rejection
- authority enforcement
- deterministic commit boundary
- session continuity across transport observation
- snapshot recovery preservation
- fail-closed rejection behavior

## Run

```bash
go run ./cmd/vrp_proof
```

## Expected Verdict

```text
VERDICT=VRP_RUNTIME_CONTINUITY_PROOF_VALID
```

## Core Claim

```text
Transport may fail.
Execution correctness must not.
```

## What This Shows

This demo is not a VPN speed test.

It is a runtime correctness proof harness.

The goal is to show that invalid execution input does not cross the canonical commit boundary, even when transport conditions change or recovery is attempted.

## What This Does Not Expose

This repository does not include:

- private production runtime internals
- full transport implementation
- private relay logic
- commercial integration code
- security-sensitive internals

It is a public proof surface.