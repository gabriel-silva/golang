package main

import (
	"fmt"
	"math"
)

type IGeometrica interface {
	area() float64
}

type Quadrado struct {
	lado float64
}

func (q Quadrado) area() float64 {
	return q.lado * q.lado
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func main() {
	var g IGeometrica = Quadrado{3}
	fmt.Printf("A area do quadrado é: %.2f\n", g.area())

	g = Circulo{5}
	fmt.Printf("A area do circulo é: %.2f\n", g.area())

}
