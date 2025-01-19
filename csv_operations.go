package csvdata

import (
	"encoding/csv"
	"go-restful-app/models"
	"os"
	"strconv"
)

const csvFileName = "fixlets.csv" // Path to fixlets.csv

// ReadCSV reads all records from the CSV file
func ReadCSV() ([]models.Record, error) {
	file, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records []models.Record
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Skip the header row
	for _, row := range rows[1:] {
		siteID, _ := strconv.Atoi(row[0])
		fixletID, _ := strconv.Atoi(row[1])
		relevantCount, _ := strconv.Atoi(row[4])
		records = append(records, models.Record{
			SiteID:                siteID,
			FixletID:              fixletID,
			Name:                  row[2],
			Criticality:           row[3],
			RelevantComputerCount: relevantCount,
		})
	}

	return records, nil
}

// WriteCSV writes records to the CSV file
func WriteCSV(records []models.Record) error {
	file, err := os.Create(csvFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	writer.Write([]string{"SiteID", "FixletID", "Name", "Criticality", "RelevantComputerCount"})

	for _, record := range records {
		row := []string{
			strconv.Itoa(record.SiteID),
			strconv.Itoa(record.FixletID),
			record.Name,
			record.Criticality,
			strconv.Itoa(record.RelevantComputerCount),
		}
		writer.Write(row)
	}

	return nil
}
