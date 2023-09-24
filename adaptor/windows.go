package adaptor

import "fmt"

type Windows struct {
}

// Windows 沒有实现Computer接口的Lighting方法需要一个适配器
func (com *Windows) InsertUSBConnectorIntoComputer() {
	fmt.Print("将 USB 插入了 windows")
}
