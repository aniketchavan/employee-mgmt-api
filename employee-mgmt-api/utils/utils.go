package utils

import (
	"fmt"
	"strconv"

	"example.com/employee-mgmt/config"
	datamodel "example.com/employee-mgmt/models"
	"github.com/xuri/excelize/v2"
)

func ParseExcelFile(filePath string) ([]datamodel.Employee, error) {

	// To skip header insertion
	var startRecordAt int
	if startRecord, isPresent := config.ConfigSet.Properties["StartRecordFrom"]; isPresent {
		var err error
		startRecordAt, err = strconv.Atoi(startRecord)
		if err != nil {
			startRecordAt = 0
		}
	}

	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open Excel file: %v", err)
	}

	// Read data from the first sheet
	// TODO: Make the sheet name configurable
	rows, err := f.GetRows("uk-500")
	if err != nil {
		return nil, fmt.Errorf("unable to read rows: %v", err)
	}

	// Parse the rows and create a slice of employees
	var employees []datamodel.Employee

	for idx, row := range rows {
		// Skip headers from request file
		if idx < startRecordAt {
			continue
		}
		if len(row) < 10 {
			continue
		}
		employee := datamodel.Employee{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			Country:     row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
