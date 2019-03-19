package main

import (
	"encoding/xml"
	"os"
	"strings"
	t "tools"
)

// AddressType 表示个人住址的数据结构
type AddressType struct {
	State      string
	City       string
	Detail     string
	PostalCode string `xml:"postalCode,attr"`
	remark     string
}

// Person 准备进行XML序列化的表示个人信息的数据结构类型
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	ID      string   `xml:"id,attr"`
	Age     int      `xml:"age"`
	Married bool
	Phone   string   `xml:"phone,omitempty"`
	Mobile  []string `xml:"mobiles>mobile"`
	height  float64  `xml:"height"`
	AddressType
	SecondAddress AddressType `xml:"secondAddress"`
	Remark        string      `xml:",comment"`
}

func main() {

	person1 := &Person{Name: "张三", ID: "111111199001013336", Age: 25, Married: false, Mobile: []string{"13222228888", "15866669999"}, height: 170, AddressType: AddressType{State: "中国", City: "北京", Detail: "海淀区中关村1号", PostalCode: "100099", remark: "路口右转"}, SecondAddress: AddressType{State: "中国", City: "上海", Detail: "徐汇区南京路1号", PostalCode: "210001", remark: "无"}, Remark: "信息有待完善"}

	var strT strings.Builder

	encoderT := xml.NewEncoder(&strT)

	errT := encoderT.Encode(person1)

	if errT != nil {
		t.Printfln("XML编码时发生错误：%v", errT.Error())
		return
	}

	t.Printfln("XML字符串为：%#v", strT.String())

	t.Printfln("\n---分隔线---\n")

	outputT, errT := xml.MarshalIndent(person1, "  ", "    ")
	if errT != nil {
		t.Printfln("XML缩进编码时发生错误：%v", errT)
		return
	}

	os.Stdout.Write(outputT)

	t.SaveStringToFile(string(outputT), `c:\test\person1.xml`)

}
