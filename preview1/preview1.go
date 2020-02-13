package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/topxeq/tk"
)

func main() {

	// 获取第1个命令行参数（实际上是第二个命令行参数，可执行文件名是第1个，序号为0）
	fileNameT := tk.GetParameterByIndexWithDefaultValue(os.Args, 1, "")

	if fileNameT == "" {
		tk.Pl("文件名不能为空")

		// 等待用户输入（可以输入任意字符，回车键结束）
		tk.GetUserInput("按回车键退出……")
		return
	}

	// 打开文件
	fileT, errT := os.Open(fileNameT)
	if errT != nil {
		tk.Pl("打开文件%v时发生错误：%v。", fileNameT, errT)
		tk.Pl("请按回车键结束处理……")
		tk.GetUserInput("按回车键退出……")
		return
	}

	// 保证关闭文件（在函数退出时会执行该条语句）
	defer fileT.Close()

	// 创建读取文件的缓冲式读取器
	readerT := bufio.NewReader(fileT)

	// 记录总行数
	countT := 0

	// 记录总字符数
	totalLenT := 0

	// 循环读取文件中的每一行
	for true {
		strT, errT := readerT.ReadString('\n')

		// 如果出现错误则中止循环
		if errT != nil {
			// 错误有可能是到达文件末尾，此时正常终止循环
			if errT == io.EOF {
				strT = strings.TrimRight(strT, "\r\n")
				if countT < 100 {
					tk.Pl("%v: %v", countT+1, strT)
				}

				totalLenT += len(strT)
				countT++
			}
			break
		}

		// 去除行尾可能存在的\r字符（Windows中的文本文件一般的行尾结束符是连续的\r\n两个字符）
		strT = strings.TrimRight(strT, "\r\n")

		// 100行以内会输出预览
		if countT < 100 {
			tk.Pl("%v: %v", countT+1, strT)
		}

		// 增加总字符数和总行数
		totalLenT += len(strT)
		countT++
	}

	tk.Pl("共%v行，平均每行%v个字符。", countT, totalLenT/countT)
	tk.GetUserInput("按回车键退出……")
	return

}
