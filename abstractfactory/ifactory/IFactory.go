package abstractfactory

import (
	iproduct "design/wangzheng/go/abstractfactory/Iproduct"
	factory "design/wangzheng/go/abstractfactory/factory"
	"fmt"
)

// 抽象工厂，生产鞋子和衬衫
type IFactory interface {
	MakeShoes() iproduct.IShoe
	MakeShirts() iproduct.IShirt
}

// 入口
func GetFactory(brandName string) (IFactory, error) {
	switch brandName {
	case "Addidas":
		return &factory.AddidasFactory{}, nil
	case "Nike":
		return &factory.NikeFactory{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")

	}

}
