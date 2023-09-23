package abstractfactory

import (
	iproduct "design/wangzheng/go/abstractfactory/Iproduct"
	product "design/wangzheng/go/abstractfactory/product"
)

type NikeFactory struct {
}

func (*NikeFactory) MakeShoes() iproduct.IShoe {
	return &product.NikeShoe{}
}

func (*NikeFactory) MakeShirts() iproduct.IShirt {
	return &product.NikeShirt{}
}
