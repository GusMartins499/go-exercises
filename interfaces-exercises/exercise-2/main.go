/* Faça um programa em Go utilizando uma interface Animal com os métodos
   Nome() string e Som() string, implementada por Cachorro, Gato e Vaca.
   Percorra um []Animal, imprimindo o seguinte resultado:

   Cachorro: Au au
   Gato: Miau
   Vaca: Muuu */

package main

import "fmt"

type Animal interface {
	nome() string
	som() string
}

type Cachorro struct{}
type Gato struct{}
type Vaca struct{}

func (a Cachorro) nome() string {
	return "Cachorro"
}

func (a Cachorro) som() string {
	return "Au au"
}

func (a Gato) nome() string {
	return "Gato"
}

func (a Gato) som() string {
	return "Miau"
}

func (a Vaca) nome() string {
	return "Vaca"
}

func (a Vaca) som() string {
	return "Muuu"
}

func main() {
	animals := []Animal{
		Cachorro{},
		Gato{},
		Vaca{},
	}

	for _, a := range animals {
		fmt.Printf("%s: %s\n", a.nome(), a.som())
	}
}
