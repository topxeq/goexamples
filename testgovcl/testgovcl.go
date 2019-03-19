package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ying32/govcl/vcl"
)

// TMainForm 是代表主窗体的自定义结构类型
type TMainForm struct {
	*vcl.TForm // 应继承自vcl.TForm

	// 各个界面控件的成员变量
	Button1 *vcl.TButton
	Label1  *vcl.TLabel
	Label2  *vcl.TLabel
}

// 设置表示主窗体实例的全局变量mainForm
var (
	mainForm *TMainForm
)

// 用于控制两个 goroutine 是否工作的标志
// runFlagG 为 false 时表示不工作，true 时表示可以有效进行工作
var runFlagG bool = false

// 根据当前时间刷新 Label1 标签显示文字的 goroutine
func updateLabel1(idxA int) {
	for {
		// 如果runFlagG为false，则休眠1秒钟后继续循环
		// 不进行有效工作
		if !runFlagG {
			time.Sleep(1 * time.Second)

			continue
		}

		// 用vcl.ThreadSync来执行所需的操作界面的代码
		vcl.ThreadSync(func() {
			// 获取Label1的文字
			textT := mainForm.Label1.Caption()

			// 按一定格式显示文字，idxA是goroutine的序号
			// 输出字符串中用下划线分割的第二部分是md5对textT进行编码后的结果
			// 第三部分是当前时间
			mainForm.Label1.SetCaption(fmt.Sprintf("%v_%v_%v", idxA, md5.New().Sum([]byte(textT))[0:1], time.Now().Format("2006-01-02 15:04:05")))
		})

		time.Sleep(1 * time.Millisecond)
	}
}

// 用随机落点法计算π值的 goroutine
//每隔一定时间刷新 label2 标签显示文字
func updateLabel2() {

	// 用随机落点法计算圆周率π
	inCircleCount := 0
	pointCountT := 0

	var x, y float64
	var piT float64

	for {
		if !runFlagG {
			time.Sleep(1 * time.Second)
			continue
		}

		pointCountT++

		x = rand.Float64()
		y = rand.Float64()

		if x*x+y*y < 1 {
			inCircleCount++
		}

		piT = (4.0 * float64(inCircleCount)) / float64(pointCountT)

		vcl.ThreadSync(func() {

			// 获取Label1的文字
			textT := mainForm.Label1.Caption()

			// 选取分割后第一个部分，即序号部分
			textT = strings.Split(textT, "_")[0]

			// 显示所取的updateLabel1函数启动的goroutine序号以及当前计算的Π值
			mainForm.Label2.SetCaption(fmt.Sprintf("[%v]计算的π值：%v", textT, piT))
		})

		time.Sleep(1 * time.Millisecond)
	}
}

// OnButtonClick 是按钮Button1点击事件的处理函数
func (f *TMainForm) OnButtonClick(sender vcl.IObject) {

	// 切换runFlagG的状态
	runFlagG = !runFlagG

	if runFlagG {
		// 如果goroutine正在运行，设置按钮Button1的文字为“停止”
		f.Button1.SetCaption("停止")
	} else {
		// 如果goroutine不在有效运行，设置按钮Button1的文字为“开始”
		f.Button1.SetCaption("开始")
	}
}

// OnFormCreate 函数会在主窗体创建时被首先调用
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("并发测试")        // 设置主窗体标题
	f.SetBounds(0, 0, 360, 240) // 设置主窗体大小

	// 设置默认字体样式，对窗体内所有控件有效
	// 除非该控件自己再设置字体
	f.Font().SetName("Simhei")
	f.Font().SetSize(11)

	// 设置第一个标签控件
	f.Label1 = vcl.NewLabel(f) // 新建标签对象

	f.Label1.SetParent(f)               // 设置其父控件为主窗体
	f.Label1.SetBounds(20, 20, 240, 28) // 设置位置和大小，参数依次为左上角的X、Y坐标及宽、高
	f.Label1.SetCaption("")             // 设置标签文字

	// 设置第二个标签控件
	f.Label2 = vcl.NewLabel(f)
	f.Label2.SetParent(f)
	f.Label2.SetBounds(20, 60, 240, 28)
	f.Label2.SetCaption("")

	// 设置按钮控件
	f.Button1 = vcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetBounds(100, 108, 48, 28)
	f.Button1.SetCaption("开始")

	// 设置按钮控件的点击事件处理函数
	f.Button1.SetOnClick(f.OnButtonClick)

	// 启动5个updateLabel1的goroutine
	go updateLabel1(1)
	go updateLabel1(2)
	go updateLabel1(3)
	go updateLabel1(4)
	go updateLabel1(5)

	// 启动1个updateLabel1的goroutine
	go updateLabel2()

}

func main() {

	// 主函数中的常规操作
	vcl.Application.Initialize()

	vcl.Application.SetMainFormOnTaskBar(true)

	vcl.Application.CreateForm(&mainForm)

	vcl.Application.Run()
}
