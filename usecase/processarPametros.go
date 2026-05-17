package usecase

import (
	"flag"
)

type Parametros struct {
	Arquivo    string
	Pasta      string
	Quantidade int
}

func ProcessarArgumentos() (Parametros, error) {
	var argumentos Parametros

	arquivo := flag.String("arquivo", "", "Arquivo Excell")
	pasta := flag.String("saida", ".", "Pastaa de Saida")
	quantidade := flag.Int("qtd", 2, "Quantidade de Arquivos")

	flag.Parse()

	if err := validarArquivo(*arquivo); err != nil {
		return argumentos, err
	}

	argumentos.Arquivo = *arquivo
	argumentos.Pasta = *pasta
	argumentos.Quantidade = *quantidade

	return argumentos, nil
}
