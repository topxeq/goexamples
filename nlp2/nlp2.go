package main

import (
	"strings"
	"tools"

	"github.com/topxeq/xiaoxian"
)

var textEnG = `They walked the whole night and all the next day, from morning till
evening, but they did not get out of the forest, and were very hungry,
for they had nothing to eat but two or three berries, which grew on the
ground. `

var textCnG = `但奇迹般的是，随着深度学习技术的横空出世，人工智能又神奇地焕发出了再一次的青春。`

func main() {

	textT := xiaoxian.CleanEnglish(textEnG)

	listT := xiaoxian.TokenizeEn(textT)

	tools.Printfln("英文分词结果：%#v", listT)

	listT, errT := xiaoxian.TokenizeEnOL(textT)

	if errT != nil {
		tools.Printfln("\n在线英文分词时发生错误：%v", errT.Error())
	} else {
		tools.Printfln("\n在线英文分词结果：%#v", listT)
	}

	tokenizerT, errT := xiaoxian.NewTokenizerCn("./dict.txt", "")

	if errT != nil {
		tools.Printfln("\n创建中文分词器时发生错误：%v", errT.Error())
		return
	}

	listT = tokenizerT.Tokenize(textCnG, false)

	tools.Printfln("\n中文分词结果：%#v", listT)

	errT = tokenizerT.LoadUserDict("./userdict.txt")

	if errT != nil {
		tools.Printfln("\n载入用户词典时发生错误：%v", errT.Error())
		return
	}

	listT = tokenizerT.Tokenize(textCnG, false)

	tools.Printfln("\n载入词典后中文分词结果：%#v", listT)

	pageResultT, errStrT, tokenT := xiaoxian.TokenizeCnBaiduOL(textCnG, false, "", "XXXXXXX", "XXXXXXXX")

	if errStrT != "" {
		tools.Printfln("\n进行在线中文分词时发生错误：%v，token: %v", errStrT, tokenT)
		return
	}

	tools.Printfln("\n在线中文分词结果：%#v", strings.Split(pageResultT, " "))

}
