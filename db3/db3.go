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

	// 查询表中所有符合条件的记录总数

	var countT int64

	errT = dbT.QueryRow("select count(*) from TEST").Scan(&countT)
	if errT != nil {
		t.Printfln("执行SQL查询语句时发生错误：%v", errT.Error())
		return
	}

	t.Printfln("库表中共有%v条记录", countT)

	// 删除库表

	_, errT = dbT.Exec(`drop table TEST`)
	if errT != nil {
		t.Printfln("删除库表时发生错误：%v", errT.Error())
		return
	}

	// 再次查询时会提示出错，因为库表已经被删除了
	errT = dbT.QueryRow("select count(*) from TEST").Scan(&countT)
	if errT != nil {
		t.Printfln("执行SQL查询语句时发生错误：%v", errT.Error())
		return
	}

	t.Printfln("库表中共有%v条记录", countT)

}
