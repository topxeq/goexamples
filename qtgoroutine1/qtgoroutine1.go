package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// 用于控制两个 goroutine 是否工作的标志
// runFlagG 为 false 时表示不工作，true 时表示可以有效进行工作
var runFlagG bool = false

// 根据当前时间刷新 label1 标签显示文字的 goroutine
func updateLabel1(labelA *widgets.QLabel, idxA int) {
	for {
		// 如果runFlagG为false，则休眠1秒钟后继续循环
		// 不进行有效工作
		if !runFlagG {
			time.Sleep(1 * time.Second)

			continue
		}

		// 获取第一个标签的文本
		textT := labelA.Text()

		// 获取textT的md5编码，并与序号和时间显示到第一个标签中
		// 即设置labelA的文字为“序号_MD5编码的第一个字符_时间”

		labelA.SetText(fmt.Sprintf("%v_%v_%v", idxA, md5.New().Sum([]byte(textT))[0:1], time.Now().Format("2006-01-02 15:04:05.000")))

		// 休眠1毫秒，防止系统资源占用过高
		time.Sleep(1 * time.Millisecond)

	}
}

// 用随机落点法计算π值的 goroutine
//每隔一定时间刷新 label2 标签显示文字
func updateLabel2(labelA *widgets.QLabel, anotherLabelA *widgets.QLabel) {

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

		// 获取第一个标签的文字
		textT := anotherLabelA.Text()

		// 获取其中第一项内容（即goroutine序号）
		textT = strings.Split(textT, "_")[0]

		// 设置第二个标签的文字为第一个标签的goroutine序号和Π值
		labelA.SetText(fmt.Sprintf("[%v]计算的π值：%1.5f", textT, piT))

		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	appT := widgets.NewQApplication(len(os.Args), os.Args)

	// 新建主窗口
	windowT := widgets.NewQMainWindow(nil, 0)
	windowT.SetMinimumSize2(250, 200)
	windowT.SetWindowTitle("并发测试")

	// 新建一个widget，用于包含两个纵向排列的子widget
	widget1T := widgets.NewQWidget(nil, 0)
	widget1T.SetLayout(widgets.NewQVBoxLayout())
	windowT.SetCentralWidget(widget1T)

	// widget1T中纵向排列的第一个widget
	// 用于包含两个CustomLabel标签
	widget2T := widgets.NewQWidget(nil, 0)
	widget2T.SetLayout(widgets.NewQHBoxLayout())
	widget1T.Layout().AddWidget(widget2T)

	// widget1T中纵向排列的第二个widget
	// 用于包含按钮buttonT
	widget3T := widgets.NewQWidget(nil, 0)
	widget3T.SetLayout(widgets.NewQHBoxLayout())
	widget1T.Layout().AddWidget(widget3T)

	// 第一个标签，用于显示当前时间等信息
	label1 := widgets.NewQLabel(nil, 0)
	label1.SetText("")
	label1.SetFixedHeight(100)
	label1.SetFixedWidth(220) // 设置固定的宽度

	widget2T.Layout().AddWidget(label1)

	// 创建字体对象，用于更改label2的字体
	fontT := gui.NewQFont2("Helvetica", -1, -1, false)
	fontT.SetPointSize(16) // 设置字体的大小

	// 第二个标签，用于表示计算中的π值等信息
	label2 := widgets.NewQLabel(nil, 0)
	label2.SetText("")
	label2.SetFixedHeight(100)
	label2.SetFixedWidth(360)
	label2.SetFont(fontT) // 设置所用的字体

	widget2T.Layout().AddWidget(label2)

	// 用于控制两个线程运行与否的按钮
	buttonT := widgets.NewQPushButton2("开始", nil)

	// 设置buttonT按钮处理点击事件的函数
	buttonT.ConnectClicked(func(bool) {

		// 切换runFlagG的状态
		// runFlagG控制了goroutine是否可以有效运行
		// true表示可以运行，false代表不应工作
		// 从另一个角度说也代表了goroutine是否正在运行
		runFlagG = !runFlagG

		if runFlagG {
			// 如果goroutine正在运行，设置按钮buttonT的文字为“停止”
			buttonT.SetText("停止")
		} else {
			// 如果goroutine不在有效运行，设置按钮buttonT的文字为“开始”
			buttonT.SetText("开始")
		}

	})

	widget3T.Layout().AddWidget(buttonT)

	// 运行5个updateLabel1函数生成的goroutine
	go updateLabel1(label1, 0)
	go updateLabel1(label1, 1)
	go updateLabel1(label1, 2)
	go updateLabel1(label1, 3)
	go updateLabel1(label1, 4)

	// 运行1个updateLabel2函数生成的goroutine
	go updateLabel2(label2, label1)

	windowT.Show()

	appT.Exec()
}
