/* Faça um programa em Go utilizando um type switch (switch v := x.(type))
   para percorrer []any{42, "Go", 3.14, true}, imprimindo o seguinte
   resultado:

   42 é int
   Go é string
   3.14 é float64
   true é bool */

package main

import "fmt"

func main() {
	values := []any{42, "Go", 3.14, true}
	for _, v := range values {
		switch v := v.(type) {
		case int:
			fmt.Println(v, "é int")
		case string:
			fmt.Println(v, "é string")
		case float64:
			fmt.Println(v, "é float64")
		case bool:
			println(v, "é bool")
		default:
			println(v, "é de um tipo desconhecido")
		}
	}
}
