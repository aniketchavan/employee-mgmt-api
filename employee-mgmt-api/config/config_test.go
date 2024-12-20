package config

import (
	"strconv"
	"testing"
)

func TestGetConfigProperty(t *testing.T) {
	InitializeConfigurations() // Ensure configurations are loaded

	// Existing property
	value := len(ConfigSet.Properties)
	if value == 0 {
		t.Errorf("PropertySet cannot be null")
	}

	// Non-existing property
	fieldValue := ConfigSet.Properties["MySqlUsername"]
	if fieldValue == "" {
		t.Errorf("expected non-nil value, got %v", value)
	}

	fieldValue = ConfigSet.Properties["MySqlPort"]
	if fieldValue != "" {
		_, err := strconv.Atoi(fieldValue)
		if err != nil {
			t.Errorf("expected int value, got string value- %v", value)
		}
	}
}
