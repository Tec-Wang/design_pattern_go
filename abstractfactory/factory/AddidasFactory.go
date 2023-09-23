package abstractfactory

import (
	iproduct "design/wangzheng/go/abstractfactory/Iproduct"
	abstractfactory "design/wangzheng/go/abstractfactory/product"
)

type AddidasFactory struct {
}

func (*AddidasFactory) MakeShoes() iproduct.IShoe {
	return &abstractfactory.AddidasShoe{}

}

func (*AddidasFactory) MakeShirts() iproduct.IShirt {
	return &abstractfactory.AddidasShirt{}
}
