package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dop251/goja"
)

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())

	// 新建Javascript虚拟机
	vmT := goja.New()

	// 设置虚拟机中的变量a为数字18
	vmT.Set("a", vmT.ToValue(18))

	// 设置虚拟机中的变量text1为字符串
	vmT.Set("text1", vmT.ToValue("[计算结果] "))

	// 设置供虚拟中Javascript代码调用的Go语言函数
	// 用于生成随机整数
	vmT.Set("getRandomInt", func(call goja.FunctionCall) goja.Value {
		// 获取调用该函数时传入的参数
		// maxA将作为生成随机数的最大值上限
		maxA := call.Argument(0).ToInteger()

		// 生成随机整数
		randomNumberT := rand.Intn(int(maxA))

		// 转换为虚拟机中可以接受的类型
		rs := vmT.ToValue(randomNumberT)

		// 返回该值
		return rs
	})

	// 设置供虚拟中Javascript代码调用的Go语言函数
	// 用于输出信息，代替console.log的功能
	vmT.Set("goPrint", func(call goja.FunctionCall) goja.Value {
		// 获取调用该函数时传入的参数
		// 用于输出信息
		strT := call.Argument(0).ToString().String()

		// 用fmt.Printf输出信息并加上换行符
		fmt.Printf(strT + "\n")

		return nil
	})

	// 运行Javascript脚本
	_, errT := vmT.RunString(`
		// 将console.log函数指向Go语言代码中定义的goPrint函数
		console = { 
			log: goPrint
		};

		result1 = a + 2; // 计算a+2的数值

		console.log(text1 + result1); // 输出信息

		result2 = getRandomInt(20); // 调用Go语言中的函数获取20以内的随机整数
	`)

	if errT != nil {
		panic(errT)
	}

	// 从虚拟机中获取变量并输出其中的值
	fmt.Printf("result1: %v\n", vmT.Get("result1").ToInteger())

	fmt.Printf("result2: %v\n", vmT.Get("result2").ToInteger())

	// 继续调用虚拟机来计算表达式的值
	// 注意此时虚拟机中的环境（变量等）都还有效
	valueT, _ := vmT.RunString("result2 * 100")
	{
		valueIntT, _ := valueT.Export().(int64)

		fmt.Printf("表达式结果：%v\n", valueIntT)
	}
}
