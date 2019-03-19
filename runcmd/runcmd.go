package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"tools"
)

func main() {
	// 分配用于存放调用程序输出信息的缓冲区
	var outBufT bytes.Buffer

	// 分配用于存放调用程序运行发生错误时输出错误信息的缓冲区
	var errBufT bytes.Buffer

	// 创建执行findstr命令的对象，第二个参数开始是命令行参数
	// 一个命令行参数作为exec.Command函数的一个参数
	cmdT := exec.Command(`c:\Windows\System32\findstr.exe`, `package`, `c:\goprjs\src\tools\*.go`)

	// 把调用程序的标准输出指向outBufT
	// 标准输出一般指的是程序正常执行时输出的信息
	cmdT.Stdout = &outBufT

	// 把调用程序的标准错误输出指向errBufT
	// 标准错误输出一般指的是程序执行出现异常时输出的信息
	cmdT.Stderr = &errBufT

	// 执行该命令（或程序）
	errT := cmdT.Run()

	// 在Windows下输出信息需要转换编码
	if errT != nil {
		fmt.Printf("运行命令时发生错误：%v\n", errT.Error())
		fmt.Printf("错误信息：\n%v\n", string(tools.ConvertBytesFromGB18030ToUTF8(errBufT.Bytes())))
	} else {
		fmt.Printf("命令输出1：\n%v\n", string(tools.ConvertBytesFromGB18030ToUTF8(outBufT.Bytes())))
	}

	// 重新创建执行wc命令的对象，该命令必须在系统可找到的路径中
	// 即在环境变量PATH中存在该文件所在的目录
	cmdT = exec.Command(`wc`)

	// 指定该命令的标准输入为三行字符串
	cmdT.Stdin = strings.NewReader(`第一行。
	This is a good example.
	最后一行。
	`)

	// 重置标准输出和标准错误输出的缓冲区
	outBufT.Reset()
	errBufT.Reset()

	// 新建命令执行对象后，需要再次设置标准输出和标准错误输出缓冲区
	cmdT.Stdout = &outBufT
	cmdT.Stderr = &errBufT

	// 执行wc命令
	errT = cmdT.Run()

	if errT != nil {
		fmt.Printf("运行命令时发生错误：%v\n", errT.Error())
		fmt.Printf("错误信息：\n%v\n", string(tools.ConvertBytesFromGB18030ToUTF8(errBufT.Bytes())))
	} else {
		fmt.Printf("命令输出2：\n%v\n", string(tools.ConvertBytesFromGB18030ToUTF8(outBufT.Bytes())))
	}

}
