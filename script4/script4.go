package main

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

var codeT = `

-- 输出字符串
print("开始运行")

-- Lua中定义函数
-- 接受两个参数，并返回一个结果值
function AddTwice(a, b)
	return a + b + b
end

`

func main() {

	// 新建一个Lua虚拟机
	vmT := lua.NewState()

	// 打开所有的Lua标准库
	lua.OpenLibraries(vmT)

	// 执行Lua脚本
	if errT := lua.DoString(vmT, codeT); errT != nil {
		panic(errT)
	}

	// 获取Lua中的函数
	vmT.Global("AddTwice")

	// 传入两个参数
	vmT.PushInteger(3)
	vmT.PushInteger(8)

	// 调用该函数，声明有两个参数和一个返回值
	vmT.Call(2, 1)

	// 获取返回结果（参数中是结果序号）
	result2, _ := vmT.ToInteger(1)

	fmt.Printf("返回结果为：%v", result2)

	// 从栈上弹出一个数值
	vmT.Pop(1)

}
