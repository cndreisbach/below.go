package below

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func Max(a ...int) int {
	max := -(int(^uint(0)>>1) - 1) // smallest int
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
