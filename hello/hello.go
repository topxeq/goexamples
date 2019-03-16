package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	dir, _ := os.Executable()
	fmt.Println(dir)

}
