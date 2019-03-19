package main

import (
	"fmt"
	"tools"

	"github.com/tidwall/gjson"
	"github.com/topxeq/xiaoxian"
)

var textEn1G = `We visited the Summer Palace when it was Chinese New Year.`

var textEn2G = `They walked the whole night and all the next day, from morning till
evening, but they did not get out of the forest, and were very hungry,
for they had nothing to eat but two or three berries, which grew on the
ground. `

var textEn3G = `They walked the whole night and all the next day, from morning till
evening, but they did not get out of the forest, and were very hungry,
for they had nothing to eat but two or three berries, which grew on the
ground. And as they were so weary that their legs would carry them no
longer, they lay down beneath a tree and fell asleep.

It was now three mornings since they had left their father’s house.
They began to walk again, but they always got deeper into the forest.
If help did not come soon, they must die of hunger and weariness!`

func main() {

	resultT, errT := xiaoxian.GetArticleDifficultyEnOL(textEn1G)

	if errT != nil {
		tools.Printfln("在线获取英文篇章难度级别时发生错误：%v", errT.Error())
	} else {
		tools.Printfln("在线获取英文篇章难度级别结果：%#v", resultT)
	}

	jsonObjectT := gjson.Parse(resultT)

	deepScoreT := jsonObjectT.Get("DeepScore").String()

	var totalWordCountT int

	fmt.Sscanf(jsonObjectT.Get("TotalWordCount").String(), "%d", &totalWordCountT)

	totalWordLemmaCount := jsonObjectT.Get("TotalWordLemmaCount").String()

	var totalSyllableCount int

	fmt.Sscanf(jsonObjectT.Get("TotalSyllableCount").String(), "%d", &totalSyllableCount)

	avgSyllableCount := float64(totalSyllableCount) / float64(totalWordCountT)

	fleschReadingEaseT := jsonObjectT.Get("FleschReadingEase").String()

	tools.Printfln("textEn1G难度级别评分（0-18）：%v", deepScoreT)
	tools.Printfln("单词总数：%v", totalWordCountT)
	tools.Printfln("单词原型总数：%v", totalWordLemmaCount)
	tools.Printfln("单词平均音节数：%v", avgSyllableCount)
	tools.Printfln("Flesch阅读难度指数：%v\n", fleschReadingEaseT)

	resultT, errT = xiaoxian.GetArticleDifficultyEnOL(textEn2G)

	if errT != nil {
		tools.Printfln("在线获取英文篇章难度级别时发生错误：%v", errT.Error())
	}

	jsonObjectT = gjson.Parse(resultT)

	deepScoreT = jsonObjectT.Get("DeepScore").String()

	fmt.Sscanf(jsonObjectT.Get("TotalWordCount").String(), "%d", &totalWordCountT)

	totalWordLemmaCount = jsonObjectT.Get("TotalWordLemmaCount").String()

	fmt.Sscanf(jsonObjectT.Get("TotalSyllableCount").String(), "%d", &totalSyllableCount)

	avgSyllableCount = float64(totalSyllableCount) / float64(totalWordCountT)

	fleschReadingEaseT = jsonObjectT.Get("FleschReadingEase").String()

	tools.Printfln("textEn2G难度级别评分（0-18）：%v", deepScoreT)
	tools.Printfln("单词总数：%v", totalWordCountT)
	tools.Printfln("单词原型总数：%v", totalWordLemmaCount)
	tools.Printfln("单词平均音节数：%v", avgSyllableCount)
	tools.Printfln("Flesch阅读难度指数：%v\n", fleschReadingEaseT)

	resultT, errT = xiaoxian.GetArticleDifficultyEnOL(textEn3G)

	if errT != nil {
		tools.Printfln("在线获取英文篇章难度级别时发生错误：%v", errT.Error())
	}

	jsonObjectT = gjson.Parse(resultT)

	deepScoreT = jsonObjectT.Get("DeepScore").String()

	fmt.Sscanf(jsonObjectT.Get("TotalWordCount").String(), "%d", &totalWordCountT)

	totalWordLemmaCount = jsonObjectT.Get("TotalWordLemmaCount").String()

	fmt.Sscanf(jsonObjectT.Get("TotalSyllableCount").String(), "%d", &totalSyllableCount)

	avgSyllableCount = float64(totalSyllableCount) / float64(totalWordCountT)

	fleschReadingEaseT = jsonObjectT.Get("FleschReadingEase").String()

	tools.Printfln("textEn3G难度级别评分（0-18）：%v", deepScoreT)
	tools.Printfln("单词总数：%v", totalWordCountT)
	tools.Printfln("单词原型总数：%v", totalWordLemmaCount)
	tools.Printfln("单词平均音节数：%v", avgSyllableCount)
	tools.Printfln("Flesch阅读难度指数：%v", fleschReadingEaseT)

}
