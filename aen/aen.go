package aen

func AverageOfEvenNumber(array []int) int {
	var sum int
	var countOfEvenNumbers int = 0
	for _,number := range array {
		if number % 2 == 0 {
			sum += number
			countOfEvenNumbers += 1
		}
	}
	return sum / countOfEvenNumbers
}