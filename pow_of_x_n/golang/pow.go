package pow

func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1
	case n == 1:
		return x
	case n<0:
		return 1/myPow(x, -n)
	}

	if(n&1 == 0) {
		return myPow(x*x, n>>1)
	} else {
		return myPow(x*x, n>>1)*x
	}
}