package csv

import (
	"encoding/csv"
	"os"
)

func Decode(pathToFile string, loc string) [][]string {
	file, err := os.Open(pathToFile)
	if err != nil {
		panic(err)
	}
	
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var matrix [][]string
	for _, row := range data {
		if isValid(row, loc) {
			var a []string
			for i, col := range row {
				a = addField(a, i, col)
			}
			matrix = append(matrix, a)
		}
	}
	defer file.Close()
	return matrix
}

func addField(a []string, index int, field string) []string {
	if index == 0 {
		a = append(a, field)
	}
	if index == 6 {
		a = append(a, field[:6])
	}
	if  index == 7 {
		a = append(a, shortPizzaName(field))
	}
	return a
}

func shortPizzaName(s string) string {
	if s == "Vegetarische Pizza" {
		return "vegi"
	}
	if s == "Pizza mit Fleisch" {
		return "fleisch"
	}
	return "no"
}

func isValid(row []string, loc string) bool {
	return row[1] == "Ja" && row[2] == loc
}
