package service

import (
	"errors"
	"go-restful-app/csvdata"
	"go-restful-app/models"
	"strconv"
)

func GetAllRecords() ([]models.Record, error) {
	return csvdata.ReadCSV()
}

func GetRecordByID(id string) (models.Record, error) {
	records, err := csvdata.ReadCSV()
	if err != nil {
		return models.Record{}, err
	}

	intID, _ := strconv.Atoi(id)
	for _, record := range records {
		if record.FixletID == intID {
			return record, nil
		}
	}
	return models.Record{}, errors.New("fixlet not found")
}

func CreateRecord(record models.Record) (models.Record, error) {
	records, err := csvdata.ReadCSV()
	if err != nil {
		return models.Record{}, err
	}

	record.FixletID = getNextID(records)
	records = append(records, record)
	if err := csvdata.WriteCSV(records); err != nil {
		return models.Record{}, err
	}
	return record, nil
}

func UpdateRecord(id string, updatedRecord models.Record) (models.Record, error) {
	records, err := csvdata.ReadCSV()
	if err != nil {
		return models.Record{}, err
	}

	intID, _ := strconv.Atoi(id)
	for i, record := range records {
		if record.FixletID == intID {
			records[i] = updatedRecord
			records[i].FixletID = intID
			if err := csvdata.WriteCSV(records); err != nil {
				return models.Record{}, err
			}
			return records[i], nil
		}
	}
	return models.Record{}, errors.New("fixlet not found")
}

func DeleteRecord(id string) error {
	records, err := csvdata.ReadCSV()
	if err != nil {
		return err
	}

	intID, _ := strconv.Atoi(id)
	for i, record := range records {
		if record.FixletID == intID {
			records = append(records[:i], records[i+1:]...)
			return csvdata.WriteCSV(records)
		}
	}
	return errors.New("fixlet not found")
}

func getNextID(records []models.Record) int {
	maxID := 0
	for _, record := range records {
		if record.FixletID > maxID {
			maxID = record.FixletID
		}
	}
	return maxID + 1
}
