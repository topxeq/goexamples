package main

import (
	"regexp"
	"strings"
	"tools"

	"github.com/topxeq/doc2vec/common"
	"github.com/topxeq/doc2vec/segmenter"
)

func regReplace(strA, patternA, replaceA string) string {
	regexpT, errT := regexp.Compile(patternA)

	if errT != nil {
		return strA
	}

	return regexpT.ReplaceAllString(strA, replaceA)
}

func prepareText(textA string) string {
	textT := strings.TrimSpace(strings.Replace(textA, "\r", "", -1))

	textListT := strings.Split(textT, "\n")

	seg := segmenter.GetSegmenter()

	seg.LoadDictionary("dict.txt")
	seg.LoadUserDictionary("userdict.txt")

	for i, v := range textListT {

	}

	qWords := []string{}

	for item := range seg.Cut(textT, false) {
		word := common.SBC2DBC(item.Text())
		qWords = append(qWords, word)
	}

	rs := strings.Join(qWords, " ")

	rs = regReplace(rs, `\s+`, " ")

	return rs
}

func train() {

}

func main() {
	fileContentT := tools.LoadStringFromFile("./testdata.txt", "")

	textT := prepareText(fileContentT)

	tools.Printfln("%v", textT)
}
