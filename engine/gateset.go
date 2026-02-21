package engine

import "math"
import "math/cmplx"

func g_pauli_x() [2][2]complex128 {
	return [2][2]complex128{
		{0, 1},
		{1, 0},
	}
}

func g_pauli_y() [2][2]complex128 {
	return [2][2]complex128{
		{0, -1i},
		{1i, 0},
	}
}

func g_pauli_z() [2][2]complex128 {
	return [2][2]complex128{
		{1, 0},
		{0, -1},
	}
}

func g_hadamard() [2][2]complex128 {
	var coef complex128 = complex(1/math.Sqrt(2), 0)
	return [2][2]complex128{
		{coef, coef},
		{coef, -coef},
	}
}

func g_s_gate() [2][2]complex128 {
	return [2][2]complex128{
		{1, 0},
		{0, 1i},
	}
}

func g_t_gate() [2][2]complex128 {
	var val complex128 = cmplx.Exp((math.Pi/4)*1i)
	return [2][2]complex128{
		{1, 0},
		{0, val},
	}
}

func g_r_x(theta float64) [2][2]complex128 {
	var cos_val complex128 = complex(math.Cos(theta/2), 0)
	var sin_val complex128 = complex(math.Sin(theta/2), 0)

	return [2][2]complex128{
		{cos_val, sin_val * -1i},
		{sin_val * -1i, cos_val},
	}
}

func g_r_y(theta float64) [2][2]complex128 {
	var cos_val complex128 = complex(math.Cos(theta/2), 0)
	var sin_val complex128 = complex(math.Sin(theta/2), 0)

	return [2][2]complex128{
		{cos_val, -sin_val},
		{sin_val, cos_val},
	}
}

func g_r_z(theta float64) [2][2]complex128 {
	var coef1 complex128 = cmplx.Exp(complex(0, -theta/2))
	var coef2 complex128 = cmplx.Exp(complex(0, theta/2))

	return [2][2]complex128{
		{coef1, 0},
		{0, coef2},
	}
}
