package main

import (
	"encoding/hex"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	// 创建一个图形窗口应用
	appT := widgets.NewQApplication(len(os.Args), os.Args)

	// 创建一个最小为 250*200 大小的主窗口
	// 并设置窗口标题为 "十六进制编码"
	windowT := widgets.NewQMainWindow(nil, 0)
	windowT.SetMinimumSize2(250, 200)
	windowT.SetWindowTitle("十六进制编码")

	// 创建一个 widget
	// 使用QVBoxLayout布局，即纵向顺序排放控件的布局
	// 并使其居中
	widgetT := widgets.NewQWidget(nil, 0)
	widgetT.SetLayout(widgets.NewQVBoxLayout())
	windowT.SetCentralWidget(widgetT)

	// 在上述widget中创建两个纯文本编辑框
	// 编辑框内都有提示文本
	input1 := widgets.NewQPlainTextEdit(nil)
	input1.SetPlaceholderText("在此输入要编码的文本……")
	input1.SetFixedHeight(100)
	widgetT.Layout().AddWidget(input1)

	input2 := widgets.NewQPlainTextEdit(nil)
	input2.SetPlaceholderText("此处将显示编码后的十六进制文本")
	input2.SetFixedHeight(100)
	widgetT.Layout().AddWidget(input2)

	// 再创建一个按钮并设置点击事件的处理函数
	buttonT := widgets.NewQPushButton2("编码", nil)

	// 设置buttonT按钮处理点击事件的函数
	buttonT.ConnectClicked(func(bool) {
		// 将input2编辑框中的文字设置成input1编辑框中文字十六进制编码后的文本
		input2.SetPlainText(hex.EncodeToString([]byte(input1.ToPlainText())))

		// 弹出消息框提示操作成功
		widgets.QMessageBox_Information(nil, "OK", "编码完毕", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	widgetT.Layout().AddWidget(buttonT)

	// 显示主窗口
	windowT.Show()

	// 运行该图形应用程序
	// 直至遇到调用app.Exit()函数
	// 或者用户关闭该程序的主窗口
	appT.Exec()
}
