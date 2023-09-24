package adaptor

import (
	"fmt"
)

// 适配器需要实现Client调用的Computer的接口的方法
// 并且持有一个被转换的对象
type WindowsAdaptor struct {
	windows Windows
}

func (adaptor *WindowsAdaptor) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	adaptor.windows.InsertUSBConnectorIntoComputer()
}
