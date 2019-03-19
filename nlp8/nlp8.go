package main

import (
	"log"
	"strings"
	"tools"

	"github.com/topxeq/xiaoxian"
)

func main() {

	modelT, errT := xiaoxian.TrainDoc2VecModel(".", "text*.txt", "trainData.txt", "d2v.model", 300, 30)

	if errT != nil {
		log.Fatalf("训练模型时发生错误：%v", errT)
	}

	modelInternalT := *modelT.Model()

	listT := modelInternalT.TXWord2Docs(`深度学习`, 5)

	tools.Printfln("单词关联的文档：\n%v", strings.Join(listT, "\n"))

}
