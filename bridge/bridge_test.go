package bridge

import "testing"

func TestBridge(t *testing.T) {
	// 使用抽象层
	var computer Computer
	var printer Printer

	{
		// mac & hp
		computer = &Mac{}
		printer = &Hp{}
	}

	computer.SetPrinter(printer)
	computer.Print()

	{
		// windows & hp
		computer = &Windows{}
		printer = &Hp{}
	}

	computer.SetPrinter(printer)
	computer.Print()

	{
		// mac & Epson
		computer = &Mac{}
		printer = &Epson{}
	}

	computer.SetPrinter(printer)
	computer.Print()

	{
		// Windows & Epson
		computer = &Windows{}
		printer = &Epson{}
	}

	computer.SetPrinter(printer)
	computer.Print()
}

// 如果想要增加功能
// 分别增加抽象层和实施层。
