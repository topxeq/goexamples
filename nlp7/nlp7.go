package main

import (
	"log"
	"tools"

	"github.com/topxeq/xiaoxian"
)

// 第一个测试文本字符串
var testDoc1G = `虽然电脑还缺少人类所具有的很多思考模式、逻辑创新、情感产生和变化的能力，但是在处理一些基于经验的、需要海量处理和计算（例如图片、语音、视频的识别等）的机械任务上，人工智能已经具备条件帮助人类去更快更准地完成。而进一步地，以大数据为基础的逻辑判断和行为决断（例如无人驾驶和医疗机器人），是深度学习下一步高歌猛进的目标。`

// 第二个测试文本字符串，仅比testDoc1G略有改动
var testDoc2G = `虽然计算机还缺少人类所具有的很多思考模式、逻辑创新、情感产生和变化的能力，但是在处理一些基于经验的、需要海量处理和计算（例如图片、语音、视频的识别等）的机械任务上，人工智能已经具备条件帮助人类去更快更准地完成。而进一步地，以大数据为基础的逻辑判断和行为决断（例如无人驾驶、电子竞技和医疗机器人），是深度学习下一步高歌猛进的目标。`

// 第三个测试文本字符串，与testDoc1G和testDoc2G差异较大
var testDoc3G = `道可道，非常道；名可名，非常名。无名，天地之始，有名，万物之母。故常无欲，以观其妙，常有欲，以观其徼。`

func main() {

	// 通过训练获得一个向量模型对象
	// 训练的目录指定为当前目录（“.”表示程序运行时当前所处的目录）
	// 第二个参数"text*.txt"表示将训练目录中所有符合该条件的文件加入
	// 第三个参数"trainData.txt"表示将清理的文本存入该文件
	// "d2v.model"是指定的存放训练后模型的文件路径
	// 最后两个数字参数分别表示生成向量的维度（300）和训练的轮数（30）
	modelT, errT := xiaoxian.TrainDoc2VecModel(".", "text*.txt", "trainData.txt", "d2v.model", 300, 30)

	// 用向量模型将testDoc1G生成一个向量
	vector1 := modelT.GetDocVectorMust(testDoc1G)

	tools.Printfln("文档向量1：%v", vector1)

	// model2T是从保存的模型文件中读取的模型对象
	// 为了演示载入模型的用法
	model2T, errT := xiaoxian.LoadD2VModel("d2v.model")

	if errT != nil {
		log.Fatalf("载入模型时发生错误：%v", errT)
	}

	// vector2是用载入的模型生成的向量
	vector2 := model2T.GetDocVectorMust(testDoc1G)

	tools.Printfln("文档向量2：%v", vector2)

	// 计算两个向量的相似度
	tools.Printfln("文档相似度1：%v", xiaoxian.CalCosineSimilarityOfVectors(vector1, vector2))

	// 直接计算两个字符串（代表文档）在该模型下的相似度
	tools.Printfln("文档相似度2：%v", model2T.GetSimilarityOfDocsEx(testDoc1G, testDoc1G))
	tools.Printfln("文档相似度3：%v", model2T.GetSimilarityOfDocsEx(testDoc1G, testDoc2G))
	tools.Printfln("文档相似度4：%v", model2T.GetSimilarityOfDocsEx(testDoc1G, testDoc3G))

}
