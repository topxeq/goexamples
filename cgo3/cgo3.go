package main

import (
	"C"
	"fmt"

	"github.com/topxeq/txtk"
)

//export printInGo
func printInGo(value string) {
	fmt.Println(value)
}

//export getRandomInt
func getRandomInt(maxA int) int {
	return txtk.GetRandomIntLessThan(maxA)
}

// 必须要有一个主函数main，可以没有内容
func main() {
}
