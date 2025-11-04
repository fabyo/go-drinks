package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/xuri/excelize/v2"
)

// Estrutura JSON da API
type Drink struct {
	IDDrink       string `json:"idDrink"`
	Name          string `json:"strDrink"`
	Category      string `json:"strCategory"`
	Alcoholic     string `json:"strAlcoholic"`
	Glass         string `json:"strGlass"`
	Instructions  string `json:"strInstructions"`
	Thumb         string `json:"strDrinkThumb"`
}

type DrinksResponse struct {
	Drinks []Drink `json:"drinks"`
}

func main() {
	// 1) Pegando os dados na API
	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?f=a"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("erro ao chamar API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API retornou status %d", resp.StatusCode)
	}

	var dr DrinksResponse
	if err := json.NewDecoder(resp.Body).Decode(&dr); err != nil {
		log.Fatalf("erro ao decodificar JSON: %v", err)
	}

	if len(dr.Drinks) == 0 {
		log.Println("Nenhum drink retornado pela API")
		return
	}

	log.Printf("Recebidos %d drinks da API\n", len(dr.Drinks))

	// 2) Cria o Excel
	f := excelize.NewFile()
	sheet := "Drinks"
	f.SetSheetName(f.GetSheetName(0), sheet)

	headers := []string{
		"ID", "Nome", "Categoria", "Alcoólico", "Copo", "Instruções", "Imagem",
	}

	for i, h := range headers {
		col := string('A' + i) // A, B, C...
		cell := fmt.Sprintf("%s1", col)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			log.Fatalf("erro ao escrever cabeçalho: %v", err)
		}
	}

	for i, d := range dr.Drinks {
		row := i + 2
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), d.IDDrink)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), d.Name)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", row), d.Category)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", row), d.Alcoholic)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", row), d.Glass)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", row), d.Instructions)
		_ = f.SetCellValue(sheet, fmt.Sprintf("G%d", row), d.Thumb)
	}

	_ = f.SetColWidth(sheet, "A", "G", 25)
	_ = f.SetColWidth(sheet, "F", "F", 60)

	filename := "drinks_a.xlsx"
	if err := f.SaveAs(filename); err != nil {
		log.Fatalf("erro ao salvar Excel: %v", err)
	}

	log.Printf("Excel gerado com sucesso: %s\n", filename)
}
