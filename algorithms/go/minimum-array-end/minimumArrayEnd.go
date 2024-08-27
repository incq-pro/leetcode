package minimum_array_end

func minEnd(n int, x int) int64 {
	var r = int64(x)
	var a = int64(n - 1)
	const base int64 = 1
	for i := 0; i < 64 && a != 0; i++ {
		b := base << i
		if r&b == 0 {
			if (a & base) != 0 {
				r |= b
			}
			a >>= 1
		}
	}
	return r
}

func MinEnd(n int, x int) int64 {
	return minEnd(n, x)
}
