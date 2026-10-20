/* Faça um programa em Go utilizando uma struct Pessoa (Nome, Idade) com o
   método Apresentar() string e o método FazerAniversario() com receptor
   ponteiro, imprimindo o seguinte resultado:

   Olá, eu sou Ana e tenho 25 anos
   Olá, eu sou Ana e tenho 26 anos */

package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, eu sou %s e tenho %d anos", p.Nome, p.Idade)
}

func (p *Pessoa) FazerAniversario() {
	p.Idade++
}

func main() {
	p := Pessoa{
		Nome:  "Ana",
		Idade: 25,
	}

	fmt.Println(p.Apresentar())
	p.FazerAniversario()
	fmt.Println(p.Apresentar())
}
