package engine

func applySingleQubitGate(q int, g [2][2]complex128, s *state) {
	// First, we need to iterate over all pairs which are states which differ
	// only on the selected qubit (q). This is more efficient than multiplying the
	// entire state by a huge 2^n x 2^n matrix.

	mask := 1 << q
	for i := 0; i < s.size; i++ {
		if (i & mask) == 0 {
			low_idx := i
			high_idx := i | mask

			s0 := s.val[low_idx]
			s1 := s.val[high_idx]

			// Matrix multiplication of the pair with the gate (g)
			// |s_i0'| = |g_00 g_01| |s_i0|
			// |s_i1'|   |g_10 g_11| |s_i1|

			// s_i0' = (g_00 * s_i0) + (g_01 * s_i1)
			n_s0 := (g[0][0] * s0) + (g[0][1] * s1)
			// s_i1' = (g_10 * s_i0) + (g_11 * s_i1)
			n_s1 := (g[1][0] * s0) + (g[1][1] * s1)

			s.val[low_idx] = n_s0
			s.val[high_idx] = n_s1
		}
	}

	normalize(s)
}

func applyControlledGate(qt int, qc int, g [2][2]complex128, s *state) {
	if qt == qc {
		panic("Control qubit needs to be different from the target qubit")
	}

	// Similar as the single qubit gate one, but this time the gate is only
	// applied to target qubit (qt) if the control qubit (qc) is 1.

	mask_t := 1 << qt
	mask_c := 1 << qc
	for i := 0; i < s.size; i++ {
		if (i & mask_t) == 0 && (i & mask_c) != 0 {
			low_idx := i
			high_idx := i | mask_t

			s0 := s.val[low_idx]
			s1 := s.val[high_idx]

			// Matrix multiplication of the pair with the gate (g)
			// |s_i0'| = |g_00 g_01| |s_i0|
			// |s_i1'|   |g_10 g_11| |s_i1|

			// s_i0' = (g_00 * s_i0) + (g_01 * s_i1)
			n_s0 := (g[0][0] * s0) + (g[0][1] * s1)
			// s_i1' = (g_10 * s_i0) + (g_11 * s_i1)
			n_s1 := (g[1][0] * s0) + (g[1][1] * s1)

			s.val[low_idx] = n_s0
			s.val[high_idx] = n_s1
		}
	}

	normalize(s)
}
