package excell

import (
	"fmt"
	"myscript/usecase"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

var (
	CelulaBaseIRT                    = "M5"
	CelulaSegurancaSocial            = "L5"
	CelulaBaseSegurancaSocial        = "J5"
	CelulaSalarioIliquido            = "I5"
	CelulaSalarioBase         string = "F5"
	CelulaData                string = "A2"
	NumeroSheet               int    = 0
	Layout                    string = "2006-01"
)

func Duplicar(argumentos usecase.Parametros) error {
	quantidade := argumentos.Quantidade
	arquivo := argumentos.Arquivo
	pasta := argumentos.Pasta

	for i := 1; i <= quantidade; i++ {
		file, err := excelize.OpenFile(arquivo)

		if err != nil {
			return fmt.Errorf("Erro ao abrir o ficherio, :%v", err)
		}

		defer file.Close()

		sheet := file.GetSheetList()[NumeroSheet]
		dataString, err := file.GetCellValue(sheet, CelulaData)
		if err != nil {
			return fmt.Errorf("Erro ao ler Campo da Data %v", err)
		}

		data, err := time.Parse(Layout, dataString)
		if err != nil {
			return fmt.Errorf("Erro ao Converter a Data Encontrada: %v", err)
		}

		if data.Year() > 2025 {
			file.SetCellValue(sheet, CelulaSalarioBase, "75000")
		}

		novaData := data.AddDate(0, i, 0)
		if err := file.SetCellValue(sheet, CelulaData, novaData.Format(Layout)); err != nil {
			return err
		}

		novoNome := fmt.Sprintf("MAPA_IRT_%02d_%d_.xlsx", int(novaData.Month()), novaData.Year())
		if err := file.SaveAs(filepath.Join(pasta, novoNome)); err != nil {
			return err
		}
	}

	return nil
}
