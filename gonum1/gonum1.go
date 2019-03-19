package main

import (
	t "tools"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// 定义参与运算的矩阵 a和 b
	var a *mat.Dense
	a = mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6})

	b := mat.NewDense(3, 2, []float64{
		1, 2,
		3, 4,
		5, 6})

	// 定义用于存放计算结果的矩阵c
	var c mat.Dense

	// 进行矩阵乘法运算
	c.Mul(a, b)

	// 用gonum/mat包中的格式来输出结果矩阵c的信息
	cOutput := mat.Formatted(&c)

	t.Printfln("结果矩阵1：\n%v", cOutput)

	// 按自定义的格式输出结果矩阵c的信息
	row, col := c.Caps() // row, col 是矩阵的行数与列数

	t.Printf("\n结果矩阵1的手动控制输出:\n")

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			t.Printf(" %8.2f", c.At(i, j)) // 每行内的数据连续输出
		}

		t.Printfln("") // 输出完每行数据后输出一个换行符
	}

	// 修改矩阵a中第 1行第 2列的数据为 5
	a.Set(0, 1, 5)

	// 再次进行矩阵乘法运算
	c.Mul(a, b)

	// 输出第二次矩阵乘法的结果
	cOutput = mat.Formatted(&c)

	t.Printfln("\n\n结果矩阵2：\n%v", cOutput)

}
