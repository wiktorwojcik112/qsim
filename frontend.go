package main

type simulation struct {
	s *state
}

func createSim(nQubits int) simulation {
	s := newState(nQubits)
	sim :=  simulation{s: s}

	return sim
}

func (sim *simulation) setState(idx int) {
	for i := 0; i < sim.s.size; i++ {
		if i == idx {
			sim.s.val[i] = 1
		} else {
			sim.s.val[i] = 0
		}
	}
}

func (sim *simulation) getProbabilities() []float64 {
	probs := make([]float64, sim.s.size)
	for i := 0; i < sim.s.size; i++ {
		probs[i] = probOf(i, sim.s)	
	}

	return probs
}

func (sim *simulation) X(q int) {
	applySingleQubitGate(q, g_pauli_x(), sim.s)
}

func (sim *simulation) Y(q int) {
	applySingleQubitGate(q, g_pauli_y(), sim.s)
}

func (sim *simulation) Z(q int) {
	applySingleQubitGate(q, g_pauli_z(), sim.s)
}

func (sim *simulation) H(q int) {
	applySingleQubitGate(q, g_hadamard(), sim.s)
}

func (sim *simulation) S(q int) {
	applySingleQubitGate(q, g_s_gate(), sim.s)
}

func (sim *simulation) T(q int) {
	applySingleQubitGate(q, g_t_gate(), sim.s)
}

func (sim *simulation) R_x(q int, theta float64) {
	applySingleQubitGate(q, g_r_x(theta), sim.s)
}

func (sim *simulation) R_y(q int, theta float64) {
	applySingleQubitGate(q, g_r_y(theta), sim.s)
}

func (sim *simulation) R_z(q int, theta float64) {
	applySingleQubitGate(q, g_r_z(theta), sim.s)
}

func (sim *simulation) CN(qt int, qc int) {
	applyControlledGate(qt, qc, g_pauli_x(), sim.s)
}

func (sim *simulation) measureAll() {
	measure(sim.s)
}

func (sim *simulation) measure(q int) {
	measureQubit(q, sim.s)
}

func (sim *simulation) getResult() int {
	return getResult(sim.s)
}






