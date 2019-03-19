package main

import (
	"strings"
	"tools"

	"github.com/tidwall/gjson"
	"github.com/topxeq/txtk"

	"github.com/topxeq/xiaoxian"
)

var textEn1G = `We visited the Summer Palace when it was Chinese New Year.`

var textEn2G = `They walked the whole night and all the next day, from morning till
evening, but they did not get out of the forest, and were very hungry,
for they had nothing to eat but two or three berries, which grew on the
ground. `

var textCnG = `我们都非常非常的高兴。汤姆和杰瑞也是这样。大家围在一起又唱又跳，欢乐极了。`

func main() {

	resultT, errT := xiaoxian.ParseSentenceEnOL(textEn1G)

	if errT != nil {
		tools.Printfln("在线分析英文句子时发生错误：%v", errT.Error())
	} else {
		tools.Printfln("在线分析英文句子结果：%#v", resultT)
	}

	jsonObjectT := gjson.Parse(resultT)

	Confidence := txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Confidence").String(), 0)
	Intensity := txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Intensity").String(), 0)
	Modality := txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Modality").String(), 0)
	Mood := jsonObjectT.Get("Mood").String()
	Polarity := txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Polarity").String(), 0)
	Subjectivity := txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Subjectivity").String(), 0)
	Tense := jsonObjectT.Get("Tense").String()

	tmps := strings.ToLower(Mood)
	if tmps == "indicative" {
		tmps = "陈述语气"
	} else if tmps == "imperative" {
		tmps = "祈使语气、命令式"
	} else if tmps == "conditional" {
		tmps = "假设、条件语气"
	} else if tmps == "subjunctive" {
		tmps = "虚拟语气"
	} else if tmps == "question" {
		tmps = "疑问语气"
	} else if tmps == "exclamation" {
		tmps = "感叹语气"
	}

	tools.Printfln("情感倾向（消极 -1.0 ——> 1.0 积极）：%v", Polarity)
	tools.Printfln("情感强烈度：%v", Intensity)
	tools.Printfln("肯定性（虚构 -1.0 ——> 1.0 事实）：%v", Modality)
	tools.Printfln("主观性（弱 0.0 —— 1.0 强）：%v", Subjectivity)
	tools.Printfln("自信度：%v", Confidence)
	tools.Printfln("语气：%v", tmps)
	tools.Printfln("时态：%v", Tense)

	resultT, errT = xiaoxian.ParseSentenceEnOL(textEn2G)

	if errT != nil {
		tools.Printfln("在线分析英文句子时发生错误：%v", errT.Error())
	} else {
		tools.Printfln("在线分析英文句子结果：%#v", resultT)
	}

	jsonObjectT = gjson.Parse(resultT)

	Confidence = txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Confidence").String(), 0)
	Intensity = txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Intensity").String(), 0)
	Modality = txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Modality").String(), 0)
	Mood = jsonObjectT.Get("Mood").String()
	Polarity = txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Polarity").String(), 0)
	Subjectivity = txtk.StrToFloat64WithDefaultValue(jsonObjectT.Get("Subjectivity").String(), 0)
	Tense = jsonObjectT.Get("Tense").String()

	tmps = strings.ToLower(Mood)
	if tmps == "indicative" {
		tmps = "陈述语气"
	} else if tmps == "imperative" {
		tmps = "祈使语气、命令式"
	} else if tmps == "conditional" {
		tmps = "假设、条件语气"
	} else if tmps == "subjunctive" {
		tmps = "虚拟语气"
	} else if tmps == "question" {
		tmps = "疑问语气"
	} else if tmps == "exclamation" {
		tmps = "感叹语气"
	}

	tools.Printfln("情感倾向（消极 -1.0 ——> 1.0 积极）：%v", Polarity)
	tools.Printfln("情感强烈度：%v", Intensity)
	tools.Printfln("肯定性（虚构 -1.0 ——> 1.0 事实）：%v", Modality)
	tools.Printfln("主观性（弱 0.0 —— 1.0 强）：%v", Subjectivity)
	tools.Printfln("自信度：%v", Confidence)
	tools.Printfln("语气：%v", tmps)
	tools.Printfln("时态：%v", Tense)

	pageResultT, errStrT, tokenT := xiaoxian.SentimentCnBaiduOL(textCnG, "", "XXXXXXX", "XXXXXXXXX")

	if errStrT != "" {
		tools.Printfln("\n进行在线中文情感分析时发生错误：%v，token: %v", errStrT, tokenT)
		return
	}

	tools.Printfln("\n在线中文情感分析结果：%#v", pageResultT)

}
