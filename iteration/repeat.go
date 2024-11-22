package iteration

func Repeat(a string, n int) string {
	var repeated string

	if n < 1 {
		repeated = a
	}
	for i := 0; i < n; i += 1 {
		repeated += a
	}
	return repeated
}
