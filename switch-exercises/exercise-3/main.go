/* Faça um programa em Go utilizando switch com vários valores no mesmo
   case (case 'a', 'e', 'i', 'o', 'u') para classificar as letras
   'a', 'b', 'e' e 'z', imprimindo o seguinte resultado:

   a é vogal
   b é consoante
   e é vogal
   z é consoante */

package main

import "fmt"

func main() {
	classifyLetter := func(letter rune) {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			fmt.Println(string(letter), "é vogal")
		default:
			fmt.Println(string(letter), "é consoante")
		}
	}

	classifyLetter('a')
	classifyLetter('b')
	classifyLetter('e')
	classifyLetter('z')
}
