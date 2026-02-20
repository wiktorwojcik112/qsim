package engine

import "math/cmplx"
import "math"

type state struct {
	val     []complex128
	nQubits int
	size    int
}

func newState(nQubits int) *state {
	size := int(math.Pow(2, float64(nQubits)))

	s := state{
		val:     make([]complex128, size),
		nQubits: nQubits,
		size:    size,
	}

	s.val[0] = 1 // |0...0> by default

	return &s
}

func ampOf(i int, s *state) complex128 {
	if i < 0 || i >= s.size {
		panic("Invalid index")
	}

	return complex(cmplx.Abs(s.val[i]), 0)
}

func probOf(i int, s *state) float64 {
	// Probability of state i

	if i < 0 || i >= s.size {
		panic("Invalid index")
	}

	amp := cmplx.Abs(s.val[i])
	return amp * amp
}

func qubitProbOf(q int, s *state) float64 {
	// Probability that qubit (q) is 1.

	var prob float64
	mask := 1 << q
	for i := 0; i < s.size; i++ {
		if (i & mask) != 0 {
			prob += probOf(i, s)			
		}
	}

	return prob
}

func normalize(s *state) {
	// Normalize so sum(|a_i|^2) equals 1.
	// This is meant to prevent drifts caused by the fact that float64 has finite
	// precision so after many steps error could cumulate and cause invalid results.

	// The normalization value: sqrt(sum(Re_i^2 + Im_i^2))
	var amp_sum float64
	for i := range s.size {
		amp_sum += probOf(i, s)		
	}

	norm := math.Sqrt(amp_sum)

	for i := range s.size {
		s.val[i] /= complex(norm, 0)
	}
}
