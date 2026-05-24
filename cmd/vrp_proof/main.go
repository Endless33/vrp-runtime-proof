package main

import (
	"fmt"

	"github.com/Endless33/vrp-runtime-proof/internal/proof"
)

func main() {
	fmt.Println("=== VRP CONTINUITY RUNTIME PROOF ===")
	fmt.Println()

	report := proof.RunProof()

	for _, result := range report.Results {
		status := "FAILED"

		if result.Passed {
			status = "PASSED"
		}

		fmt.Printf(
			"%-30s %s\n",
			result.Name,
			status,
		)

		fmt.Printf(
			"reason=%s\n\n",
			result.Reason,
		)
	}

	fmt.Printf("accepted=%d\n", report.Accepted)
	fmt.Printf("rejected=%d\n", report.Rejected)
	fmt.Printf("replay_rejected=%d\n", report.ReplayRejected)
	fmt.Printf("stale_epoch_rejected=%d\n", report.StaleEpochRejected)
	fmt.Printf("non_authority_rejected=%d\n", report.NonAuthorityRejected)
	fmt.Printf("session_preserved=%v\n", report.SessionPreserved)
	fmt.Printf("recovery_preserved=%v\n", report.RecoveryPreserved)

	fmt.Println()
	fmt.Printf("VERDICT=%s\n", report.Verdict)
}