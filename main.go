package main

import (
	"fmt"
	myfmt "practice_go/phase_01_core/00_myfmt"
	// Add your current exercise import here, e.g.:
	// variablesfunctions "practice_go/phase_01_core/01_variables_functions"
	// math2 "practice_go/phase_02_packages/12_math2"
)

func main() {
	test_cases := []string{"", "123", "12x", "-", "89", "+20", "4 5"}

	for i := range test_cases {
		val, err := myfmt.Atoi(test_cases[i])

		fmt.Printf("input: %s, value: %d, error_response: %s \n", test_cases[i], val, err)
	}
}
