package main

import (
	"fmt"
	"log"
	"myscript/excell"
	"myscript/usecase"
	"strconv"
)

func main() {
	argumentos, err := usecase.ProcessarArgumentos()
	if err != nil {
		log.Fatal(err)
	}

	arquivo := argumentos["arquivo"]
	pasta := argumentos["pasta"]
	quantidade := argumentos["quantidade"]
	qtd, err := strconv.Atoi(quantidade)
	if err != nil {
		qtd = 1
	}

	err = excell.Duplicar(arquivo, pasta, qtd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sucesso")
}
