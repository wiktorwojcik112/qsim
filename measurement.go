package main

import "math/rand/v2"

func measure(s *state) {
	// Measures the probabilities for each state.
	random := rand.Float64()
	var run_sum float64

	res := 0
	for i := 0; i < s.size; i++ {
		run_sum += probOf(i, s)
		if random <= run_sum {
			res = i
			break
		}
	}

	for i := 0; i < s.size; i++ {
		if i == res {
			s.val[i] = 1
		} else {
			s.val[i] = 0
		}
	}
}

func measureQubit(q int, s *state) {
	// Measures the probability that the qubit (q) is going to be 1.
	prob := qubitProbOf(q, s)
	random := rand.Float64()

	res := 0
	if random > prob {
		// The result is 0
		res = 0
	} else {
		// The result is 1
		res = 1
	}

	mask := 1 << q
	for i := 0; i < s.size; i++ {
		if (i & mask) == 0 && res != 0 || (i & mask) != 0 && res == 0 {
			s.val[i] = 0	
		}
	}

	normalize(s)
}

func getResult(s *state) int {
	// Measures the probabilities for each state.
	random := rand.Float64()
	var run_sum float64

	for i := 0; i < s.size; i++ {
		run_sum += probOf(i, s)
		if random <= run_sum {
			return i
		}
	}

	return 0 // Should never happen
}
