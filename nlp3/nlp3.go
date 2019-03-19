package main

import (
	"tools"

	"github.com/topxeq/xiaoxian"
)

var textEnG = `They walked the whole night and all the next day, from morning till
evening, but they did not get out of the forest, and were very hungry,
for they had nothing to eat but two or three berries, which grew on the
ground.`

var textCnG = `但奇迹般的是，随着深度学习技术的横空出世，人工智能又神奇地焕发出了再一次的青春。`

func main() {

	listT, errT := xiaoxian.TagEnOL(textEnG)

	if errT != nil {
		tools.Printfln("在线英文词性识别时发生错误：%v", errT.Error())
	} else {
		tools.Printfln("在线英文词性识别结果：%#v", listT)
	}

	posTaggerT, errT := xiaoxian.NewPosTaggerCn("./dict.txt", "./userdict.txt")

	if errT != nil {
		tools.Printfln("\n创建中文词性识别器时发生错误：%v", errT.Error())
		return
	}

	listT = posTaggerT.Tag(textCnG, false, "")

	tools.Printfln("\n中文词性识别结果：%#v", listT)

}
