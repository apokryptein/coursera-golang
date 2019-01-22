// kinematics.go
// Author: Tomlinson

package main

import (
	"fmt"
	"math"
)


func main() {
	var accel, iVelo, iDis, secs float64

	fmt.Printf("\nAcceleration: ")
	fmt.Scan(&accel)
	fmt.Printf("Initial Velocity: ")
	fmt.Scan(&iVelo)
	fmt.Printf("Initial Displacement: ")
	fmt.Scan(&iDis)
	fmt.Printf("Seconds: ")
	fmt.Scan(&secs)

	fn := GenDisplaceFn(accel, iVelo, iDis)

	fmt.Printf("Displacement after %.02f seconds: %.02f\n", secs, fn(secs))
}

func GenDisplaceFn(accel, iVelo, iDis float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5 * accel * math.Pow(t, 2) + (iVelo * t) + iDis
	}
	return fn
}

