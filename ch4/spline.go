package ch4

func NaturalSpline(xs []float64, ys []float64) [][]float64 {
	n := len(xs)-1

	h := make([]float64, n)
	for i := 0; i < n; i++ {
		h[i] = xs[i+1] - xs[i]
	}

	// a is 1-indexed, skipped a[0]
	alpha := make([]float64, n)
	for i := 1; i < n; i++ {
		alpha[i] = 3.0/h[i]*(ys[i+1]-ys[i])-3.0/(h[i-1])*(ys[i]-ys[i-1])
	}

	l := make([]float64, n+1)
	u := make([]float64, n+1)
	z := make([]float64, n+1)
	
	l[0] = 1.0
	u[0] = 0.0
	z[0] = 0.0

	for i:= 1; i < n; i++ {
		l[i] = 2.0*(xs[i+1]-xs[i-1])-h[i-1]*u[i-1]
		u[i] = h[i]/l[i]
		z[i] = (alpha[i]-h[i-1]*z[i-1])/l[i]
	}

	l[n] = 1.0
	z[n] = 0.0
	b := make([]float64, n)
	c := make([]float64, n+1)
	d := make([]float64, n)
	c[n] = 0.0
	for j := n-1; j >= 0; j-- {
		c[j] = z[j]-u[j]*c[j+1]
		b[j] = (ys[j+1]-ys[j])/h[j] - h[j]*(c[j+1]+2.0*c[j])/3.0
		d[j] = (c[j+1]-c[j])/(3.0*h[j])
	}

	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = ys[i]
	}
	return [][]float64{a, b, c, d}
}