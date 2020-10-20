package asteroidCollision

func asteroidCollision(asteroids []int) []int {

	i := 0
	for i < len(asteroids)-1 {
		if asteroids[i] > 0 && asteroids[i+1] < 0 {
			if -asteroids[i+1] > asteroids[i] {
				asteroids = append(asteroids[:i], asteroids[i+1:]...)
				if i > 0 {
					i--
				}
			} else if -asteroids[i+1] < asteroids[i] {
				if i == len(asteroids)-2 {
					asteroids = asteroids[:i+1]
				} else {
					asteroids = append(asteroids[:i+1], asteroids[i+2:]...)
				}
			} else {
				if i == 0 {
					asteroids = asteroids[2:]
				} else if i == len(asteroids)-2 {
					asteroids = asteroids[:i]
				} else {
					asteroids = append(asteroids[0:i], asteroids[i+2:]...)
				}
				if i > 0 {
					i--
				}
			}
		} else {
			i++
		}
	}

	return asteroids
}
