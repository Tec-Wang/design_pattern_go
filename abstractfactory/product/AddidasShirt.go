package abstractfactory

type AddidasShirt struct {
}

func (*AddidasShirt) Wear() string {
	return "AddidasShirt"
}
