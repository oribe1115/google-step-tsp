package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// CSVRead csvから座標データを読み取って返す
func CSVRead(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	coordData := make([][]float64, 0)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		x, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		coordData = append(coordData, []float64{x, y})
	}

	return coordData, nil
}

// CSVWrite csvに経路を書き込む
func CSVWrite(filename string, tour *Tour) error {
	record := make([][]string, 0)
	record = append(record, []string{"index"})
	for _, id := range *tour {
		record = append(record, []string{fmt.Sprintf("%d", id)})
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(record)
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}
