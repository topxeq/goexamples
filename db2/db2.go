package main

import (
	"database/sql"
	t "tools"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dbFileT := `c:\test\test.db`

	if !t.FileExists(dbFileT) {
		t.Printfln("数据库文件%v不存在", dbFileT)
		return
	}

	// 打开已存在的库
	dbT, errT := sql.Open("sqlite3", dbFileT)

	if errT != nil {
		t.Printfln("打开数据库时发生错误：%v", errT.Error())
		return
	}

	defer dbT.Close()

	// 用“准备”方式做SQL查询
	stmtT, errT := dbT.Prepare("select CODE from TEST where ID = ?")
	if errT != nil {
		t.Printfln("准备SQL查询语句时发生错误：%v", errT.Error())
		return
	}

	defer stmtT.Close()

	var codeT string

	errT = stmtT.QueryRow("5").Scan(&codeT)

	if errT != nil {
		t.Printfln("从查询结果中获取字段数值时发生错误：%v", errT.Error())
		return
	}

	t.Printfln("第5条记录中的CODE字段值为：%v", codeT)

	// 删除TEST表中所有记录
	_, errT = dbT.Exec("delete from TEST")
	if errT != nil {
		t.Printfln("删除数据库表记录时发生错误：%v", errT.Error())
		return
	}

	// 重新插入3条记录
	_, errT = dbT.Exec("insert into TEST(ID, CODE) values(5, '汤姆'), (10, '杰瑞'), (18, '史诺比')")
	if errT != nil {
		t.Printfln("插入新数据库表记录时发生错误：%v", errT.Error())
		return
	}

	// 再次进行SQL查询
	rowsT, errT := dbT.Query("select * from TEST")
	if errT != nil {
		t.Printfln("执行SQL查询语句时发生错误：%v", errT.Error())
		return
	}

	defer rowsT.Close()

	// 遍历查询结果
	for rowsT.Next() {
		var idT int
		var codeT string

		errT = rowsT.Scan(&idT, &codeT)
		if errT != nil {
			t.Printfln("遍历查询结果时发生错误：%v", errT.Error())
			return
		}

		t.Printfln("ID: %v, CODE: %v", idT, codeT)
	}

	// 检查查询结果的错误
	errT = rowsT.Err()
	if errT != nil {
		t.Printfln("查询结果有错误：%v", errT.Error())
	}

}
