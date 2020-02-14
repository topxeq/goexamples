package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/topxeq/tk"
)

func main() {
	var errT error

	// 获取第1个命令行参数（实际上是第二个命令行参数，可执行文件名是第1个，序号为0）
	fileNameT := tk.GetParameterByIndexWithDefaultValue(os.Args, 1, "")

	// 如果命令行参数中没有指定文件名则报错
	if fileNameT == "" {
		tk.Pl("文件名不能为空")
		return
	}

	// 如果文件不存在也报错
	if !tk.IfFileExists(fileNameT) {
		tk.Pl("文件 %v 不存在。", fileNameT)
		return
	}

	// limitLineCountT限制每个分块文件的大小（行数）
	// 从命令行参数中可以用-size=100000这样的参数来设置，默认为5000000行
	limitLineCountT := tk.GetSwitchWithDefaultIntValue(os.Args, "-size=", 5000000)

	// 总行数
	lineCountT := 0

	// 分块文件数
	fileCountT := 0

	// 打开原始文件准备进行读取
	fileT, errT := os.Open(fileNameT)
	if errT != nil {
		tk.Pl("打开文件时发生错误：%v", errT)
		return
	}

	// 创建一个缓冲式读取器对象
	readerT := bufio.NewReader(fileT)

	// ifEOFT用于判断是否读到了文件末尾
	ifEOFT := false

	// 临时变量，用于存储字符串
	var tmps string

	// 反复循环从源文件中读取行，直至读到文件末尾
	// 每次读取最多limitLineCountT行，写入临时文件中，超出则继续写到下一个临时文件中
	// 临时文件名按数字进行排序，存于变量fileCountT中
	for !ifEOFT {
		// 分配指定大小的切片（可以理解为Go语言中的可变长数组）准备放置读取到的文本行
		bufT := make([]string, 0, limitLineCountT)

		fileCountT++
		tk.Pl("正在读取第%v组数据", fileCountT)

		// 临时文件名，tk.Spr函数相当于fmt.Sprintf函数
		// 本例中的临时文件名将依次为sub00000001.txt、sub00000002.txt...
		subFileNameT := tk.Spr("sub%08d.txt", fileCountT)

		// 默认将临时文件放在执行时的当前目录下
		subPathT := filepath.Join("./", subFileNameT)

		// 循环读取limitLineCountT次，试图读取limitLineCountT行文本
		for j := 0; j < limitLineCountT; j++ {
			strT, errT := readerT.ReadString('\n')
			if errT != nil {
				// 读到文件结尾时的处理
				if errT == io.EOF {
					tmps = tk.Trim(strT)
					if tmps != "" {
						bufT = append(bufT, tmps)
					}

					ifEOFT = true
				} else {
					tk.Pl("文件读取失败:%v", errT)
					fileT.Close()
					os.Exit(1)
				}
				break
			}

			tmps = tk.Trim(strT)

			// 本例中空行将被丢弃，即不处理空行（包括含有空格等空白字符的行）
			if tmps != "" {
				bufT = append(bufT, tmps)
			}
		}

		// 对读取到的最多limitLineCountT行文本进行排序
		tk.Pl("正在排序第%v组数据", fileCountT)
		sort.Sort(sort.StringSlice(bufT))

		// 保存排序后的文本到临时文件
		tk.Pl("正在保存第%v组数据到临时文件%v", fileCountT, subPathT)
		rse := tk.SaveStringListBuffered(bufT, subPathT, "\n")
		if tk.IsErrorString(rse) {
			tk.Pl("保存临时文件%v失败:%v", subPathT, tk.GetErrorString(rse))
			fileT.Close()
			os.Exit(1)
		}

		// 记录总共处理的行数
		lineCountT += len(bufT)
	}

	fileT.Close()

	tk.Pl("共读取了%v行，写入了%v个临时文件", lineCountT, fileCountT)

	// 排序写
	tk.Pl("进行多文件排序并去除重复行……")

	// 存放临时文件读取器的变量
	filesT := make([]*os.File, fileCountT)
	readersT := make([]*bufio.Reader, fileCountT)

	// 用于进行对多个文件读取的第一行进行大小比对排序的变量
	strBufT := make([]string, fileCountT)
	compareBufT := make([]int, fileCountT)
	selIndexT := 0

	// 用于保存当前写入的行，用于去除重复行
	currentLineT := ""

	// 统计整体读取的行数和写入的行数
	readCountT := 0
	writeCountT := 0

	// 打开多个临时文件用于同时读取
	for i := 1; i <= fileCountT; i++ {
		subPathT := filepath.Join("./", tk.Spr("sub%08d.txt", i))

		tk.Pl("打开临时文件%v准备读取", subPathT)
		filesT[i-1], errT = os.Open(subPathT)
		if errT != nil {
			tk.Pl("打开文件时发生错误：%v", errT)
			os.Exit(1)
		}

		readersT[i-1] = bufio.NewReader(filesT[i-1])

	}

	// 创建一个新文件用于写入最终结果，默认为当前目录下的output.txt文件
	outputFileT, errT := os.Create("./output.txt")
	if errT != nil {
		tk.Pl("创建输出文件时发生错误：%v", errT)
		os.Exit(1)
	}

	// 创建写入器
	outputWriterT := bufio.NewWriter(outputFileT)

	// 用于判断是否是写入的第一行
	// 如果不是第一行，将再写入每一行文本之前，先写入一个回车换行符
	notFirstFlagT := false

	// 循环读取并写入结果文件
	for true {

		var lineT string

		// 记录一共被关闭了多少个临时文件，表示已经有多少个临时文件被读取完毕
		var closedFileT = 0

		// 是否读到文件结尾
		var eofT bool

		// 从各个文件中都读取一行，空行将被丢弃
		for k := 0; k < fileCountT; k++ {
			if readersT[k] == nil {
				closedFileT++
				continue
			}

			// 如果某个文件对应的一行已空，则再读一行
			if strBufT[k] == "" {

				foundT := false
				for readersT[k] != nil {
					lineT, eofT, errT = tk.ReadLineFromBufioReader(readersT[k])

					if errT != nil {
						tk.Pl("从临时文件%v中读取数据时发生错误：%v", k, errT)
						os.Exit(1)
					}

					lineT = tk.Trim(lineT)

					if eofT {
						readersT[k] = nil
						filesT[k].Close()
					}

					if lineT == "" {
						continue
					}

					foundT = true
					break
				}

				if foundT {
					strBufT[k] = lineT
				}
			}
		}

		// 进行计数式比对，找出排名最靠前的一行
		var compareT int

		for ii := 0; ii < fileCountT; ii++ {
			compareBufT[ii] = 0
		}

		for ii := 0; ii < (fileCountT - 1); ii++ {
			if strBufT[ii] == "" {
				continue
			}

			for jj := ii + 1; jj < fileCountT; jj++ {
				if strBufT[jj] == "" {
					compareBufT[ii]++
					continue
				}

				compareT = strings.Compare(strBufT[ii], strBufT[jj])
				if compareT > 0 {
					compareBufT[jj]++
				} else if compareT < 0 {
					compareBufT[ii]++
				}
			}
		}

		maxT := 0
		for kk := 0; kk < fileCountT; kk++ {
			if compareBufT[kk] > maxT {
				maxT = compareBufT[kk]
				selIndexT = kk
			}
		}

		// 处理只有一个文件时的比对
		if fileCountT == 1 && strBufT[0] != "" {
			maxT = 1
			selIndexT = 0
		}

		// 如果所有行都是空行，说明已经读取完毕所有文件，将退出循环
		if maxT <= 0 {
			tk.Pl("读取缓冲区全部为空")
			break
		}

		readCountT++

		// 如果将要写入的一行与上一行一样，说明是重复行，则丢弃
		// 由此实现去除重复行的效果
		// 注意此方法仅对排序后的文本才是正确的
		if currentLineT != "" {
			if strBufT[selIndexT] == currentLineT {
				// tk.Pl("发现重复行：%v", currentLineT)
				strBufT[selIndexT] = ""
				continue
			}
		}

		currentLineT = strBufT[selIndexT]
		strBufT[selIndexT] = ""

		if notFirstFlagT {
			outputWriterT.WriteString("\r\n")
		} else {
			notFirstFlagT = true
		}

		// 将最终选出的文本行写入结果文件
		_, errT = outputWriterT.WriteString(currentLineT)
		if errT != nil {
			tk.Pl("向输出文件中写入数据时发生错误：%v", errT)
			os.Exit(1)
		}

		writeCountT++

		// 所有文件如果都已关闭，说明都已读取完，循环将终止
		if closedFileT == fileCountT {
			break
		}

	}

	// 由于使用的是bufio，即缓冲方式写入文件，注意一定要用Flush来保证在内存中的数据被确保真正写入文件中
	outputWriterT.Flush()
	outputFileT.Close()

	tk.Pl("处理完毕（共写入%v行），按q键加回车退出……", writeCountT)
}
