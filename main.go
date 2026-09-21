package main

import (
	"fmt"
	"math"
	myfmt "practice_go/phase_01_core/00_myfmt"
	// Add your current exercise import here, e.g.:
	// variablesfunctions "practice_go/phase_01_core/01_variables_functions"
	// math2 "practice_go/phase_02_packages/12_math2"
)

func main() {
	fmt.Printf("%s\n", myfmt.Itoa(-1))
	fmt.Println(myfmt.Itoa(0))
	fmt.Printf("%s\n", myfmt.Itoa(10))
	fmt.Println(myfmt.Itoa(20))
	fmt.Println(myfmt.Itoa(math.MinInt))
}
