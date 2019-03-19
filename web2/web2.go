package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	t "tools"
)

func dynamicHandler(w http.ResponseWriter, r *http.Request) {

	urlPathT := strings.TrimPrefix(r.RequestURI, "/dynamic/")

	t.Printfln("urlPath: %v", urlPathT)

	var htmlTextT string

	switch urlPathT {
	case "index.html", "", "next/test.html":
		tmplFilePathT := filepath.Join(`c:\test\tmpl`, "indextmpl.html")

		t.Printfln("filePath: %v", tmplFilePathT)

		if !t.FileExists(tmplFilePathT) {
			http.NotFound(w, r)
			return
		}

		templateT, errT := template.ParseFiles(tmplFilePathT)

		if errT != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dataT := map[string]interface{}{
			"RemoteAddr":    r.RemoteAddr,
			"RandomStrings": []string{t.GenerateRandomString(5, 8, true, true, true, false, false, false), t.GenerateRandomString(5, 8, true, true, true, false, false, false), t.GenerateRandomString(5, 8, true, true, true, false, false, false)},
		}

		var sb strings.Builder

		templateT.ExecuteTemplate(&sb, "indextmpl.html", dataT)

		htmlTextT = sb.String()
	default:
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(htmlTextT))
}

func main() {

	portT := "8838"

	if len(os.Args) >= 2 {
		portT = os.Args[1]
	}

	http.HandleFunc("/dynamic/", dynamicHandler)

	err := http.ListenAndServe(":"+portT, nil)
	if err != nil {
		t.Printfln("打开HTTP监听端口时发生错误：%v", err.Error())
	}

}
