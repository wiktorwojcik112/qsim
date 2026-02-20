package main

import "fmt"
import "qsim/engine"

func main() {
	// Bell state

	sim := engine.CreateSim(2)
	
	sim.SetState(0b00)
	sim.H(0)
	sim.CN(1, 0)

	probs := sim.GetProbabilities()
	for i := 0; i < len(probs); i++ {
		fmt.Printf("%02b: %f\n", i, probs[i])
	}

	res := sim.GetResult()
	fmt.Printf("Result: %d\n", res)
}
