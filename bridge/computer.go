package bridge

// 电脑抽象类，具有打印功能，打印功能需要兼容不同品牌的打印机
type Computer interface {
	Print()
	SetPrinter(Printer)
}
