package main

import (
	"encoding/json"
	t "tools"
)

func main() {

	jsonT := `{
		"name": "Bob",
		"type": "dog",
		"weight": 3.3,
		"color": "brown",
		"parents": [
			{
			"name": "Tom",
			"type": "dog",
			"weight": 8.3,
			"color": "Black"
			},
			{
				"name": "Mary",
				"type": "dog",
				"weight": 6.7,
				"color": "White"
			}
		]
	}`

	var bufT interface{}

	errT := json.Unmarshal([]byte(jsonT), &bufT)

	if errT != nil {
		t.Printfln("JSON解码时发生错误: %v", errT.Error())
		return
	}

	t.Printfln("bufT: %#v", bufT)

	parents := bufT.(map[string]interface{})["parents"]

	t.Printfln("parents: %#v", parents)

	parentsLen := len(parents.([]interface{}))

	t.Printfln("parentsLen: %v", parentsLen)

	firstParent := parents.([]interface{})[0]

	t.Printfln("firstParent: %#v", firstParent)

	firstParentName := firstParent.(map[string]interface{})["name"]

	t.Printfln("firstParentName: %#v", firstParentName)

	t.Printfln("firstParentWeight: %#v", bufT.(map[string]interface{})["parents"].([]interface{})[0].(map[string]interface{})["weight"])

	t.Printfln("firstParent中共有%v个key/value对", len(firstParent.(map[string]interface{})))

	for k, v := range firstParent.(map[string]interface{}) {
		t.Printfln("第一条狗的%v属性值为：%v", k, v)
	}

}
