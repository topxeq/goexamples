package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 设置所绘制的每个数据点的默认大小
	plotter.DefaultGlyphStyle.Radius = vg.Points(2.0)

	// 新建一张图表
	p, _ := plot.New()

	//设置图表标题与X、Y轴说明文字
	p.Title.Text = "Lines and Points"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 设置Y轴的数值范围
	p.Y.Max = 4.0
	p.Y.Min = -2.0

	// 手动设置Y轴上显示的各个参考点数值及其显示字符串
	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{{-2.0, "-2 (Min)"}, {-1.0, ""}, {0, "0"}, {1.0, ""}, {2.0, "2"}, {3.0, ""}, {4.0, "4 (Max)"}})

	// 手动建立第一条折线的数据，共包含3个点
	points1 := plotter.XYs{
		{X: -8.0, Y: 2.0},
		{X: 1.0, Y: 2.5},
		{X: 7.0, Y: 4.0},
	}

	// 设置画线（不包含突出的点）使用的颜色
	plotter.DefaultLineStyle.Color = color.RGBA{R: 0x00, G: 0, B: 0xFF, A: 0xFF}

	// 新建一条不含突出点的折线
	line1, _ := plotter.NewLine(points1)

	// 为第二条（Sin函数）曲线分配1000点容量的空间
	points2 := make(plotter.XYs, 0, 1000)

	// 循环获取每个x点的Sin函数值
	for x := -10.0; x <= 10.0; x = x + 0.1 {
		y := math.Sin(x)

		points2 = append(points2, plotter.XY{X: x, Y: y})
	}

	// 为第三条（Cos函数）曲线分配1000点容量的空间
	points3 := make(plotter.XYs, 0, 1000)

	// 循环获取每个x点的Cos函数值
	for x := -10.0; x <= 10.0; x = x + 0.1 {
		y := math.Cos(x)

		points3 = append(points3, plotter.XY{X: x, Y: y})
	}

	// 设置画Sin函数散点图的颜色
	plotter.DefaultGlyphStyle.Color = color.RGBA{R: 0xFF, G: 0, B: 0, A: 0xFF}

	// 新建Sin函数的散点图
	scatter1, _ := plotter.NewScatter(points2)

	// 设置画Cos函数散点图的颜色
	plotter.DefaultGlyphStyle.Color = color.RGBA{R: 0, G: 0xFF, B: 0, A: 0xFF}

	// 新建Cos函数的散点图
	scatter2, _ := plotter.NewScatter(points3)

	// 在图表中加入一条折线与两个散点的线
	p.Add(scatter1, scatter2, line1)

	// 增加图例说明
	p.Legend.Add("line1", line1)
	p.Legend.Add("points2", scatter1)
	p.Legend.Add("points3", scatter2)

	// 保存图表到图片文件
	p.Save(8*vg.Inch, 8*vg.Inch, "points.png")
}
