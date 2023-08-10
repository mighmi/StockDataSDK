package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetTickerFromUser(t *testing.T) {
	// Define a test table with input values and expected results
	testCases := []struct {
		input    string
		expected string
	}{
		{"pbr.a", "PBR-A"},
		{"ewz", "EWZ"},
		{"GOOG", "GOOG"},
		{"bDorY", "BDORY"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// Replace os.Args with the input value for testing, then reset
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{"test_prog", tc.input}

			result := GetTickerFromUser()
			if result != tc.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
			}
		})
	}
}

// TestGetData, mok a server?

func TestWriteToFile(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input: `{
				"Meta Data": {
					"1. Information": "Daily Prices (open, high, low, close) and Volumes",
					"2. Symbol": "FAKESTOCK",
					"3. Last Refreshed": "2023-08-04",
					"4. Output Size": "Compact",
					"5. Time Zone": "US/Eastern"
				},
				"Time Series (Daily)": {
					"2023-08-04": {
						"1. open": "32.5300",
						"2. high": "32.8550",
						"3. low": "32.0425",
						"4. close": "32.0600",
						"5. volume": "30396398"
					},
					"2023-08-03": {
						"1. open": "32.8000",
						"2. high": "33.0300",
						"3. low": "32.2700",
						"4. close": "32.2900",
						"5. volume": "27908785"
					}
				}
			}`,
			expected: `{"2023-08-03":{"1. open":"32.8","2. high":"33.03","3. low":"32.27","4. close":"32.29","5. volume":"27908785"},"2023-08-04":{"1. open":"32.53","2. high":"32.855","3. low":"32.0425","4. close":"32.06","5. volume":"30396398"}}`,
		},
	}

	for _, tc := range testCases {
		t.Run("Test case", func(t *testing.T) {
			// Convert the JSON string to an io.Reader
			reader := strings.NewReader(tc.input)

			// Call the WriteToFile function with the formatted JSON data
			// Delete test file after
			filename := "test_data"
			WriteToFile(filename, ReformatJson(reader))
			defer os.Remove("data/test_data.txt")

			// Read the written data from the file
			writtenData, err := os.ReadFile("data/" + filename + ".txt")
			if err != nil {
				t.Fatalf("Error reading written data from file: %v", err)
			}

			// Compare the written data with the expected output
			actual := string(writtenData)
			expected := tc.expected

			if actual != expected {
				t.Errorf("Actual output does not match expected output for test case.\nActual:\n%s\nExpected:\n%s", actual, expected)
			}
		})
	}
}
