package main

import (
	"encoding/hex"
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	// 将在浏览器的控制台输出信息
	fmt.Printf("这是一个Go WebAssembly的例子。\n")

	// 获取网页DOM对象中的输入框
	var input1 = js.Global().Get("document").Call("getElementById", "input1")

	// 获取网页中的按钮
	var button1 = js.Global().Get("document").Call("getElementById", "button1")

	//设置准备绑定在按钮button1上的Go语言编写的函数
	callbackFuncT := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// 获取输入框input1中的内容
		textT := input1.Get("value").String()

		// 在控制台内输出该内容
		fmt.Printf("textT: %v\n", textT)

		// 将输入框input1中的内容变成之前内容十六进制编码后的结果
		input1.Set("value", hex.EncodeToString([]byte(textT)))

		return nil
	})

	// 将该函数绑定在按钮的点击事件上
	button1.Call("addEventListener", "click", callbackFuncT)

	// 最后程序不能退出，否则按钮点击时将提示Go/Wasm代码已经终止运行
	// 因此使用无限循环（也可以用其他方法）来使程序保持运行
	for {
		// 每次休眠100毫秒避免系统资源占用过高
		time.Sleep(100 * time.Millisecond)
	}
}
