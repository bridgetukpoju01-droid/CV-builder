package main

import (
	"strings"
	"testing"
)

func TestStringToArt(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLines int
		contains      []string
		notContains   []string
		exactOutput   string
	}{
		{
			name:          "Single digit",
			input:         "1",
			expectedLines: 5,
			contains:      []string{"|", "  |"},
			exactOutput:   "  |  \n  |  \n  |  \n  |  \n  |  \n",
		},
		{
			name:          "Multiple digits",
			input:         "12",
			expectedLines: 5,
			contains:      []string{},
			notContains:   []string{"||"},
		},
		{
			name:          "Empty string",
			input:         "",
			expectedLines: 0,
			contains:      []string{},
		},
		{
			name:          "Invalid character",
			input:         "a",
			expectedLines: 0,
			contains:      []string{},
		},
		{
			name:          "With newline",
			input:         "1\n2",
			expectedLines: 10,
			contains:      []string{"|", "_"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToArt(tt.input)

			if tt.expectedLines == 0 && result != "" {
				t.Errorf("Expected empty string, got %q", result)
				return
			}

			if tt.expectedLines > 0 {
				lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
				if len(lines) != tt.expectedLines {
					t.Errorf("Expected %d lines, got %d", tt.expectedLines, len(lines))
				}
			}

			if tt.exactOutput != "" && result != tt.exactOutput {
				t.Errorf("Expected %q, got %q", tt.exactOutput, result)
			}

			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Result should contain %q", expected)
				}
			}

			for _, unexpected := range tt.notContains {
				if strings.Contains(result, unexpected) {
					t.Errorf("Result should not contain %q", unexpected)
				}
			}
		})
	}
}
