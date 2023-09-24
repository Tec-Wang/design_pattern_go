package adaptor

import "testing"

func TestAdaptor(t *testing.T) {
	// 客户端
	client := &Client{}
	client.InsertLightningConnectorIntoComputer(&Mac{})

	// windows属于不支持的第三方类，会编译不通过
	// client.InsertLightningConnectorIntoComputer(&Windows{})

	// 使用适配器
	client.InsertLightningConnectorIntoComputer(&WindowsAdaptor{})
}
