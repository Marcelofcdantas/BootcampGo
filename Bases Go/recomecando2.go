package main

import "fmt"

func main() {
	exerc1()
	exerc2()
}

func exerc1() {
	palavra := "palavra"
	fmt.Println(len(palavra))

	for _, letra := range palavra {
		fmt.Println(string(letra))
	}
}

func exerc2() {
	salario := 1000
	empregado := true
	idade := 22
	tempoAtividade := 3
	if (tempoAtividade > 1) && (idade > 22) && empregado {
		if salario > 100000 {
			fmt.Println("empréstimo sem juros")
		} else {
			fmt.Println("empréstimo com juros")
		}
	} else {
		fmt.Println("emprestimo negado")
	}
}
