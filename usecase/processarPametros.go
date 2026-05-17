package usecase

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ProcessarArgumentos() (map[string]string, error) {
	input := os.Args
	argumentos := make(map[string]string)
	var pasta string
	var quantidade string

	if len(input) < 2 {
		return argumentos, errors.New("Faltando argumentos")
	}
	arquivo := input[1]

	if len(input) > 2 {
		pasta = input[2]
	}

	if len(input) > 3 {
		quantidade = input[3]
	}

	if !strings.Contains(arquivo, "xlsx") {
		return argumentos, errors.New("Arquivo invalido,Tente um arquivo Excell (xlsx)")
	}

	qtd, err := strconv.Atoi(quantidade)
	if err != nil || qtd < 1 {
		quantidade = "1"
	}

	pasta = filepath.Clean(pasta)
	if filepath.IsLocal(pasta) {
		return argumentos, errors.New("Caminho da Pasta inválido")
	}

	argumentos["arquivo"] = arquivo
	argumentos["pasta"] = pasta
	argumentos["quantidade"] = quantidade

	return argumentos, nil
}
