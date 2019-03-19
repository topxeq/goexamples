package main

import (
	"fmt"
	"net"
	"strings"
	"time"
	t "tools"
)

func main() {

	// 在本机的8819端口监听UDP协议的连接
	listenerT, errT := net.ListenPacket("udp", ":8819")
	if errT != nil {
		t.Printfln("监听端口时发生错误：%v\n", errT.Error())
		return
	}

	// 保证UDP监听器会被关闭
	defer listenerT.Close()

	// 声明用于接收客户端信息的缓冲区
	bufT := make([]byte, 4096)

	// 用于收到消息的计数
	messageCountT := 0

	// 循环监听数据
	for {
		// countT中是成功接收消息后实际读取到的字节数
		// addressT是客户端的地址
		countT, addressT, errT := listenerT.ReadFrom(bufT)

		if errT != nil {
			t.Printfln("接收数据时发生错误：%v", errT.Error())

			continue
		}

		// 将收到的数据清理
		messageT := strings.TrimSpace(string(bufT[:countT]))

		// 根据收到的字符串进行处理
		switch messageT {
		case "exit": // 收到“exit”则关闭连接，程序将终止

			listenerT.WriteTo([]byte(fmt.Sprintf("连接将被关闭，共收到%v条消息\n", messageCountT)), addressT)

			return
		default: // 默认是做加上当前系统时间的简单回复

			// 对收到的消息进行计数
			messageCountT++

			t.Printfln("收到：%v", messageT)

			// 生成回复字符串
			responseT := fmt.Sprintf("[%v] 已收到%v\n", time.Now(), messageT)

			// 在连接上写入回复字符串
			listenerT.WriteTo([]byte(responseT), addressT)
		}

	}
}
