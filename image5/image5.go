package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
	"strings"

	"github.com/fogleman/gg"
)

// Circle 表示圆形的对象，(X0, Y0)是圆心的坐标，R是半径
type Circle struct {
	X0, Y0, R float64
}

// GetColorValue 确定坐标为 (x, y) 的点的色彩
func (p *Circle) GetColorValue(x, y float64) byte {

	// 计算该点与圆心的相对坐标 (dx, dy)
	// 注意这种同时给两个变量赋值的方法
	var dx, dy float64 = p.X0 - x, p.Y0 - y

	// 计算该点与圆心的距离 d
	d := math.Sqrt(dx*dx + dy*dy)

	if d > p.R {
		// 如果 距离 d 大于半径，则该点一定在圆形之外
		// 返回0值，表示没有该种色调
		return 0
	}

	// 否则该点在圆形内部（包括边）
	// 此时根据与圆心距离的远近确定颜色强度
	return byte(255.0 * d / p.R)
}

// Square 表示正方形的对象
// (X0, Y0)是其中心点的坐标，R是半径，即中心点到任何一条边的垂直距离
type Square struct {
	X0, Y0, R float64
}

// InShape 确定坐标为 (x, y) 的点是否在该形状内
func (p *Square) InShape(x, y float64) bool {

	// 计算正方形上方那条边的纵坐标yt和下方那条边的纵坐标yb
	yt := p.Y0 - p.R
	yb := p.Y0 + p.R

	// 计算正方形左侧那条边的横坐标xl和右侧那条边的横坐标xr
	xl := p.X0 - p.R
	xr := p.X0 + p.R

	// 如果 y > yb 则该点不在形状内
	if y > yb {
		return false
	}

	// 如果 y < yt 则该点不在形状内
	if y < yt {
		return false
	}

	// 如果 x < xl 则该点不在形状内
	if x < xl {
		return false
	}

	// 如果 x > xr 则该点不在形状内
	if x > xr {
		return false
	}

	// 默认返回true，表示该点在形状内
	return true

}

// GetColorValue 确定坐标为 (x, y) 的点的色彩
func (p *Square) GetColorValue(x, y float64) byte {
	// 如果不在该形状内，则色彩强度为0
	if !p.InShape(x, y) {
		return 0
	}

	// 对x，y坐标取模后用一定公式确定色彩强度
	return byte((int(x)%5)*32 + (int(y)%5)*32)
}

// Triangle 表示等边三角形的对象，金字塔型对称放置
// (X0, Y0)是其中心点的坐标，R是半径，即中心点到任何一个顶点的距离
type Triangle struct {
	X0, Y0, R float64
}

// InShape 确定坐标为 (x, y) 的点是否在该形状内
func (p *Triangle) InShape(x, y float64) bool {

	// 求得该三角形上顶点的纵坐标yt和两个下顶点的纵坐标yb
	yt := p.Y0 - p.R
	yb := p.Y0 + p.R/2

	// 如果 y 大于 yb，则该点在形状外
	if y > yb {
		return false
	}

	// 如果 y 小于 yt，则该点在形状外
	if y < yt {
		return false
	}

	// 求该等边三角形的边长
	sideLength := math.Sqrt(3) * p.R

	// x坐标在中心点左侧时与右侧时需要用相反的计算方式确定该点是否在形状内
	if x < p.X0 {

		// 用矢量方向法判断该点是在三角形左侧那条边的左边还是右边
		x1 := p.X0
		y1 := yt

		x2 := p.X0 - sideLength/2
		y2 := yb

		x3 := x
		y3 := y

		s := (x1-x3)*(y2-y3) - (y1-y3)*(x2-x3)

		if s <= 0 {
			// s <= 0 表示在右侧，即在形状内
			return true
		}

	} else {

		// 用矢量方向法判断该点是在三角形右侧那条边的左边还是右边
		x1 := p.X0
		y1 := yt

		x2 := p.X0 + sideLength/2
		y2 := yb

		x3 := x
		y3 := y

		s := (x1-x3)*(y2-y3) - (y1-y3)*(x2-x3)

		if s > 0 {
			// s > 0 表示在左侧，即在形状内
			return true
		}
	}

	// 默认返回false，表示该点不在形状内
	return false

}

