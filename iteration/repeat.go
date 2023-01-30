package iteration

func Repeat(character string) (repeated string) {
	for counter := 0; counter < 5; counter++ {
		repeated += character
	}
	return
}
