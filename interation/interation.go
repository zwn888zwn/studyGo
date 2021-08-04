package iteration

func Repeat(character string, times int) string {
	var repeatString string
	for i := 0;i<times;i++ {
		repeatString += character
	}
	return repeatString
}
