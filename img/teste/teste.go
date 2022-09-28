// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main() {

	si := []int{10, 20, 30, 40, 50}
	fmt.Println(funcao1(si...))

	sint := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(funcao2(sint))

}

func funcao1(x ...int) int {
	total := 0
	for _, valor := range x {
		total += valor
	}
	return total
}

func funcao2(y []int) int {
	total := 0
	for _, value := range y {
		total += value
	}
	return total
}
