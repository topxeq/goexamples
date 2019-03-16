package main

import (
	"fmt"
)

type Person struct {
	Height int // 代表身高（单位为厘米）
	Weight int // 代表体重（单位为公斤）
}

func main() {
	genders := []string{"男", "女"}

	AdamInfo := &Person{Height: 175, Weight: 60}
	EveInfo := map[string]int{"Height": 165, "Weight": 50}

	fmt.Println("夏娃是" + genders[1] + "的。")

	if EveInfo["Weight"] > AdamInfo.Weight {
		fmt.Println("很遗憾，夏娃的体重现在是" + string(EveInfo["Weight"]) + "公斤。")
	} else if EveInfo["Weight"] == AdamInfo.Weight {
		fmt.Println("很遗憾，夏娃的体重和亚当一样。")
	} else {
		fmt.Println("重要的事儿说3遍!")

		for i := 0; i < 3; i++ {
			fmt.Print("夏娃没有亚当重，她的体重只有")
			fmt.Print(EveInfo["Weight"])
			fmt.Println("公斤。")
		}
	}
}
