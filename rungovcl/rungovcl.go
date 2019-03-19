package main

import (
	"encoding/hex"
	"strings"

	"github.com/ying32/govcl/vcl"
)

// TMainForm 是代表主窗体的自定义结构类型
type TMainForm struct {
	*vcl.TForm // 应继承自vcl.TForm

	// 各个界面控件的成员变量
	Button1 *vcl.TButton
	Label1  *vcl.TLabel
	Label2  *vcl.TLabel
	Edit1   *vcl.TEdit
	Edit2   *vcl.TEdit
}

// 设置表示主窗体实例的全局变量mainForm
var (
	mainForm *TMainForm
)

func main() {

	// 主函数中的常规操作
	vcl.Application.Initialize()

	vcl.Application.SetMainFormOnTaskBar(true)

	vcl.Application.CreateForm(&mainForm)

	vcl.Application.Run()
}

// OnFormCreate 函数会在主窗体创建时被首先调用
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("第一个例子")       // 设置主窗体标题
	f.SetBounds(0, 0, 320, 240) // 设置主窗体大小

	// 设置窗体的字体样式，对窗体内所有控件有效
	// 除非该控件自己再设置字体
	f.Font().SetName("Simhei")
	f.Font().SetSize(11)

	// 设置第一个标签控件
	f.Label1 = vcl.NewLabel(f) // 新建标签对象

	f.Label1.SetParent(f)              // 设置其父控件为主窗体
	f.Label1.SetBounds(20, 20, 50, 28) // 设置位置和大小，参数依次为左上角的X、Y坐标及宽、高
	f.Label1.SetCaption("原文本")         // 设置标签文字

	// 设置第一个文本编辑框控件
	f.Edit1 = vcl.NewEdit(f)
	f.Edit1.SetParent(f)
	f.Edit1.SetBounds(120, 20, 160, 28)

	// 设置第二个标签控件
	f.Label2 = vcl.NewLabel(f)
	f.Label2.SetParent(f)
	f.Label2.SetBounds(20, 60, 50, 28)
	f.Label2.SetCaption("编码后文本")

	// 设置第二个文本编辑框控件
	f.Edit2 = vcl.NewEdit(f)
	f.Edit2.SetParent(f)
	f.Edit2.SetBounds(120, 60, 160, 28)

	// 设置按钮控件
	f.Button1 = vcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetBounds(100, 108, 48, 28)
	f.Button1.SetCaption("编码")

	// 设置按钮控件的点击事件处理函数
	f.Button1.SetOnClick(f.OnButtonClick)

}

// OnButtonClick 是按钮Button1点击事件的处理函数
func (f *TMainForm) OnButtonClick(sender vcl.IObject) {

	// 获取文本编辑框Edit1中的文本
	strT := f.Edit1.Text()

	// 如果文本去除两边空白后是空字符串则弹出提示信息后退出
	if strings.TrimSpace(strT) == "" {
		vcl.ShowMessage("原文本为空字符串")
		return
	}

	// 将该文本编码为十六进制文本后显示在Edit2中
	f.Edit2.SetText(hex.EncodeToString([]byte(strT)))
}
