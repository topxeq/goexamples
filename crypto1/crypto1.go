package main

import (
	t "tools"
)

// encryptBytes 是用密码codeA来加密原文dataA的函数
func encryptBytes(dataA []byte, codeA []byte) []byte {

	// 获取原始数据的长度
	dataLenT := len(dataA)

	// 获取密码的长度
	codeLenT := len(codeA)

	// 创建放置密文的缓冲区，其大小与原文的字节长度应该是一样的
	encBufT := make([]byte, dataLenT)

	// 循环进行加密
	for i := 0; i < dataLenT; i++ {

		// codeA[i%codeLenT] 可以保证循环取到合理索引范围内的密码字节
		encBufT[i] = dataA[i] + codeA[i%codeLenT] + byte(i)
	}

	return encBufT
}

// decryptBytes 是用密码codeA来解密密文dataA的函数
func decryptBytes(dataA []byte, codeA []byte) []byte {

	// 获取密文数据的长度
	dataLenT := len(dataA)

	// 获取密码的长度
	codeLenT := len(codeA)

	// 创建放置还原的原文的缓冲区，其大小与密文的字节长度应该是一样的
	decBufT := make([]byte, dataLenT)

	// 循环进行解密
	for i := 0; i < dataLenT; i++ {

		// codeA[i%codeLenT] 可以保证循环取到合理索引范围内的密码字节
		decBufT[i] = dataA[i] - codeA[i%codeLenT] - byte(i)
	}

	return decBufT
}

func main() {

	originText := "This is an example."

	codeT := "test"

	t.Printfln("原文是：%#v", originText)
	t.Printfln("密码是：%#v", []byte(codeT))

	encBytes := encryptBytes([]byte(originText), []byte(codeT))

	t.Printfln("加密后的字节切片是：%#v", encBytes)

	decBytes := decryptBytes(encBytes, []byte(codeT))

	t.Printfln("解密后的字节切片是%#v", decBytes)
	t.Printfln("解密后还原的原文是%#v", string(decBytes))

}
