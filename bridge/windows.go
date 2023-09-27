package bridge

import "fmt"

// 精确抽象类，具有打印功能
type Windows struct {
	Printer Printer
}

func (windows *Windows) SetPrinter(printer Printer) {
	windows.Printer = printer
}

func (windows *Windows) Print() {
	fmt.Print("windows 打印文件 :")
	windows.Printer.PrintFile()
	fmt.Println("windows 打印完成！")
	fmt.Println("-------------")
}
