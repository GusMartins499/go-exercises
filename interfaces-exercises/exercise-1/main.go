/* Faça um programa em Go utilizando uma interface Forma com o método
   Area() float64, implementada por Retangulo (4 x 5) e Circulo (raio 2),
   e uma função imprimirArea(f Forma), imprimindo o seguinte resultado:

   Área: 20.00
   Área: 12.57 */

package main

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func main() {
	r := Retangulo{10, 15}
	escreverArea(r)

	c := Circulo{10}
	escreverArea(c)

}
