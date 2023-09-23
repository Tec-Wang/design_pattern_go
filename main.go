package main

import (
	ifactory "design/wangzheng/go/abstractfactory/ifactory"
	builder "design/wangzheng/go/builder"
	housebuilder "design/wangzheng/go/builder/builder"
	"runtime"

	factory "design/wangzheng/go/factory"
	"fmt"
)

func main() {
	// factory
	// factory_test()
	// test_absFactory()
	test_bulder()

	// 纯纯测试代码
	// 数组是值类型
	ints := [...]int{1, 2, 3}
	a := ints
	a[0] = 222
	fmt.Println(a)
	fmt.Println(ints)
}

func String2Bytes(s string) []byte {
	runtime.SetFinalizer()
}

func factory_test() {
	factory := factory.IFactory{}

	// ak := factory.GetGun("ak")
	// fmt.Printf("ak.Shoot(): %v\n", ak.Shoot())
	// de := factory.GetGun("desertEagle")
	// fmt.Printf("de.Shoot(): %v\n", de.Shoot())

	// 获取十次ak
	for i := 0; i < 10; i++ {
		i2 := factory.GetGun("ak")
		fmt.Println(&i2)
	}
}

func test_absFactory() {
	brandName := "Addidas"
	clothesFactory, error := ifactory.GetFactory(brandName)
	if error != nil {
		fmt.Errorf("wrong when get factory")
	}
	shirts := clothesFactory.MakeShirts()
	shoes := clothesFactory.MakeShoes()
	fmt.Println(shirts.Wear())
	fmt.Println(shoes.Wear())
	// 测试接口隐式兼容，两个接口的方法签名是一样的，就会隐式兼容
	shoes = shirts
	fmt.Println(shoes.Wear())
}

func test_bulder() {
	// 创建想要执行的builder
	normalBuilder := housebuilder.NormalBulder{}
	iglooBuilder := housebuilder.IglooBuilder{}
	// 使用director创建不同类型
	director := builder.Director{}

	igHouse := director.MakeIGlooHouse(&iglooBuilder)
	fmt.Print(igHouse)

	normalHouse := director.MakeNormalHouse(&normalBuilder)
	fmt.Print(normalHouse)
}
