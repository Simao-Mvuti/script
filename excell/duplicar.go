package excell

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

var (
	CelulaData  string = "A2"
	NumeroSheet int    = 0
	Layout      string = "2006-01"
)

func Duplicar(arquivo string, pasta string, quantidade int) error {
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

		novaData := data.AddDate(0, i, 0)

		if err := file.SetCellValue(sheet, CelulaData, novaData.Format(Layout)); err != nil {
			return err
		}

		novoNome := fmt.Sprintf("MAPA_IRT_%02d_%d_.xlsx", int(novaData.Month()), novaData.Year())
		if err := file.SaveAs(novoNome); err != nil {
			return err
		}
	}

	return nil
}
