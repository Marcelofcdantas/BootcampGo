package main

import "fmt"

func main() {
	exerc1()
	exerc2()
	resp1, resp2 := exerc3(3)
	fmt.Printf("%v de %s \n", resp2, resp1)

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

func exerc3(numero int) (string, int) {
	var mesDoAno string
	switch {
	case numero == 1:
		mesDoAno = "Janeiro"
	case numero == 2:
		mesDoAno = "Fevereiro"
	case numero == 3:
		mesDoAno = "Março"
	case numero == 4:
		mesDoAno = "Abril"
	case numero == 5:
		mesDoAno = "Maio"
	case numero == 6:
		mesDoAno = "Junho"
	case numero == 7:
		mesDoAno = "Julho"
	case numero == 8:
		mesDoAno = "Agosto"
	case numero == 9:
		mesDoAno = "Setembro"
	case numero == 10:
		mesDoAno = "Outubro"
	case numero == 11:
		mesDoAno = "Novembro"
	case numero == 12:
		mesDoAno = "Dezembro"
	default:
		mesDoAno = "Mês inválido"
	}
	return mesDoAno, numero
}

func exerc4() {
	var employees = map[string]int {
		"Benjamim": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
		
	}
