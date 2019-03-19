package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	// 连接服务器，127.0.0.1是特殊的IP地址，表示本机
	connectionT, errT := net.Dial("tcp", "127.0.0.1:8818")

	if errT != nil {
		fmt.Printf("连接服务器时发生错误：%v\n", errT.Error())
		return
	}

	// 确保连接会被关闭
	defer connectionT.Close()

	// inputT用于放置命令行输入的字符串
	var inputT string

	// 新建一个从标准输入读取信息的bufio.Reader类型变量
	stdInputT := bufio.NewReader(os.Stdin)

	// 重复循环从命令行读取字符串并发送到服务器，然后等待响应
	for {
		// 从命令行接收一行字符串
		inputT, errT = stdInputT.ReadString('\n')

		// 将字符串写入到与服务器的连接，注意最后一定要加上“\n”
		fmt.Fprintf(connectionT, "%v\n", strings.TrimSpace(inputT))

		// 接收服务器的响应信息
		responseT, errT := bufio.NewReader(connectionT).ReadString('\n')

		if errT != nil {

			// 如果服务器已关闭连接，将会受到io.EOF的错误，此时应退出goroutine
			if errT == io.EOF {
				fmt.Printf("服务器已关闭连接：%v\n", errT.Error())
				return
			}

			// 遇到其他错误则输出信息后中止循环
			fmt.Printf("从服务器接受响应时发生错误：%v\n", errT.Error())

			break
		}

		// 将响应信息去掉首尾空白字符
		responseT = strings.TrimSpace(responseT)

		// 在命令行界面上输出从服务器收到的响应信息
		fmt.Printf("服务器响应：%v\n", responseT)
	}

}
