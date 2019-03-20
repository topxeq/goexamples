package main

import "fmt"

func getClosure() func(c int) int {
	totalCount := 0

	return func(d int) int {
		totalCount += d

		return totalCount
	}
}

func main() {

	f1 := getClosure()

	for i := 0; i < 5; i++ {

		countT := f1(3)

		fmt.Println(countT)
	}

	fmt.Println("-----")

	f2 := getClosure()

	for i := 0; i < 5; i++ {

		countT := f2(3)

		fmt.Println(countT)
	}

}
