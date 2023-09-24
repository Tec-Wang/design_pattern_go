package adaptor

import "fmt"

type Client struct {
}

func (*Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightingPort()
}
