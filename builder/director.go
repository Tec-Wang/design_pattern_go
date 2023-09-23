package builder

import (
	b "design/wangzheng/go/builder/builder"
	house "design/wangzheng/go/builder/product"
)

type Director struct{}

// 客户端持有的对象，是个入口
func (director *Director) MakeNormalHouse(normalBulder *b.NormalBulder) house.NormalHouse {
	normalBulder.Reset()
	normalBulder.SetDoorType()
	normalBulder.SetNumFloor()
	normalBulder.SetPublicArea()
	normalBulder.SetWindowType()
	return normalBulder.GetHouse()
}

func (director *Director) MakeIGlooHouse(iglooBuilder *b.IglooBuilder) house.IglooHouse {
	iglooBuilder.Reset()
	iglooBuilder.SetDoorType()
	iglooBuilder.SetNumFloor()
	iglooBuilder.SetWindowType()
	return iglooBuilder.GetHouse()
}
