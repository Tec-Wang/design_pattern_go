package bridge

import "fmt"

// 惠普打印机
type Hp struct {
}

func (hp *Hp) PrintFile() {
	fmt.Println("惠普打印机打印中。。。。。。")
}
