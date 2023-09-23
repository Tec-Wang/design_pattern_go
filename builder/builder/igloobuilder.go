package builder

import house "design/wangzheng/go/builder/product"

type IglooBuilder struct {
	Builder
	house.IglooHouse
}

func (builder *IglooBuilder) Reset() {
	builder.IglooHouse = house.IglooHouse{}
}

func (builder *IglooBuilder) SetWindowType() {
	builder.WindowType = "igloo Window"
}
func (builder *IglooBuilder) SetDoorType() {
	builder.DoorType = "igloo Door"
}
func (builder *IglooBuilder) SetNumFloor() {
	builder.Float = 2
}
func (builder *IglooBuilder) GetHouse() house.IglooHouse {
	return house.IglooHouse{
		WindowType: builder.WindowType,
		DoorType:   builder.DoorType,
		Float:      builder.Float,
	}
}
