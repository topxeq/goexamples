package main

import (
	t "tools"

	"github.com/topxeq/xiaoxian"
)

var textEnG = `They walked the whole night and all the next day, from morning till
evening, but they did not get out of the forest, and were very hungry,
for they had nothing to eat but two or three berries, which grew on the
ground. And as they were so weary that their legs would carry them no
longer, they lay down beneath a tree and fell asleep.

It was now three mornings since they had left their father’s house.
They began to walk again, but they always got deeper into the forest.
If help did not come soon, they must die of hunger and weariness!`

var textCnG = `人工智能（Artificial Intelligence，简称AI）从孕育诞生至今，已经至少有近80年历史了。80年的光阴，虽然在历史的长河中不过是浪花一朵，但如果以人的一生来说，已经是进入耄耋老年了。但奇迹般的是，随着深度学习技术的横空出世，人工智能又神奇地焕发出了再一次的青春。深度学习系统AlphaGo及其升级版本一再战胜围棋领域的多位世界冠军级选手，最后甚至到了一败难求、人类选手只能仰视的地步，不能不说是引起了世人广泛关注人工智能领域的决定性事件。指纹识别、人脸识别、无人驾驶等应用了深度学习方法而又贴近人们日常生活的技术，可以说深刻地改变了人类的生活和消费方式，也因此让人工智能更加深入人心，激起了人工智能尤其是深度学习领域的学习热潮。`

func main() {

	textT := xiaoxian.CleanEnglish(textEnG)

	listT := xiaoxian.SplitArticleEn(textT)

	t.Printfln("英文分句结果：\n")

	for i, v := range listT {
		t.Printfln("第%v句：%v", i+1, v)
	}

	t.Printfln("\n-----\n")

	listT, errT := xiaoxian.SplitArticleEnOL(textT)

	if errT != nil {
		t.Printfln("在线英文分句时发生错误：%v", errT.Error())
	} else {
		t.Printfln("在线英文分句结果\n")

		for i, v := range listT {
			t.Printfln("第%v句：%v", i+1, v)
		}

	}

	t.Printfln("\n-----\n")

	listT = xiaoxian.SplitArticleCn(textCnG)

	t.Printfln("中文分句结果\n")

	for i, v := range listT {
		t.Printfln("第%v句：%v", i+1, v)
	}

}
