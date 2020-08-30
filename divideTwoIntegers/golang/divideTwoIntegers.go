package divideTwoIntegers


func divide(dividend int, divisor int) int {
	if (dividend == 0 || divisor == 0) {
		return 0
	}

	if (dividend == -2147483648 && divisor == -1) {
		return 2147483647
	}

	shouldNegative  := false
	if(dividend<0) {
		shouldNegative = !shouldNegative
		dividend = -dividend
	}
	
	if(divisor<0) {
		shouldNegative = !shouldNegative
		divisor = -divisor
	}

	count :=0;

	if(divisor == 1) {
		count = dividend
	} else if(divisor == 2) {
		count = dividend>>1
	} else {
		for dividend >= divisor {
			dividend -=divisor
			count++
		}
	}

	if(shouldNegative) {
		count = -count;
	}

	return count
}
