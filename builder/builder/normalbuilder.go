package builder

import house "design/wangzheng/go/builder/product"

type NormalBulder struct {
	Builder
	house.NormalHouse
}

func (builder *NormalBulder) Reset() {
	// 实际上在go语言中没有必要。
	builder.NormalHouse = house.NormalHouse{}
}

func (builder *NormalBulder) SetWindowType() {
	builder.WindowType = "normal Window"
}
func (builder *NormalBulder) SetDoorType() {
	builder.DoorType = "normal Door"
}
func (builder *NormalBulder) SetNumFloor() {
	builder.Float = 2
}

func (builder *NormalBulder) SetPublicArea() {
	builder.PublicArea = 80
}

func (builder *NormalBulder) GetHouse() house.NormalHouse {
	return house.NormalHouse{
		WindowType: builder.WindowType,
		DoorType:   builder.DoorType,
		Float:      builder.Float,
		PublicArea: builder.PublicArea,
	}
}
