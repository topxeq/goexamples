package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"

	t "tools"
)

var textT = []string{
	"三人行，必有我师",
	"得道多助，失道寡助",
	"飘风不到夕，暴雨不终朝",
	"",
	"We are so nice.",
}

func main() {

	// 打开一个已有图片
	fileT, errT := os.Open(`c:\test\rgb.png`)

	if errT != nil {
		t.Printfln("打开图片文件时发生错误：%v", errT.Error())
		return
	}

	defer fileT.Close()

	imgT, errT := png.Decode(fileT)
	if errT != nil {
		t.Printfln("图片解码时发生错误：%v", errT.Error())
		return
	}

	// 要从image.Image对象强制转换为image.RGBA图像对象
	imageT := imgT.(*image.RGBA)

	// 载入字体文件内容
	fontBytesT, errT := ioutil.ReadFile(`C:\Windows\Fonts\simhei.ttf`)

	if errT != nil {
		t.Printfln("载入字体时发生错误：%v", errT.Error())
		return
	}

	// 解析字体
	fontT, errT := freetype.ParseFont(fontBytesT)

	if errT != nil {
		t.Printfln("分析字体时发生错误：%v", errT.Error())
		return
	}

	// 设置前景色，即绘制文字用的颜色
	foreColorT := image.White

	// 设置字体大小
	fontSizeT := 28.0

	// 绘制字体需要创建一个freetype.Context类型的环境对象
	contextT := freetype.NewContext()

	// 设置DPI
	contextT.SetDPI(72)

	// 设置使用的字体
	contextT.SetFont(fontT)

	// 设置字体大小
	contextT.SetFontSize(fontSizeT)

	// 设置绘制文字的区域
	contextT.SetClip(imageT.Bounds())

	// 设置在哪个图形对象上绘制
	contextT.SetDst(imageT)

	// 设置绘制源，这里直接用绘制文字的颜色传入即可
	contextT.SetSrc(foreColorT)

	// 设置绘制文字的位置pt
	// PointToFixed函数将相对文字大小pt转换为实际像素大小
	// 因为PointToFixed函数返回的是特殊的freetype.Int26_6类型的数值
	// 如果转换成整数需要右移六个二进制位
	pt := freetype.Pt(10, 10+int(contextT.PointToFixed(fontSizeT)>>6))

	for _, s := range textT {
		// 实际绘制文字的语句
		_, errT = contextT.DrawString(s, pt)
		if errT != nil {
			t.Printfln("绘制文字时发生错误：%v", errT.Error())
			return
		}

		// 每次画完文字将pt的纵坐标下移1.5倍字体大小
		pt.Y += contextT.PointToFixed(fontSizeT * 1.5)
	}

	// 保存图像为PNG格式的图片文件
	fileT, errT = os.Create("c:\\test\\font.png")

	if errT != nil {
		fmt.Println(errT)
		return
	}

	defer fileT.Close()

	png.Encode(fileT, imageT)
}
