// handlers/handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmployeeHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
	}{
		{"GetEmployees", "GET", "/employees", http.StatusOK},
		{"CreateEmployee", "POST", "/employees", http.StatusCreated},
		{"UpdateEmployee", "PUT", "/employees/1", http.StatusOK},
		{"DeleteEmployee", "DELETE", "/employees/1", http.StatusNoContent},
		{"InvalidMethod", "PATCH", "/employees", http.StatusMethodNotAllowed},
		{"NotFound", "GET", "/nonexistent", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(EmployeeHandler) // Replace with your actual handler function

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

		})
	}
}
