package engine

type Simulation struct {
	s *state
}

func CreateSim(nQubits int) Simulation {
	s := newState(nQubits)
	sim :=  Simulation{s: s}

	return sim
}

func (sim *Simulation) SetState(idx int) {
	for i := 0; i < sim.s.size; i++ {
		if i == idx {
			sim.s.val[i] = 1
		} else {
			sim.s.val[i] = 0
		}
	}
}

func (sim *Simulation) GetProbabilities() []float64 {
	probs := make([]float64, sim.s.size)
	for i := 0; i < sim.s.size; i++ {
		probs[i] = probOf(i, sim.s)	
	}

	return probs
}

func (sim *Simulation) X(q int) {
	applySingleQubitGate(q, g_pauli_x(), sim.s)
}

func (sim *Simulation) Y(q int) {
	applySingleQubitGate(q, g_pauli_y(), sim.s)
}

func (sim *Simulation) Z(q int) {
	applySingleQubitGate(q, g_pauli_z(), sim.s)
}

func (sim *Simulation) H(q int) {
	applySingleQubitGate(q, g_hadamard(), sim.s)
}

func (sim *Simulation) S(q int) {
	applySingleQubitGate(q, g_s_gate(), sim.s)
}

func (sim *Simulation) T(q int) {
	applySingleQubitGate(q, g_t_gate(), sim.s)
}

func (sim *Simulation) R_x(q int, theta float64) {
	applySingleQubitGate(q, g_r_x(theta), sim.s)
}

func (sim *Simulation) R_y(q int, theta float64) {
	applySingleQubitGate(q, g_r_y(theta), sim.s)
}

func (sim *Simulation) R_z(q int, theta float64) {
	applySingleQubitGate(q, g_r_z(theta), sim.s)
}

func (sim *Simulation) CN(qt int, qc int) {
	applyControlledGate(qt, qc, g_pauli_x(), sim.s)
}

func (sim *Simulation) MeasureAll() {
	measure(sim.s)
}

func (sim *Simulation) Measure(q int) {
	measureQubit(q, sim.s)
}

func (sim *Simulation) GetResult() int {
	return getResult(sim.s)
}






