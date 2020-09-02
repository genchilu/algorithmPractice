package pow

func myPow(x float64, n int) float64 {
	if(n==0) {
		return 1
	}

	if(x==0 || x==1) {
		return x
	}

	if(x==-1) {
		if (n%2==0) {
			return 1.0
		} else {
			return -1.0
		}
	}

	isNative := false
	if n < 0 {
		n = -n
		isNative = true
	}

	result := 1.0

	if(isNative) {
		for i:=0;i<n;i++ {
			result *= x
			if(result>100000) {
				return 0.0
			}
		}
	} else {
		for i:=0;i<n;i++ {
			result *= x
		}
	}

	if(isNative) {
		return 1.0/result;
	}

    return result
}