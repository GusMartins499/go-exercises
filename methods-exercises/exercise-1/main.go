/* Faça um programa em Go utilizando uma struct Retangulo (Largura, Altura)
   com os métodos Area() e Perimetro() para um retângulo 5 x 3, imprimindo
   o seguinte resultado:

   Área: 15.00
   Perímetro: 16.00 */

package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

func (r Retangulo) Perimetro() float64 {
	return (r.altura * 2) + (r.largura * 2)
}

func main() {
	r := Retangulo{
		largura: 5,
		altura:  3,
	}

	fmt.Printf("Área: %f\n", r.Area())
	fmt.Printf("Perímetro: %f\n", r.Perimetro())
}
