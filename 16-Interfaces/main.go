package main

import (
	"fmt"
	"math"
)

//As interfaces possuem SOMENTE assinaturas de métodos, nada mais!
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (ret retangulo) area() float64 {
	return ret.altura * ret.largura
}

type ciruclo struct {
	raio float64
}

func (circ ciruclo) area() float64 {
	return math.Pi * math.Pow(circ.raio, 2)
}

func main() {
	ret := retangulo{10, 15}
	escreverArea(ret)

	circ := ciruclo{10}
	escreverArea(circ)
}
