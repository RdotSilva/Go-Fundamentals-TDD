package iteration

// Take a character and repeat it N number of times to create a new string of repeated characters
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}