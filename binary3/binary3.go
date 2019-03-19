package main

import (
	"bytes"
	"encoding/gob"
	"os"
	t "tools"
)

// Customer 是表示客户信息的结构类型
type Customer struct {
	Name   string
	Age    int
	Gender string
	Mobile string
	Email  string
}

func main() {

	// 生成模拟的客户信息，共包含三条记录
	customers := make([]Customer, 3)

	customers[0] = Customer{Name: "张三", Age: 28, Gender: "男", Mobile: "1322226688", Email: "zhangsan@company1.com"}
	customers[1] = Customer{Name: "李四", Age: 24, Gender: "女", Mobile: "15766669999", Email: "lisi@company2.com"}
	customers[2] = Customer{Name: "王五", Age: 16, Gender: "男", Mobile: "19355558985", Email: "wangwu@company3.com"}

	t.Printfln("customers: %v", customers)

	// 创建ctm格式的文件用于写入客户信息记录
	file1T, errT := os.Create(`c:\test\customerInfo.ctm`)

	if errT != nil {
		t.Printfln("创建客户信息文件时发生错误：%v", errT.Error())
		return
	}

	// 创建编码器对象
	encoderT := gob.NewEncoder(file1T)

	// 写入文件头
	_, errT = file1T.Write([]byte{0x07, 0x01, 0x00, 0x08})

	if errT != nil {
		t.Printfln("创建客户信息文件时发生错误：%v", errT.Error())
		file1T.Close()
		return
	}

	// 写入记录条数（长度）
	encoderT.Encode(int64(len(customers)))

	// 循环写入所有记录
	for _, v := range customers {
		encoderT.Encode(v)
	}

	// 关闭文件
	file1T.Close()

	// 为读取而打开文件
	file2T, errT := os.Open(`c:\test\customerInfo.ctm`)

	if errT != nil {
		t.Printfln("打开客户信息文件时发生错误：%v", errT.Error())
		return
	}

	// file2T可以用defer语句来关闭
	defer file2T.Close()

	// 创建解码器
	decoderT := gob.NewDecoder(file2T)

	// 分配用于存储文件头的字节切片变量
	fileHeadT := make([]byte, 4)

	// 读取文件头
	_, errT = file2T.Read(fileHeadT)

	if errT != nil {
		t.Printfln("读取文件头时发生错误：%v", errT.Error())
		return
	}

	t.Printfln("文件头：%#v", fileHeadT)

	// 判断是否是正确的文件头
	if bytes.Compare(fileHeadT, []byte{0x07, 0x01, 0x00, 0x08}) != 0 {
		t.Printfln("该文件不是cfm格式的文件")
		return
	}

	// 读取记录条数
	var recordCountT int64

	decoderT.Decode(&recordCountT)

	t.Printfln("记录条数：%v", recordCountT)

	// 按记录条数分配相应空间
	records := make([]Customer, recordCountT)

	// 循环读取每条记录
	for i := 0; i < int(recordCountT); i++ {
		decoderT.Decode(&records[i])
	}

	t.Printfln("客户信息记录：%v", records)

}
