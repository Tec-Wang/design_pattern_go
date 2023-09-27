package bridge

import "fmt"

// 精确抽象类，具有打印功能
type Mac struct {
	Printer Printer
}

func (mac *Mac) SetPrinter(printer Printer) {
	mac.Printer = printer
}

func (mac *Mac) Print() {
	fmt.Print("mac 打印文件 :")
	mac.Printer.PrintFile()
	fmt.Println("mac 打印完成！")
	fmt.Println("-------------")
}
