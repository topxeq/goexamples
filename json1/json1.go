package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	t "tools"
)

// AddressType 表示个人住址的数据结构
type AddressType struct {
	State      string
	City       string
	Detail     string
	PostalCode string
	remark     string
}

// Person 表示个人信息的数据结构类型
type Person struct {
	Name    string
	ID      string
	Age     int
	Married bool
	Phone   string
	Mobile  []string
	height  float64
	AddressType
	SecondAddress AddressType
	Remark        string
}

func main() {

	person1 := new(Person)

	fileT, errT := os.Open(`c:\test\person1.json`)

	if errT != nil {
		t.Printfln("打开JSON文件时发生错误：%v", errT.Error())
		fileT.Close()
		return
	}

	decoderT := json.NewDecoder(fileT)

	errT = decoderT.Decode(person1)

	fileT.Close()

	if errT != nil {
		t.Printfln("JSON解码时发生错误：%v", errT.Error())
		return
	}

	t.Printfln("person1：%#v", person1)

	t.Printfln("\n---分隔线---\n")

	person2 := &Person{}

	bytesT, errT := ioutil.ReadFile(`c:\test\person1.json`)

	if errT != nil {
		t.Printfln("再次读取JSON文件时发生错误：%v", errT.Error())
		return
	}

	errT = json.Unmarshal(bytesT, person2)
	if errT != nil {
		t.Printfln("第二次JSON解码时发生错误：%v", errT)
		return
	}

	t.Printfln("person2：%#v", person2)
}
