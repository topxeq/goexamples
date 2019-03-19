package main

import (
	"database/sql"
	"os"
	t "tools"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// 如果存在该库（SQLite库是放在单一的文件中的）则删除该文件
	if t.FileExists(`c:\test\test.db`) {
		os.Remove(`c:\test\test.db`)
	}

	// 创建新库
	dbT, errT := sql.Open("sqlite3", `c:\test\test.db`)

	if errT != nil {
		t.Printfln("创建数据库时发生错误：%v", errT.Error())
		return
	}

	defer dbT.Close()

	//创建表
	sqlStmtT := `
	create table TEST (ID integer not null primary key, CODE text);
    `
	_, errT = dbT.Exec(sqlStmtT)
	if errT != nil {
		t.Printfln("创建表时发生错误：%v", errT.Error())
		return
	}

	// 开始一个数据库事务
	txT, errT := dbT.Begin()
	if errT != nil {
		t.Printfln("新建事务时发生错误：%v", errT.Error())
		return
	}

	// 准备一个SQL语句，用于向表中插入记录
	stmtT, errT := txT.Prepare("insert into TEST(ID, CODE) values(?, ?)")
	if errT != nil {
		t.Printfln("准备SQL语句插入记录时发生错误：%v", errT.Error())
		return
	}

	defer stmtT.Close()

	// 向表中插入10条记录
	// 每条记录的ID字段用循环变量的值赋值
	// CODE字段用随机产生的字符串
	for i := 0; i < 10; i++ {
		_, errT = stmtT.Exec(i, t.GenerateRandomString(5, 8, true, true, true, false, false, false))
		if errT != nil {
			t.Printfln("执行SQL插入记录语句时发生错误：%v", errT.Error())
			return
		}
	}

	// 执行事务，此时新纪录才会被真正插入到表中
	txT.Commit()

	// 进行SQL查询
	rowsT, errT := dbT.Query("select ID, CODE from TEST")
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
