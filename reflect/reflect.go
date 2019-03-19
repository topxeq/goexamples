package main

import (
	"encoding/json"
	"reflect"
	t "tools"
)

var jsonTextG = `{
"ID": "12345",
"name": "张三",
"曾用名": ["李四", "王五"],
"年龄": 28,
"电话": {"座机": "66668888", "手机": "13333338888"},
"pets": [
	{"name": "Tom", "type": "cat"},
	{"name": "Jerry", "type": "mouse"}
	]
}
`

func analyzeJsonObject(vA interface{}, compareKeyA string, lastKeyA string, listA *[]string) {

	valueT := reflect.ValueOf(vA)

	switch valueT.Kind() {
	case reflect.String:
		if compareKeyA != "" {
			if lastKeyA != compareKeyA {
				break

			}
		}

		*listA = append(*listA, valueT.String())
	case reflect.Slice:
		for i := 0; i < valueT.Len(); i++ {
			v := valueT.Index(i)

			analyzeJsonObject(v.Interface(), compareKeyA, "", listA)
		}
	case reflect.Map:
		iter := valueT.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			analyzeJsonObject(v.Interface(), compareKeyA, k.String(), listA)

		}
	default:
		t.Printfln("遇到未处理的数据类型: %v, 值为: %v\n", valueT.Kind(), valueT)
	}
}

func main() {

	var v interface{}

	errT := json.Unmarshal([]byte(jsonTextG), &v)

	if errT != nil {
		t.Printfln("JSON解析错误：%v", errT.Error())
		return
	}

	t.Printfln("解析JSON文本后获得的数据类型为：%v\n", reflect.TypeOf(v))

	listT := make([]string, 0, 10)

	analyzeJsonObject(v, "", "", &listT)

	t.Printfln("结果1：%#v\n", listT)

	listT = make([]string, 0, 10)

	analyzeJsonObject(v, "name", "", &listT)

	t.Printfln("结果2：%#v\n", listT)

}
