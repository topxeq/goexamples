package main

import (
	"fmt"
	"log"

	// 应用所需的包
	"github.com/dgraph-io/badger"
)

func main() {

	// 准备创建数据库的设置参数
	optionsT := badger.DefaultOptions

	// 设置数据库的工作目录
	optionsT.Dir = `c:\test\db`
	optionsT.ValueDir = `c:\test\db`

	// 创建或打开数据库
	dbT, errT := badger.Open(optionsT)
	if errT != nil {
		log.Fatal(errT)
	}

	// 确保退出前关闭数据库
	defer dbT.Close()

	// 准备用于测试的数据，是映射类型的数据
	dataT := make(map[string]string)

	dataT["ok"] = "yes"
	dataT["名字"] = "张三"

	// 新建一个事务
	transT := dbT.NewTransaction(true)

	// 遍历测试数据并存入数据库
	for k, v := range dataT {

		// 将对应的键值对存入
		errT := transT.Set([]byte(k), []byte(v))

		if errT != nil {
			fmt.Printf("设置KV对时发生错误：%v", errT.Error())
		}

	}

	// 提交事务，此时才真正写入数据库
	_ = transT.Commit()

	// 输出分隔线
	fmt.Printf("\n-----\n")

	// 再次新建一个事务用于查询
	transT = dbT.NewTransaction(true)

	// 查询键名为ok对应的键值
	itemT, errT := transT.Get([]byte("ok"))

	if errT != nil {
		log.Fatalf("获取KV对时发生错误：%v", errT)
	}

	// 获取键值
	valueT, errT := itemT.ValueCopy(nil)

	if errT != nil {
		log.Fatalf("获取KV对值时发生错误：%v", errT)
	}

	fmt.Printf("获取到的键名为%v的键值：%v\n", "ok", string(valueT))

	// 用只读模式打开数据库后遍历其中所有的键值对
	errT = dbT.View(func(txn *badger.Txn) error {

		// 准备遍历数据库中键值对的设置参数，这里用的是默认设置
		optionsT := badger.DefaultIteratorOptions

		// 设置预获取的数量
		optionsT.PrefetchSize = 10

		// 创建遍历用的枚举对象
		iteratorT := txn.NewIterator(optionsT)

		// 确保枚举对象被关闭
		defer iteratorT.Close()

		// 进行遍历
		for iteratorT.Rewind(); iteratorT.Valid(); iteratorT.Next() {

			// 获取一个枚举值
			itemT := iteratorT.Item()

			// 获取该枚举值中的键值
			k := itemT.Key()

			// 调用匿名函数获取键值并处理
			errT := itemT.Value(func(v []byte) error {
				fmt.Printf("键名：%s，键值：%s\n", k, v)
				return nil
			})

			if errT != nil {
				return errT
			}

		}

		return nil
	})

	if errT != nil {
		log.Fatalf("遍历KV对时发生错误：%v", errT)
	}

	fmt.Printf("\n-----\n")

}
