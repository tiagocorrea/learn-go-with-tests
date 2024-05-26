package iteration

func Repeat(character string, N int) string {
	var repeated string
	for i := 0; i < N; i++ {
		repeated += character
	}
	return repeated
}
