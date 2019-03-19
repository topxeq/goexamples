package main

import (
	"fmt"

	"github.com/d5/tengo/objects"
	"github.com/d5/tengo/script"
)

// 定义将要执行的Tengo语言脚本代码
var codeT = `

sum := 0 // 变量初始化和赋值语句与Go语言基本一致

a := 0

// 循环语句也很像Go语言中的
// 变量maxA是将从Go语言中传递进来的
for a < maxA {
	sum += a
	a ++
}

// 也有printf等内置函数
// times函数是在Go语言中定义的
// 这里演示了在脚本语言中调用Go语言中代码的能力
printf("%d\n", times(2, 3, 4))

`

// 供脚本语言调用的函数，用于计算不定个数的参数的累乘积
// 其函数形式（参数和返回值）必须是这样
func times(objsA ...objects.Object) (objects.Object, error) {
	lenT := len(objsA)

	intListT := make([]int, lenT)

	// 用一个循环将函数不定个数参数中的所有数值存入整数切片中
	for i, v := range objsA {
		// 调用objects.ToInt函数将objects.Object对象转换为整数
		cT, ok := objects.ToInt(v)

		if ok {
			intListT[i] = cT
		}
	}

	// 进行累乘与那算
	r := 1

	for i := 0; i < lenT; i++ {
		r = r * intListT[i]
	}

	// 输出结果值供参考
	fmt.Printf("result: %v\n", r)

	// 也作为函数返回值返回，返回前要转换为objects.Object类型
	// objects.Int类型实现了objects.Object类型，因此可以用作返回值
	return &objects.Int{Value: int64(r)}, nil
}

func main() {

	// 新建一个脚本运行的虚拟机
	// 一般会编译为字节码准备运行
	s := script.New([]byte(codeT))

	// 向脚本执行环境（虚拟机VM）中传入变量maxA
	_ = s.Add("maxA", 20)

	//传入准备在虚拟机中执行的Go语言编写的函数times
	_ = s.Add("times", times)

	// 执行脚本
	c, err := s.Run()
	if err != nil {
		panic(err)
	}

	// 获取返回值（脚本中的sum变量）
	sumT := c.Get("sum")

	// 转换类型后输出
	fmt.Println(sumT.Int())
}
