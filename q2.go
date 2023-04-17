package main

import "fmt"

func numeromaior(lista []float64) float64 {
	max := lista[0]
	for _, num := range lista {
		if num > max {
			max = num
		}

	}
	return max
}
func numeromenor(lista []float64) float64 {

	min := lista[0]

	for _, num := range lista {
		if num < min {
			min = num
		}
	}
	return min
}

func media(lista []float64) float64 {
	var media float64
	var soma float64
	soma = 0

	for _, num := range lista {
		soma += num

	}
	media = float64(soma) / float64(len(lista))
	return media
}

func main() {

	var lista []float64
	var numero float64
	numero = -1
	println("Digite os números inteiros para calcular o maior, o menor e a média. Digite 0 para ir ao resultado!")
	for numero != 0 {

		fmt.Println("Digite um número")
		fmt.Scanln(&numero)
		if numero != 0 {
			lista = append(lista, numero)
		}

	}
	fmt.Printf("\n\nO maior valor é: %.0f \n", numeromaior(lista))
	fmt.Printf("O menor numero é: %.0f \n", numeromenor(lista))
	fmt.Printf("A media da lista é: %.2f", media(lista))

}
