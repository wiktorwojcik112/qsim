# qsim
This is a simple quantum computing simulator written in Go.

It has a quantum engine with support for the most common gates like Hadamard, X,
CNOT, etc.

## Example
The following code shows creation of the Bell State using qsim.

```go
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
```

## Roadmap
[x] Quantum Engine
[] Package organization
[] Tests
[] QASM support
