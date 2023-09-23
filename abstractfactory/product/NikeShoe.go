package abstractfactory

type NikeShoe struct {
}

func (*NikeShoe) Wear() string {
	return "nikeShoe"
}
