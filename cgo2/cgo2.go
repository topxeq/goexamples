package main

/*
#include "test.h"
*/
import "C"

import "tools"

func main() {
	C.show3Times(C.CString(string(tools.ConvertBytesFromUTF8ToGB18030([]byte("[时间]")))))
}
