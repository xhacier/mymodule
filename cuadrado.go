package figuras

type Cuadrado struct {
	Ancho float64
}

func (cua *Cuadrado) area() float64 {
	return cua.Ancho * cua.Ancho
}

func (cua *Cuadrado) perimetro() float64 {
	return 2*cua.Ancho + 2*cua.Ancho
}
