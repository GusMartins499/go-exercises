/* Faça um programa em Go utilizando uma função que recebe outra função como
   parâmetro, aplicar(numeros []int, f func(int) int) []int, para aplicar
   "dobro" e "quadrado" em []int{1, 2, 3, 4}, imprimindo o seguinte
   resultado:

   Dobro: [2 4 6 8]
   Quadrado: [1 4 9 16] */

package main

import "fmt"

func quadrado(n int) int {
	return n * n
}

func dobro(n int) int {
	return n * 2
}

func aplicar(nums []int, f func(int) int) []int {
	result := make([]int, 0, len(nums))
	for _, n := range nums {
		result = append(result, f(n))
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("Dobro: %d\nQuadrado: %d\n", aplicar(nums, dobro), aplicar(nums, quadrado))
}
