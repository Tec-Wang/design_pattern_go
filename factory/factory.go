package factory

import "fmt"

// 工厂设计模式

// Creator 抽象工厂作为父类，内嵌核心业务逻辑。同时作为创建具体子类的入口。但是GO中没有继承，为了发挥Creator的作用
var gunMap map[string]IGun = make(map[string]IGun)

type IFactory struct {
	// 子类应该有个统一的产品返回值，在这个里面其实就是AKGun和DesertEagle都实现的IGun
	// 利用map实现单例
	gun IGun
}

func NewFactory() *IFactory {
	factory := &IFactory{}
	return factory
}

func (gunFactory *IFactory) GetGun(name string) *IGun {
	// go中没有静态方法。所以这里使用子类工厂创建具体的子类的时候可以单独写一个方法。
	switch name {
	case "ak":
		gunFactory.gun = gunFactory.getGunFromGunMap("ak")
		fmt.Println(&gunFactory.gun)
		if gunFactory.gun == nil {
			gunFactory.gun = new(AKFactory).Crete()
			gunFactory.putGunFromGunMap("ak", gunFactory.gun)
		}
		return &gunFactory.gun

	case "desertEagle":
		deGun := gunFactory.getGunFromGunMap("desertEagle")
		if deGun == nil {
			deGun = new(DesertEagleFactory).Crete()
			gunFactory.putGunFromGunMap("desertEagle", deGun)
		}
		return &deGun
	default:
		return nil
	}
}

func (gunFactory *IFactory) getGunFromGunMap(name string) IGun {
	return gunMap[name]
}

func (gunFactory *IFactory) putGunFromGunMap(gunName string, gun IGun) {
	gunMap[gunName] = gun
}

type AKFactory struct {
}

func (aKFactory *AKFactory) Crete() *AKGun {
	return &AKGun{
		Gun{
			Name:  "ak",
			Price: 1888,
		},
	}

}

type DesertEagleFactory struct{}

func (desertEagleFactory *DesertEagleFactory) Crete() *DesertEagle {
	return &DesertEagle{
		Gun{
			Name:  "desertEagle",
			Price: 18888,
		},
	}
}

type IGun interface {
	Shoot() string
}

type Gun struct {
	Name  string
	Price int
}

type AKGun struct {
	Gun
}

type DesertEagle struct {
	Gun
}

func (ak *AKGun) Shoot() string {
	return "AK biubiubiu"
}

func (de *DesertEagle) Shoot() string {
	return "沙漠之鹰"
}
