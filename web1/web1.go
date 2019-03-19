package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	t "tools"
)

func japiHandler(w http.ResponseWriter, r *http.Request) {
	reqT := r.FormValue("req")

	returnObjectT := make(map[string]string)

	switch reqT {
	case "":
		returnObjectT["Status"] = "fail"
		returnObjectT["Value"] = fmt.Sprintf("请求不能为空")
	case "getTime":
		returnObjectT["Status"] = "success"
		returnObjectT["Value"] = fmt.Sprintf("%v", time.Now())
	case "generatePassword":
		lenStrT := r.FormValue("len")

		if lenStrT == "" {
			returnObjectT["Status"] = "fail"
			returnObjectT["Value"] = "需要指定长度（len）"
			break
		}

		var lenT int

		_, errT := fmt.Sscanf(lenStrT, "%d", &lenT)
		if errT != nil {
			returnObjectT["Status"] = "fail"
			returnObjectT["Value"] = "长度格式不正确"
			break
		}

		returnObjectT["Status"] = "success"
		returnObjectT["Value"] = fmt.Sprintf("%v", t.GenerateRandomString(lenT, lenT, true, true, true, true, false, false))
	default:
		returnObjectT["Status"] = "fail"
		returnObjectT["Value"] = fmt.Sprintf("未知的请求: %v", reqT)
	}

	bufT, _ := json.Marshal(returnObjectT)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/json;charset=utf-8")

	w.WriteHeader(http.StatusOK)

	callbackT := r.FormValue("callback")
	if callbackT != "" {
		w.Write([]byte(fmt.Sprintf("%v(%v);", callbackT, string(bufT))))
	} else {
		w.Write(bufT)
	}
}

func main() {

	portT := "8838"

	if len(os.Args) >= 2 {
		portT = os.Args[1]
	}

	http.HandleFunc("/japi/", japiHandler)

	err := http.ListenAndServe(":"+portT, nil)
	if err != nil {
		t.Printfln("打开HTTP监听端口时发生错误：%v", err.Error())
	}

}
