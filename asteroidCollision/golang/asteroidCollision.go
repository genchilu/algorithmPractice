package asteroidCollision

func asteroidCollision(asteroids []int) []int {

	ans := []int{}

	for _, v := range asteroids {
		add := true
		for len(ans) > 0 && ans[len(ans)-1] > 0 && v < 0 {
			if ans[len(ans)-1] == -v {
				ans = ans[:len(ans)-1]
				add = false
				break
			} else if ans[len(ans)-1] < -v {
				ans = ans[:len(ans)-1]
			} else {
				add = false
				break
			}
		}

		if add {
			ans = append(ans, v)
		}
	}

	return ans
}
