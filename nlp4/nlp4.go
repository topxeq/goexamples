package main

import (
	"strings"
	"tools"

	"github.com/topxeq/xiaoxian"
)

var textEnG = `We visited the Summer Palace when it was Chinese New Year.`

var textCnG = `我们今年春节去颐和园玩了一次。`

func main() {

	listT, errT := xiaoxian.GetNamedEntityEnOL(textEnG)

	if errT != nil {
		tools.Printfln("在线英文命名实体识别时发生错误：%v", errT.Error())
	} else {
		tools.Printfln("在线英文命名实体识别结果：%#v", listT)
	}

	pageResultT, errStrT, tokenT := xiaoxian.NerCnBaiduOL(textCnG, false, "", "XXXXXXXX", "XXXXXXXXXX")

	if errStrT != "" {
		tools.Printfln("\n进行在线中文命名实体识别时发生错误：%v，token: %v", errStrT, tokenT)
		return
	}

	tools.Printfln("\n在线中文命名实体识别结果：%#v", strings.Split(pageResultT, " "))

}
