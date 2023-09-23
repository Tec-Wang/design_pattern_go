package abstractfactory

type NikeShirt struct {
}

func (*NikeShirt) Wear() string {
	return "nikeShirt"
}
