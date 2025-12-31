package algorithms

import "math"

func TwoCristalBalls(breaks []bool) int {

	found := false
	foundIdx := -1
	size := len(breaks)
	lastJump := 0
	i := int(math.Floor(math.Sqrt(float64(size))))
	for i < size {

		if breaks[i] {
			for lastJump < i {

				if breaks[lastJump] {
					found = true
					foundIdx = lastJump
					break
				}

				lastJump++
			}

			if found {
				break
			}
		} else {
			lastJump = i
			i += int(math.Floor(math.Sqrt(float64(size))))
		}
	}

	return foundIdx
}
