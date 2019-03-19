package main

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"math"

	t "tools"
)

// sigmoid 计算sigmoid的函数
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

func main() {
	// 为即将使用的中文字体起名
	fontName := "simhei"

	// 读入改TrueType字体文件
	ttfBytes, err := ioutil.ReadFile("/xiaoxian/web/txhelp/simhei.ttf")

	// 读取文件发生异常时的处理
	if err != nil {
		t.Printfln("读取字体文件失败：%v", err.Error())
		return
	}

	// 将字体加入字体表（仅本次运行有效）
	font, _ := truetype.Parse(ttfBytes)
	vg.AddFont(fontName, font)

	// 设置gonum/plot绘图时默认字体为自定义的中文字体
	plot.DefaultFont = fontName

	// 设置所绘制的每个数据点的大小
	plotter.DefaultGlyphStyle.Radius = vg.Points(1.0)

	// 新建一张图表
	p, _ := plot.New()

	//设置图表标题与X、Y轴说明文字
	p.Title.Text = "Gonum plot示例"
	p.X.Label.Text = "X轴"
	p.Y.Label.Text = "Y轴"

	// 手动建立第一条折线的数据，共包含3个点
	points1 := plotter.XYs{
		{0.0, 0.0},
		{1.0, 1.0},
		{2.0, 4.0},
	}

	// 为第二条折线（sigmoid函数对应的曲线）分配内存空间
	points2 := make(plotter.XYs, 0, 4)

	// 让x的值从-10开始至10结束，每次步进0.1，循环计算每个x值对应的sigmoid函数计算结果值y；然后将该值加入到切片变量points2中
	for x := -10.0; x <= 10.0; x = x + 0.1 {
		y := sigmoid(x)

		points2 = append(points2, plotter.XY{X: x, Y: y})
	}

	// 同时在图表中加入两条折线并起名
	plotutil.AddLinePoints(p, "函数 y = x * x", points1, "函数 y = sigmoid(x)", points2)

	// 保存图表到图片文件
	p.Save(6*vg.Inch, 6*vg.Inch, "points.png")
}
