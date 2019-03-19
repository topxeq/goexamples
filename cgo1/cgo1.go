package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// 是否初始化随机数种子的标志
int randomizeFlagG = 0;

// 获取一个随机数
int getRandomInt(int maxA) {
	if (randomizeFlagG == 0) {
		// 初始化随机数种子
		srand(time(NULL));
	}

	// 限制随机数值范围
	int rs = rand()%maxA;

	// 为了演示，在C语言函数中输出一下生成的随机数
	printf("%d\n", rs);

	return rs;
}

// 输出一个字符串
void printString(char *str) {
    printf("%s\n", str);
}

*/
import "C"

import (
	"fmt"
	"tools"
	"unsafe"
)

func main() {
	// 调用C的标准库函数puts来输出
	C.puts(C.CString(string(tools.ConvertBytesFromUTF8ToGB18030([]byte("这是一个test.")))))

	// 将Go语言字符串转换为C语言格式的字符串
	cStrT := C.CString(string(tools.ConvertBytesFromUTF8ToGB18030([]byte("测试字符串"))))

	// 调用C语言中自定义的函数来输出
	C.printString(cStrT)

	// 确保释放C语言格式的字符串所占用的内存空间
	defer C.free(unsafe.Pointer(cStrT))

	// 调用C语言中定义的函数获取一个随机数
	rs := C.getRandomInt(20)

	// 输出该随机数及其在Go语言中的类型
	fmt.Printf("%T, %#v\n", rs, rs)
}
