package main

import (
	"fmt"
	"log"
	"myscript/excell"
	"myscript/usecase"
)

func main() {
	argumentos, err := usecase.ProcessarArgumentos()
	if err != nil {
		log.Fatal(err)
	}

	err = excell.Duplicar(argumentos)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Sucesso")
}
