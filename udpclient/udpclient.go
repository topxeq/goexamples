package main

import (
	"net"
	"os"
	t "tools"
)

func main() {

	// 程序必须有一个命令行参数，即所需发送的字符串
	if len(os.Args) < 2 {
		t.Printfln("请输入所需发送的字符串")
		return
	}

	// 建立UDP连接
	connectionT, errT := net.Dial("udp", `localhost:8819`)
	if errT != nil {
		t.Printfln("建立UDP连接时发生错误：%v", errT)
		return
	}

	// 保证连接会被关闭
	defer connectionT.Close()

	// 向服务端发送数据
	_, errT = connectionT.Write([]byte(os.Args[1]))
	if errT != nil {
		t.Printfln("发送数据时发生错误：%v", errT)
		return
	}

	// 声明用于接收服务器响应的缓冲区
	bufT := make([]byte, 4096)

	// 读取服务器响应
	countT, errT := connectionT.Read(bufT)
	if errT != nil {
		t.Printfln("读取服务器响应时发生错误：%v", errT)
		return
	}

	t.Printfln("服务器响应：%v", string(bufT[:countT]))

}
