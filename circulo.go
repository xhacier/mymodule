package figuras

type Circulo struct {
	Radio float64
}

func (cur *Circulo) area() float64 {
	return 3.1416 * (cur.Radio * cur.Radio)
}

func (cur *Circulo) perimetro() float64 {
	return 2 * 3.1416 * cur.Radio
}
