package main

import "fmt"

func sumOfMultiples(number int) int {
	sum := 0
	for i := 1; i < number; i++ {
		of3 := i % 3
		of5 := i % 5
		if of3 == 0 && of5 == 0 {
			sum += i
			continue
		}
		if of3 == 0 {
			sum += i
		}
		if of5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(sumOfMultiples(10))
}
