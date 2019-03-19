package main

import (
	"encoding/xml"
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

func (v AddressType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	stringAllT := v.State + "|" + v.City + "|" + v.Detail + "|" + v.PostalCode + "|" + v.remark

	e.EncodeElement(stringAllT, start)

	return nil
}

func (p *AddressType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var stringBufT string

	d.DecodeElement(&stringBufT, &start)

	listT := strings.Split(stringBufT, "|")

	*p = AddressType{State: listT[0], City: listT[1], Detail: listT[2], PostalCode: listT[3], remark: listT[4]}

	return nil
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
	// AddressType
	SecondAddress AddressType `xml:"secondAddress"`
	Remark        string      `xml:",comment"`
}

func main() {

	person1 := &Person{Name: "张三", ID: "111111199001013336", Age: 25, Married: false, Mobile: []string{"13222228888", "15866669999"}, height: 170, SecondAddress: AddressType{State: "中国", City: "上海", Detail: "徐汇区南京路1号", PostalCode: "210001", remark: "无"}, Remark: "信息有待完善"}

	xmlBytesT, errT := xml.MarshalIndent(person1, "", "  ")
	if errT != nil {
		t.Printfln("XML缩进编码时发生错误：%v", errT)
		return
	}

	t.Printfln("XML字符串为：%v", string(xmlBytesT))

	t.Printfln("\n---分隔线---\n")

	person2 := &Person{}

	errT = xml.Unmarshal(xmlBytesT, person2)
	if errT != nil {
		t.Printfln("XML解码时发生错误：%v", errT)
		return
	}

	t.Printfln("person2：%#v", person2)

}
