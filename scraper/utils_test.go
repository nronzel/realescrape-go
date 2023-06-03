package scraper

import (
	"testing"
)

// TestSplitUnits checks if the function correctly separates units and numbers from input
func TestSplitUnits(t *testing.T) {
	testCases := []struct {
		input  string
		number string
		unit   string
	}{
		{"200 acre", "200", "acre"},
		{"20 km", "20", "km"},
		{"no number or units", "", ""},
		{"", "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			number, unit, _ := splitUnits(tc.input)
			if number != tc.number || unit != tc.unit {
				t.Errorf("Expected %s and %s, got %s and %s", tc.number, tc.unit, number, unit)
			}
		})
	}
}

// TestParseAddress checks if the function correctly separates components of an address
func TestParseAddress(t *testing.T) {
	t.Run("should correctly scraper a full address", func(t *testing.T) {
		street, city, state, zip := parseAddress("123 Fake St, Springfield, IL 12345")
		if street != "123 Fake St" || city != "Springfield" || state != "IL" || zip != "12345" {
			t.Errorf("Expected parsed address components, got incorrect values")
		}
	})

	t.Run("should return empty strings for incomplete address", func(t *testing.T) {
		street, city, state, zip := parseAddress("Incomplete address")
		if street != "" || city != "" || state != "" || zip != "" {
			t.Errorf("Expected empty strings, got non-empty values")
		}
	})
}

// TestConvertToSqft checks if the function correctly converts acres to square feet
func TestConvertToSqft(t *testing.T) {
	testCases := []struct {
		acre   string
		sqft   string
		hasErr bool
	}{
		{"2", "87120", false},
		{"0", "0", false},
		{"-1", "0", true},
		{"invalid", "0", true},
	}

	for _, tc := range testCases {
		t.Run(tc.acre, func(t *testing.T) {
			sqft, err := convertToSqft(tc.acre)

			if (err != nil) != tc.hasErr {
				t.Errorf("convertToSqft() error = %v, wantErr %v", err, tc.hasErr)
				return
			}

			if sqft != tc.sqft {
				t.Errorf("Expected %s, got %s", tc.sqft, sqft)
			}
		})
	}
}

// TestHtyRatios checks if the function correctly calculates ratios
func TestHtyRatios(t *testing.T) {
	t.Run("should correctly calculate ratios", func(t *testing.T) {
		ratio1, ratio2 := htyRatios("2000", "10000")
		if ratio1 != "5.00" || ratio2 != "0.20" {
			t.Errorf("Expected '5.00' and '0.20', got '%s' and '%s'", ratio1, ratio2)
		}
	})
	t.Run("should return empty strings for empty input", func(t *testing.T) {
		ratio1, ratio2 := htyRatios("", "")
		if ratio1 != "" || ratio2 != "" {
			t.Errorf("Expected '' and '', got '%s' and '%s'", ratio1, ratio2)
		}
	})
}
