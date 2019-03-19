package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	t "tools"
)

type Point struct {
	X float64
	Y float32
}

func main() {

	n1 := 64
	f1 := 12.8
	b1 := false
	s1 := "abc123"
	point1 := Point{X: 1.8, Y: 3.5}

	bufT := new(bytes.Buffer)

	errT := binary.Write(bufT, binary.LittleEndian, int64(n1))

	if errT != nil {
		log.Fatalf("写入数据n1时发生错误：%v", errT.Error())
	}

	binary.Write(bufT, binary.LittleEndian, f1)
	binary.Write(bufT, binary.LittleEndian, b1)
	binary.Write(bufT, binary.LittleEndian, []byte(s1))
	binary.Write(bufT, binary.LittleEndian, point1)

	t.Printfln("bufT中内容：%#v", bufT.Bytes())

	ioutil.WriteFile(`c:\test\binaryFile1.bin`, bufT.Bytes(), 0666)

	bytesT, errT := ioutil.ReadFile(`c:\test\binaryFile1.bin`)

	if errT != nil {
		log.Fatalf("从文件中读取数据时发生错误：%v", errT.Error())
	}

	var n2 int64
	var f2 float64
	var b2 bool
	var s2buf []byte = make([]byte, 6)
	var point2 Point

	newBufT := bytes.NewReader(bytesT)

	errT = binary.Read(newBufT, binary.LittleEndian, &n2)

	if errT != nil {
		log.Fatalf("读入数据n2时发生错误：%v", errT.Error())
	}

	binary.Read(newBufT, binary.LittleEndian, &f2)
	binary.Read(newBufT, binary.LittleEndian, &b2)
	binary.Read(newBufT, binary.LittleEndian, &s2buf)
	binary.Read(newBufT, binary.LittleEndian, &point2)

	t.Printfln("n2=%#v", n2)
	t.Printfln("f2=%#v", f2)
	t.Printfln("b2=%#v", b2)
	t.Printfln("s2=%#v", string(s2buf))
	t.Printfln("point2=%#v", point2)

}
