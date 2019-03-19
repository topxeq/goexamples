package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// connectionHandler 是处理单个连接的函数
func connectionHandler(connectionA net.Conn) {

	// 确保连接最终会被关闭
	defer connectionA.Close()

	messageCountT := 0

	// 在连接上循环接收一行一行的文本字符串并作处理
	for {
		// 从连接读取字符串，每次接收一行（以“\n”换行符为分界）
		messageT, errT := bufio.NewReader(connectionA).ReadString('\n')

		if errT != nil {
			fmt.Printf("从连接读取数据时发生错误（连接将被关闭）：%v\n", errT.Error())
			return
		}

		// 去除收到字符串的首尾空白字符（包括最后的“\n”）
		messageT = strings.TrimSpace(messageT)

		// 根据收到的字符串进行处理
		switch messageT {
		case "exit": // 收到“exit”则关闭连接，本goroutine将终止

			// 用fmt.Fprintf直接在连接上写入字符串
			fmt.Fprintf(connectionA, "连接将被关闭，共收到%v条消息\n", messageCountT)

			return
		default: // 默认是做加上当前系统时间的简单回复

			// 对收到的消息进行计数
			messageCountT++

			responseT := fmt.Sprintf("[%v] 已收到%v\n", time.Now(), messageT)

			// 在连接上写入回复字符串
			connectionA.Write([]byte(responseT))
		}
	}
}

func main() {

	// 在本机的8818端口监听TCP协议的连接
	listenerT, errT := net.Listen("tcp", ":8818")
	if errT != nil {
		fmt.Printf("监听端口时发生错误：%v\n", errT.Error())
		return
	}

	// 循环监听，接受连接并对每个连接新建一个goroutine处理
	for {
		connectionT, errT := listenerT.Accept()

		if errT != nil {
			fmt.Printf("接受连接时发生错误：%v\n", errT.Error())

			// 此时连接无效，直接继续循环而不启动goroutine
			continue
		}

		// 新建goroutine来处理连接
		go connectionHandler(connectionT)
	}
}
