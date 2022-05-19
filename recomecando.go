package main

import "fmt"

func main() {
	exec1()
	exec2()
	exec3()
	exec4()
}

func exec1() {
	nome := "Marcelo"
	var idade int = 37
	fmt.Println(nome, idade)
}

func exec2() {
	var (
		temperatura         int    = 37
		umidade             string = "ok"
		pressao_atmosferica int8   = 10
	)
	fmt.Println(temperatura, umidade, pressao_atmosferica)
}

func exec3() {
	fmt.Println("var nome1 string")
	fmt.Println("var sobrenome string")
	fmt.Println("var idade int")
	fmt.Println("sobrenome1 := 6")
	fmt.Println("var licenca_para_dirigir bool = true")
	fmt.Println("var estatura_da_pessoa int")
	fmt.Println("quantidadeDeFilhos := 2")
}

func exec4() {
	var sobrenome = "Silva"
	var idade int = 25
	boolean := "false"
	var salario float32 = 4585.90
	var nome string = "Felippe"
	fmt.Println(sobrenome, idade, boolean, salario, nome)
}