// 例子网页代码，用于对网站根目录访问时作为响应内容返回
var htmlTextG = `
<html>
<head>
	<meta charset="utf-8" />
	<title>Image test</title>
</head>
<body>
	<div><center><span>欢迎</span></center></div>
	<div style="margin-top: 20px;"><center><img src="data:image/png;base64,这里的中文字符串将被替换" width=600px /></center></div>
</body>
</html>

	`

// 处理对网站根目录进行访问的HTTP请求的函数
// 将htmlTextG字符串中的子串“这里的中文字符串将被替换”替换为Base64编码的图片数据后作为网页返回
func handlePage(respA http.ResponseWriter, reqA *http.Request) {

	// 确定图片的宽与高
	// 注意这种同时声明两个同类型变量并用一个等号分别赋值的写法
	var w, h int = 600, 400

	// 为计算三个形状的圆心准备数据
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 80.0
	θ := math.Pi * 2 / 3

	// 生成三个形状，第一个形状用红色，第二个用绿色，第三个用蓝色填充
	//第一个形状用三角形，第二个用方形，第三个用原型
	// 处于照顾视觉的考虑，适当调整了几个形状的中心位置和半径
	shapeRedT := &Triangle{hw - r*math.Sin(0) + 90*0.1, hh - r*math.Cos(0) + 90*0.3, 90 * 1.18}
	shapeGreenT := &Square{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 90 * 0.9}
	shapeBlueT := &Circle{(hw - r*math.Sin(-θ)), hh - r*math.Cos(-θ), 90}

	// 新建一个RGBA色彩空间的图形对象
	imageT := image.NewRGBA(image.Rect(0, 0, w, h))

	// 循环一行一行设置每个点的颜色
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// 根据该点是否在shapeRedT内确定红色的强度
			var colorRedT byte = 0
			if shapeRedT.InShape(float64(x), float64(y)) {
				colorRedT = 255
			}

			// 根据该点 x、y坐标的奇偶数确定绿色的强度
			var colorGreenT byte = shapeGreenT.GetColorValue(float64(x), float64(y))

			// 根据该点是否在shapeBlueT内确定绿色的强度
			var colorBlueT byte = shapeBlueT.GetColorValue(float64(x), float64(y))

			// 根据汇总的RGB颜色确定最终该点的颜色
			colorT := color.RGBA{colorRedT, colorGreenT, colorBlueT, 255}

			// 设置该点的颜色
			imageT.Set(x, y, colorT)
		}
	}

	// 从image.RGBA类型的图片对象生成一个gg.Context
	// 以便准备在其上绘制文字
	contextT := gg.NewContextForImage(imageT)

	// 设置使用白色画笔
	contextT.SetColor(image.White)

	// 在图中下方居中输出包括访问端地址的文字信息
	contextT.DrawStringAnchored(fmt.Sprintf("Welcome user from %v", reqA.RemoteAddr), 300, 350, 0.5, 0.5)

	// 实际执行绘制文字操作
	contextT.Stroke()

	// 定义字节缓冲区用于接收图片编码成PNG格式的数据
	var bufT bytes.Buffer

	// 进行png格式的图形编码，并以流式方法写入http响应中
	// 编码前要用gg.Context.Image函数获取转换回image.Image类型的
	png.Encode(&bufT, contextT.Image())

	// 将图片数据编码成Base64格式的文本字符串
	base64ImageT := base64.StdEncoding.EncodeToString(bufT.Bytes())

	// 写请求响应头
	respA.Header().Set("Content-Type", "text/html")

	respA.WriteHeader(200)

	// 将htmlTextG字符串中的子串“这里的中文字符串将被替换”替换为Base64编码的图片数据后作为网页返回
	respA.Write([]byte(strings.Replace(htmlTextG, "这里的中文字符串将被替换", base64ImageT, -1)))

}

func main() {

	// 设定根路由处理函数
	http.HandleFunc("/", handlePage)

	// 在指定端口上监听
	http.ListenAndServe(":8837", nil)
}
