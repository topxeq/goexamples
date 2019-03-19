package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	t "tools"
)

type Person struct {
	Name   string
	Age    int
	Gender string
	Height float64
	Weight float64
}

func main() {

	n1 := 64
	f1 := 12.8
	b1 := false
	s1 := "abc123"
	person1 := Person{Name: "张三", Age: 28, Gender: "男", Height: 170, Weight: 60}

	bufT := new(bytes.Buffer)

	encoderT := gob.NewEncoder(bufT)

	errT := encoderT.Encode(n1)

	if errT != nil {
		log.Fatalf("写入数据n1时发生错误：%v", errT.Error())
	}

	encoderT.Encode(f1)
	encoderT.Encode(b1)
	encoderT.Encode(s1)
	encoderT.Encode(person1)

	t.Printfln("bufT中内容：%#v", bufT.Bytes())

	ioutil.WriteFile(`c:\test\binaryFile2.bin`, bufT.Bytes(), 0666)

	bytesT, errT := ioutil.ReadFile(`c:\test\binaryFile2.bin`)

	if errT != nil {
		log.Fatalf("从文件中读取数据时发生错误：%v", errT.Error())
	}

	var n2 int
	var f2 float64
	var b2 bool
	var s2 string
	var person2 Person

	newBufT := bytes.NewBuffer(bytesT)

	decoderT := gob.NewDecoder(newBufT)

	errT = decoderT.Decode(&n2)

	if errT != nil {
		log.Fatalf("读入数据n2时发生错误：%v", errT.Error())
	}

	decoderT.Decode(&f2)
	decoderT.Decode(&b2)
	decoderT.Decode(&s2)
	decoderT.Decode(&person2)

	t.Printfln("n2=%#v", n2)
	t.Printfln("f2=%#v", f2)
	t.Printfln("b2=%#v", b2)
	t.Printfln("s2=%#v", s2)
	t.Printfln("person2=%#v", person2)

}
