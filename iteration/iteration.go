package iteration

const repeatCount = 5

// Take a character and repeat it N number of times to create a new string of repeated characters
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}