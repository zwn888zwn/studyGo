package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfArray := len(numbersToSum)
	sums = make([]int,lengthOfArray)  //make 或者append加数组长度

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers)==0 {
			sums = append(sums,0)
		}else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
