package abstractfactory

type AddidasShoe struct {
}

func (*AddidasShoe) Wear() string {
	return "AddidasShoe"
}
