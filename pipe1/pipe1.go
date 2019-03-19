package main

import (
	"bufio"
	"io"
	"os/exec"
	"tools"
)

// 负责运行程序或命令的函数
// 第一个参数是要运行的程序或命令的路径
// 如果有命令行参数的话，从后面的可变参数argsA中读取
func runCmd(nameA string, argsA ...string) {

	// 创建执行命令的对象
	cmdT := exec.Command(nameA, argsA...)

	// 获取该对象的标准输出管道
	pipeT, errT := cmdT.StdoutPipe()

	if errT != nil {
		tools.Printfln("设置管道时发生错误：%v", errT)
		return
	}

	// 创建从管道中读取内容的bufio.Reader对象
	readerT := bufio.NewReader(pipeT)

	// 启动一个goroutine来读取
	go func() {

		// 循环每次读取一行，直至读到io.EOF或出错
		for {
			inputT, errT := readerT.ReadString('\n')

			if errT != nil {
				if errT == io.EOF {
					return
				}

				tools.Printfln("从管道中读取内容时发生错误：%v", errT)
				return
			}

			// 输出读取到的每一行信息
			tools.Printfln("%s", inputT)
		}
	}()

	// 启动（执行）命令
	errT = cmdT.Start()
	if errT != nil {
		tools.Printfln("启动程序时发生错误：%v", errT)
		return
	}

	// 等待命令执行结束
	errT = cmdT.Wait()
	if errT != nil {
		tools.Printfln("等待程序运行完毕时发生错误：%v", errT)
		return
	}

	return
}

func main() {
	// 执行命令repeat1
	// repeat1.exe必须在当前目录下或环境变量path中指定的某一个的目录中
	runCmd(`repeat1`)
}
