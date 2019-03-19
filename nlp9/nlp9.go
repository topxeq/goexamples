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

	filterT := map[string]int{"地": 1, "的": 1, "又": 1, "说": 1}

	listT := modelInternalT.TXWord2Words(`深度学习`, 5, filterT)

	tools.Printfln("单词关联的单词：\n%v", strings.Join(listT, "\n"))

}
