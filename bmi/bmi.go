package main

import (
	"fmt"
	"math"
	"os"
	t "tools"
)

func main() {
	args := os.Args

	var W float64
	var H float64

	fmt.Sscanf(args[1], "%f", &W)
	fmt.Sscanf(args[2], "%f", &H)

	t.Printfln("体重: %.2f", W)
	t.Printfln("身高: %.2f", H)

	BMI := W / math.Pow(H, 2)

	t.Printfln("BMI: %.2f", BMI)

	if BMI < 18.5 {
		t.Printfln("偏瘦")
	} else if (18.5 <= BMI) && (BMI < 24) {
		t.Printfln("正常")
	} else if 24 <= BMI && BMI < 28 {
		t.Printfln("偏胖")
	} else if 28 <= BMI && BMI < 30 {
		t.Printfln("肥胖")
	} else if BMI >= 30 {
		t.Printfln("重度肥胖")
	}
}
