package main

import (
	"fmt"
	"practice_go/basics"
	// Add your current exercise import here, e.g.:
	// variablesfunctions "practice_go/phase_01_core/01_variables_functions"
	// math2 "practice_go/phase_02_packages/12_math2"
)

func main() {
	// ---- YOUR SCRATCH PLAYGROUND (verifier IGNORES this file) ----
	// 1. Import the package you are practicing (see above).
	// 2. Call its funcs here, then run:  go run .
	// 3. Leave exercise code in its own folder; this file is throwaway.
	//
	// Default demo (keeps existing basics working):
	result, err := basics.Addition(3, 20)
	if err == nil {
		fmt.Println(result)
	}
	// Example for next step (uncomment when you start Phase 1):
	// fmt.Println(variablesfunctions.Demo())
}
