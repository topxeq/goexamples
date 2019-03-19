package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	t "tools"
)

const ClixVersion = "1.00"

func main() {
	args := os.Args

	argsLen := len(args)

	if argsLen < 2 {
		t.Printfln("请输入命令。")
		return
	}

	subCmd := args[1]

	switch subCmd {
	case "version":
		t.Printfln("CLIX v%v", ClixVersion)
		break

	case "getFileType":
		if argsLen < 3 {
			t.Printfln("请输入需要判断类型的文件名")
			break
		}

		s := args[2]

		typeT, errT := t.GetFileTypeByHead(s)

		if errT != nil {
			t.Printfln("获取文件类型时发生错误：%v", errT.Error())
			break
		}

		t.Printfln("文件类型是：%v", typeT)

		break

	case "sort":
		if argsLen < 3 {
			t.Printfln("请输入所需排序的数字序列，例如 1,5,6,7,2")
			break
		}

		s := args[2]

		list := strings.Split(s, ",")

		numberList := make([]float64, len(list))

		for i, v := range list {
			fmt.Sscanf(v, "%f", &numberList[i])
		}

		t.Printfln("排序之后的数字序列：%v", numberList)

		listLen := len(numberList)

		for i := 0; i < (listLen - 1); i++ {
			for j := i + 1; j < listLen; j++ {
				if numberList[i] < numberList[j] {
					numberList[i], numberList[j] = numberList[j], numberList[i]
				}
			}
		}

		t.Printfln("排序之后的数字序列：%v", numberList)

		break

	case "calbmi":
		wStr := t.GetFlag(args, "-w=")
		if wStr == "" {
			t.Printfln("请正确输入体重值")
			break
		}

		hStr := t.GetFlag(args, "-h=")
		if hStr == "" {
			t.Printfln("请正确输入身高值")
			break
		}

		var W float64
		var H float64

		fmt.Sscanf(wStr, "%f", &W)
		fmt.Sscanf(hStr, "%f", &H)

		BMI := W / math.Pow(H, 2)

		if t.FlagExists(args, "-value") {
			t.Printf("%.2f", BMI)
			break
		}

		t.Printfln("体重: %.2f", W)
		t.Printfln("身高: %.2f", H)

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

		break

	default:
		t.Printfln("无法识别的命令")
		break
	}

}
