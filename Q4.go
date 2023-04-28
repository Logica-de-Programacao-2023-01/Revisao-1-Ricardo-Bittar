package main

import (
	"errors"
	"fmt"
)

func main() {
	var estadoLido string
	var impostoLido int
	var precoBase float64

	fmt.Println("Informe o preco")
	fmt.Scanln(&precoBase)
	fmt.Println("Informe o estado de origem")
	fmt.Scanln(&estadoLido)
	fmt.Println("Informe o codigo do imposto")
	fmt.Scanln(&impostoLido)
	precoFinal, err := CalculaPrecoFinal(precoBase, impostoLido, estadoLido)

	fmt.Printf("O preco final e %.2f \n erro: %f", precoFinal, err)

}
func CalculaPrecoFinal(precoBase float64, impostoLido int, estadoLido string) (float64, error) {

	const imposto1 float64 = 0.05
	const imposto2 float64 = 0.10
	const imposto3 float64 = 0.15
	const estadoSp float64 = 0.10
	const estadoRj float64 = 0.15
	const estadoMg float64 = 0.20
	const estadoEs float64 = 0.25
	const estadoOutros float64 = 0.30
	var impostoFinal float64
	var frete float64

	if impostoLido == 1 {
		impostoFinal = imposto1

	} else if impostoLido == 2 {
		impostoFinal = imposto2
	} else if impostoLido == 3 {
		impostoFinal = imposto3
	} else {
		return 0, errors.New("imposto não encontrado")
	}
	if estadoLido == "sp" {
		frete = estadoSp

	} else if estadoLido == "rj" {
		frete = estadoRj
	} else if estadoLido == "mg" {
		frete = estadoMg
	} else if estadoLido == "es" {
		frete = estadoEs
	} else {
		frete = estadoOutros
	}

	if precoBase <= 0 {
		return 0, errors.New("preco base invalido")
	}

	return precoBase + precoBase*impostoFinal + precoBase*frete, nil
}
