package figuras

import "fmt"

type Calcular interface {
	area() float64
	perimetro() float64
}

func CacularArea(calcular Calcular) {
	fmt.Println(calcular)
	fmt.Println("Area: ", calcular.area())
	fmt.Println("Perimetro:", calcular.perimetro())
}
