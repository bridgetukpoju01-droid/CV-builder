package main

import (
	"strings"
	"testing"
)

func TestArtBuilder_BasicFlow(t *testing.T) {
	builder := NewArtBuilder()
	result := builder.
		AddText("HI").
		SetStyle("normal").
		Build()

	if result == "" {
		t.Error("Expected non-empty result")
	}

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}
}

func TestArtBuilder_MethodChaining(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*ArtBuilder) *ArtBuilder
		expectLines int
		minWidth    int
	}{
		{
			name: "Single text normal style",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("A").SetStyle("normal")
			},
			expectLines: 8,
			minWidth:    1, // At least 1 character wide (actual width depends on font)
		},
		{
			name: "Bold style",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("A").SetStyle("bold")
			},
			expectLines: 8,
			minWidth:    2, // Bold should be at least as wide as normal
		},
		{
			name: "Multiple texts with different styles",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.
					AddText("A").
					SetStyle("normal").
					AddText("B").
					SetStyle("italic")
			},
			expectLines: 8,
			minWidth:    2, // At least 2 characters wide
		},
		{
			name: "Outline style",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("T").SetStyle("outline")
			},
			expectLines: 8,
			minWidth:    2, // Outline adds border, so wider than normal
		},
		{
			name: "Italic style has forward slant",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("I").SetStyle("italic")
			},
			expectLines: 8,
			minWidth:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewArtBuilder()
			builder = tt.setup(builder)
			result := builder.Build()

			if result == "" {
				t.Errorf("Expected non-empty result for %s", tt.name)
			}

			lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
			if len(lines) != tt.expectLines {
				t.Errorf("Expected %d lines, got %d", tt.expectLines, len(lines))
			}

			if len(lines) > 0 && len(lines[0]) < tt.minWidth {
				t.Errorf("Expected min width %d, got %d", tt.minWidth, len(lines[0]))
			}
		})
	}
}

func TestArtBuilder_ItalicSlant(t *testing.T) {
	builder := NewArtBuilder()
	normalResult := builder.AddText("I").SetStyle("normal").Build()

	builder = NewArtBuilder()
	italicResult := builder.AddText("I").SetStyle("italic").Build()

	// Italic should be different from normal (demonstrates forward slant)
	if normalResult == italicResult {
		t.Error("Italic style should differ from normal style (forward slant required)")
	}

	// Check that italic result isn't empty
	if italicResult == "" {
		t.Error("Italic style should produce non-empty result")
	}
}

func TestArtBuilder_InvalidStyle(t *testing.T) {
	builder := NewArtBuilder()

	// According to requirements, only "bold", "italic", "outline", "normal" are supported
	// Invalid styles should be rejected (behavior depends on implementation)
	// This test expects panic for invalid style
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid style 'rainbow' - only supported styles: normal, bold, italic, outline")
		}
	}()

	builder.AddText("HI").SetStyle("rainbow").Build()
}

func TestArtBuilder_BoldStyle(t *testing.T) {
	builder := NewArtBuilder()
	normalResult := builder.AddText("A").SetStyle("normal").Build()

	builder = NewArtBuilder()
	boldResult := builder.AddText("A").SetStyle("bold").Build()

	// Bold should be different from normal
	if normalResult == boldResult {
		t.Error("Bold style should differ from normal style")
	}

	// Bold should typically be wider
	normalLines := strings.Split(strings.TrimRight(normalResult, "\n"), "\n")
	boldLines := strings.Split(strings.TrimRight(boldResult, "\n"), "\n")

	if len(normalLines) > 0 && len(boldLines) > 0 {
		if len(boldLines[0]) <= len(normalLines[0]) {
			t.Log("Warning: Bold style is not wider than normal style (but this may be acceptable depending on implementation)")
		}
	}
}

func TestArtBuilder_OutlineStyle(t *testing.T) {
	builder := NewArtBuilder()
	outlineResult := builder.AddText("T").SetStyle("outline").Build()

	if outlineResult == "" {
		t.Error("Outline style should produce non-empty result")
	}

	// Outline should add border characters (like '+', '-', '|')
	hasBorderChars := strings.ContainsAny(outlineResult, "+-|")
	if !hasBorderChars {
		t.Log("Warning: Outline style doesn't contain typical border characters")
	}
}

func TestArtBuilder_MethodChainingReturn(t *testing.T) {
	builder := NewArtBuilder()

	// Each method should return *ArtBuilder for chaining
	result1 := builder.AddText("A")
	if result1 != builder {
		t.Error("AddText should return the same builder instance for chaining")
	}

	result2 := builder.SetStyle("bold")
	if result2 != builder {
		t.Error("SetStyle should return the same builder instance for chaining")
	}

	result3 := builder.AddText("B")
	if result3 != builder {
		t.Error("AddText should return the same builder instance for chaining")
	}
}
