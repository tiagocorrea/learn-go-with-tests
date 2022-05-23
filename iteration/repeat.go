package iteration

const repeatCount = 5

// Repeat given `character` by many `count`.
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
