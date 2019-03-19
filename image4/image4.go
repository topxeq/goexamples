package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// TwoCircleMask 表示内含两个交叠圆形的蒙版对象
// W、H分别是蒙版的宽度和高度
// (X0, Y0)、(X1, Y1)是分别是两个圆心的坐标，R是半径
type TwoCircleMask struct {
	W, H           int
	X0, Y0, X1, Y1 float64
	R              float64
}

func (p *TwoCircleMask) ColorModel() color.Model {
	return color.AlphaModel
}

func (p *TwoCircleMask) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.W, p.H)
}

func (p *TwoCircleMask) At(x, y int) color.Color {
	// 计算该点与两个圆心的相对坐标 (dx1, dy1)和(dx2, dy2)
	// 注意这种同时给两个变量赋值的方法
	var dx1, dy1 float64 = p.X0 - float64(x), p.Y0 - float64(y)
	var dx2, dy2 float64 = p.X1 - float64(x), p.Y1 - float64(y)

	// 计算该点与两个圆心的距离 d1和d2
	d1 := math.Sqrt(dx1*dx1 + dy1*dy1)
	d2 := math.Sqrt(dx2*dx2 + dy2*dy2)

	// 判断 (x, y) 点是否在形状内
	if d1 <= p.R || d2 <= p.R {
		// 如果 距离 d1 和 d2 中的任意一个不大于半径，则该点一定在形状之内
		// 使用一定的算法规则实现渐变的蒙板（即透明度渐变）
		return color.Alpha{byte(255 * (d1 + d2) / 2 / p.R)}
	}

	// 否则该点在图形外
	return color.Alpha{0}
}

func main() {

	// 设置图片的宽度和高度
	widthT, heightT := 640, 480

	// 创建原图与目标图
	srcImageT := image.NewRGBA(image.Rect(0, 0, widthT, heightT))
	dstImageT := image.NewRGBA(image.Rect(0, 0, widthT, heightT))

	// 将原图填充成红色
	draw.Draw(srcImageT, srcImageT.Bounds(), image.NewUniform(color.RGBA{255, 0, 0, 255}), image.ZP, draw.Src)

	//将目标图填充成黑色
	draw.Draw(dstImageT, srcImageT.Bounds(), image.Black, image.ZP, draw.Src)

	// 将原图通过蒙版截取的不规则区域复制到目标图片上
	draw.DrawMask(dstImageT, dstImageT.Bounds(), srcImageT, image.ZP, &TwoCircleMask{W: widthT, H: heightT, X0: 200, Y0: 200, X1: 300, Y1: 200, R: 80}, image.ZP, draw.Over)

	// 保存图片
	fileT, errT := os.Create(`mask2.png`)

	if errT != nil {
		fmt.Println(errT)
		return
	}

	defer fileT.Close()

	png.Encode(fileT, dstImageT)
}
