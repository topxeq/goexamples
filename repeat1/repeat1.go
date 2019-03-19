package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", time.Now())

		time.Sleep(time.Second * 3)
	}
}
