package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-proof/internal/proof"
)

func main() {
	fmt.Println("=== VRP ATTACK SURFACE PROOF ===")
	fmt.Println("Invariant: hostile input must not cross the canonical execution boundary")
	fmt.Println()

	report := proof.RunAttackProof()

	for _, result := range report.Results {
		status := "FAILED"

		if result.Passed {
			status = "PASSED"
		}

		fmt.Printf("%-38s %s\n", result.Name, status)
		fmt.Printf("verdict=%s\n", result.Verdict)
		fmt.Printf("reason=%s\n\n", result.Reason)
	}

	fmt.Println("=== ATTACK SUMMARY ===")
	fmt.Printf("accepted=%d\n", report.Accepted)
	fmt.Printf("rejected=%d\n", report.Rejected)
	fmt.Printf("tamper_rejected=%d\n", report.TamperRejected)
	fmt.Printf("replay_rejected=%d\n", report.ReplayRejected)
	fmt.Printf("stale_epoch_rejected=%d\n", report.StaleEpochRejected)
	fmt.Printf("stale_authority_rejected=%d\n", report.StaleAuthorityRejected)
	fmt.Printf("fake_recovery_rejected=%d\n", report.FakeRecoveryRejected)

	fmt.Println()
	fmt.Printf("VERDICT=%s\n", report.Verdict)
}