package main

import "fmt"

func Somar(a, b int) int {
	return a + b
}

func ClassificarNota(grade float64) string {
	if grade >= 9 {
		return "A"
	} else if grade >= 7 {
		return "B"
	} else if grade >= 5 {
		return "C"
	} else {
		return "Reprovado"
	}
}

func InverterString(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(Somar(2, 3))
	fmt.Println(ClassificarNota(8))
	fmt.Println(InverterString("Original"))
}
