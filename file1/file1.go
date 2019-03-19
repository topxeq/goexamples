package main

import (
	"io"
	"os"
	t "tools"
)

func main() {

	oldFileNameT := "c:\\test\\long.txt"
	newFileNameT := "c:\\test\\sub1\\copiedFile.txt"

	// 打开源文件
	oldFileT, errT := os.Open(oldFileNameT)

	if errT != nil {
		t.Printfln("打开源文件时发生错误：%v", errT.Error())
		return
	}

	defer oldFileT.Close()

	// 创建新文件
	newFileT, errT := os.OpenFile(newFileNameT, os.O_CREATE|os.O_RDWR, 0666)

	if errT != nil {
		t.Printfln("创建新文件时发生错误：%v", errT.Error())
		return
	}

	defer newFileT.Close()

	bufT := make([]byte, 5)

	for {

		countT, errT := oldFileT.Read(bufT)

		if errT != nil {
			if errT == io.EOF {
				break
			}

			t.Printfln("从源文件中读取数据时发生错误：%v", errT.Error())
			return
		}

		_, errT = newFileT.Write(bufT[:countT])

		if errT != nil {
			t.Printfln("将数据写入新文件时发生错误：%v", errT.Error())
			return
		}
	}

	t.Printfln("已成功拷贝文件 %v 到 %v。", oldFileNameT, newFileNameT)

}
