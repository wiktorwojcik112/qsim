package main

import "fmt"

func main() {
	// Bell state

	sim := createSim(2)
	
	sim.setState(0b00)
	sim.H(0)
	sim.CN(1, 0)

	probs := sim.getProbabilities()
	for i := 0; i < len(probs); i++ {
		fmt.Printf("%02b: %f\n", i, probs[i])
	}

	res := sim.getResult()
	fmt.Printf("Result: %d\n", res)
}
