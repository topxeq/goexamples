package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// ReturnMsg 用于CustomLabel类型中getTextInGoroutine函数返回数据用
type ReturnMsg struct {
	sync.Mutex // 继承自sync.Mutex，以便可以加共享锁

	Text string // 用于存放实际的结果字符串
}

// CustomLabel 自定义的文字标签图形控件
// 继承自 widgets.QLabel （使用匿名字段的方式）
// 增加声明了一个slot函数changeTextInGoroutine用于修改标签的文字
// 和一个slot函数getTextInGoroutine用于读取标签文字
type CustomLabel struct {
	widgets.QLabel

	_ func(textA string) `slot:"changeTextInGoroutine"`

	// getTextInGoroutine函数会将读取到的标签文字存放在参数t指向的变量中
	_ func(t *ReturnMsg) `slot:"getTextInGoroutine"`
}

// 全局变量runFlagG是用于控制多个 goroutine 是否工作的标志
// runFlagG 为 false 时表示不工作，true 时表示可以有效进行工作
var runFlagG bool = false

// 根据当前时间刷新 label1 标签显示文字的 goroutine
func updateLabel1(labelA *CustomLabel, idxA int) {
	for {
		// 如果runFlagG为false，则休眠1秒钟后继续循环
		// 不进行有效工作
		if !runFlagG {
			time.Sleep(1 * time.Second)

			continue
		}

		// 新建一个ReturnMsg类型的变量msgT用于GetTextInGoroutine函数的返回结果
		msgT := new(ReturnMsg)

		// 读取标签文字时需要加锁，因为要写入msgT中
		msgT.Lock()

		labelA.GetTextInGoroutine(msgT)

		// 读取msgT内容时需要加锁
		msgT.Lock()
		textT := msgT.Text
		msgT.Unlock()

		// 按一定格式显示文字，idxA是goroutine的序号
		// 输出字符串中用下划线分割的第二部分是md5对textT进行编码后的结果，仅保留一个字节
		// 第三部分是当前时间
		labelA.ChangeTextInGoroutine(fmt.Sprintf("%v_%v_%v", idxA, md5.New().Sum([]byte(textT))[0:1], time.Now().Format("2006-01-02 15:04:05.000")))

		time.Sleep(1 * time.Millisecond)
	}
}

// 用随机落点法计算π值的 goroutine
//每隔一定时间刷新 label2 标签显示文字
func updateLabel2(labelA *CustomLabel, anotherLabelA *CustomLabel) {

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

		// 获取label1的文字
		msgT := new(ReturnMsg)

		msgT.Lock()

		anotherLabelA.GetTextInGoroutine(msgT)

		// 调用labelA的slot函数设置该标签的文字为正在计算的π值
		msgT.Lock()
		textT := msgT.Text
		msgT.Unlock()

		// 选取分割后第一个部分，即序号部分
		textT = strings.Split(textT, "_")[0]

		// 显示所取的updateLabel1函数启动的goroutine序号以及当前计算的Π值
		labelA.ChangeTextInGoroutine(fmt.Sprintf("[%v]计算的π值：%v", textT, piT))

		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	appT := widgets.NewQApplication(len(os.Args), os.Args)

	// 新建主窗口
	windowT := widgets.NewQMainWindow(nil, 0)
	windowT.SetMinimumSize2(250, 200)
	windowT.SetWindowTitle("并发安全测试")

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

	// 第一个标签，用于显示当前时间
	label1 := NewCustomLabel(nil, 0)
	label1.SetText("")
	label1.SetFixedHeight(100)
	label1.SetFixedWidth(220)

	// 设置label1的slot处理函数ChangeTextInGoroutine
	// 将label1的文字设置为传入的textA参数
	label1.ConnectChangeTextInGoroutine(func(textA string) {
		label1.SetText(textA)
	})

	// 设置label1的slot处理函数GetTextInGoroutine
	// 读出label1的文字放入msgA参数的Text字段中
	label1.ConnectGetTextInGoroutine(func(msgA *ReturnMsg) {
		msgA.Text = label1.Text()

		msgA.Unlock()
	})

	widget2T.Layout().AddWidget(label1)

	// 创建字体对象，用于更改label2的字体
	fontT := gui.NewQFont2("Helvetica", -1, -1, false)
	fontT.SetPointSize(16) // 设置字体的大小

	// 第二个标签，用于表示计算中的π值
	label2 := NewCustomLabel(nil, 0)
	label2.SetText("")
	label2.SetFixedHeight(100)
	label2.SetFixedWidth(400)
	label2.SetFont(fontT) // 设置所用的字体

	// 设置label2的slot处理函数
	// 将label2的文字设置为传入的textA参数
	label2.ConnectChangeTextInGoroutine(func(textA string) {
		label2.SetText(textA)
	})

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

	// 启动5个updateLabel1的goroutine
	go updateLabel1(label1, 0)
	go updateLabel1(label1, 1)
	go updateLabel1(label1, 2)
	go updateLabel1(label1, 3)
	go updateLabel1(label1, 4)

	// 启动1个updateLabel1的goroutine
	go updateLabel2(label2, label1)

	windowT.Show()

	appT.Exec()
}
