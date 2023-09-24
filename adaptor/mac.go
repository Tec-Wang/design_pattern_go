package adaptor

import "fmt"

type Mac struct {
}

func (com *Mac) InsertIntoLightingPort() {
	fmt.Print("将 lighting 插入了 MAC")
}
