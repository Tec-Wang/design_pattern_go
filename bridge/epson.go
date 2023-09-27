package bridge

import "fmt"

// Epson打印机，具体实施层
type Epson struct {
}

func (epson *Epson) PrintFile() {
	fmt.Println("epson打印机打印中。。。。。。")
}
