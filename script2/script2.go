package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robertkrimen/otto"
)

func main() {

	// 初始化随机数种子
	rand.Seed(time.Now().Unix())

	// 新建Javascript虚拟机
	vmT := otto.New()

	// 设置虚拟机中的变量a为数字18
	vmT.Set("a", 18)

	// 设置虚拟机中的变量text1为字符串
	vmT.Set("text1", "[计算结果] ")

	// 设置供虚拟中Javascript代码调用的Go语言函数
	vmT.Set("getRandomInt", func(call otto.FunctionCall) otto.Value {
		// 获取调用该函数时传入的参数
		// maxA将作为生成随机数的最大值上限
		maxA, _ := call.Argument(0).ToInteger()

		// 生成随机整数
		randomNumberT := rand.Intn(int(maxA))

		// 转换为虚拟机中可以接受的类型
		rs, _ := otto.ToValue(randomNumberT)

		// 返回该值
		return rs
	})

	// 在虚拟机中运行代码
	vmT.Run(`
		result1 = a + 2; // 计算a+2的数值

		console.log(text1 + result1); // 输出信息

		result2 = getRandomInt(20); // 调用Go语言中的函数获取20以内的随机整数
	`)

	// 从虚拟机中获取变量的值
	if valueT, errT := vmT.Get("result1"); errT == nil {
		if valueIntT, errT := valueT.ToInteger(); errT == nil {
			fmt.Printf("result1: %v\n", valueIntT)
		}
	}

	if valueT, errT := vmT.Get("result2"); errT == nil {
		if valueIntT, errT := valueT.ToInteger(); errT == nil {
			fmt.Printf("result2: %v\n", valueIntT)
		}
	}

	// 继续调用虚拟机来计算表达式的值
	// 注意此时虚拟机中的环境（变量等）都还有效
	valueT, _ := vmT.Run("result2 * 100")
	{
		valueIntT, _ := valueT.ToInteger()

		fmt.Printf("表达式结果：%v\n", valueIntT)
	}

}
