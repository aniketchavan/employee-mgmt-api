package utils

import (
	"testing"
)

func TestParseExcelFile(t *testing.T) {
	// Valid file
	filePath := "../test/valid_employees.xlsx"
	_, err := ParseExcelFile(filePath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Invalid file format
	filePath = "../test/invalid_file.txt"
	_, err = ParseExcelFile(filePath)
	if err == nil {
		t.Fatal("expected error, got none")
	}

	// Missing file
	filePath = "../test/missing_file.xlsx"
	_, err = ParseExcelFile(filePath)
	if err == nil {
		t.Fatal("expected error, got none")
	}

	// Empty file
	filePath = "../test/empty_file.xlsx"
	_, err = ParseExcelFile(filePath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
