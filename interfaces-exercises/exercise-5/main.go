/* Faça um programa em Go utilizando type assertion (valor, ok := f.(Tipo))
   sobre uma variável do tipo Forma (do exercício 1) que guarda um
   Circulo de raio 3, imprimindo o seguinte resultado:

   É um círculo de raio 3
   Não é um retângulo */

package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	var f Forma = Circulo{Raio: 3}

	if c, ok := f.(Circulo); ok {
		fmt.Printf("É um círculo de raio %v\n", c.Raio)
	}

	if _, ok := f.(Retangulo); !ok {
		fmt.Println("Não é um retângulo")
	}
}
