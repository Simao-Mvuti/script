package usecase

import (
	"fmt"
	"os"
	"path/filepath"
)

func validarArquivo(path string) error {
	if path == "" {
		return fmt.Errorf("arquivo vazio")
	}

	ext := filepath.Ext(path)

	if ext != ".xlsx" {
		return fmt.Errorf("arquivo precisa ser .xlsx")
	}

	_, err := os.Stat(path)

	if err != nil {
		return fmt.Errorf("arquivo não encontrado")
	}

	return nil
}
