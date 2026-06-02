package main

import (
	"strings"
)

type ArtBuilder struct {
	text  string
	style string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{
		style: "normal",
	}
}

func (a *ArtBuilder) AddText(t string) *ArtBuilder {
	a.text += t
	return a
}

func (a *ArtBuilder) SetStyle(s string) *ArtBuilder {
	switch s {
	case "normal", "bold", "italic", "outline":
		a.style = s
	default:
		panic("invalid style")
	}
	return a
}

func (a *ArtBuilder) Build() string {
	lines := make([]string, 8)

	switch a.style {

	case "bold":
		var b strings.Builder
		for _, r := range a.text {
			b.WriteRune(r)
			b.WriteRune(r)
		}
		for i := 0; i < 8; i++ {
			lines[i] = b.String()
		}

	case "italic":
		for i := 0; i < 8; i++ {
			lines[i] = strings.Repeat(" ", 7-i) + a.text
		}

	case "outline":
		border := "+" + strings.Repeat("-", len(a.text)+2) + "+"
		lines[0] = border
		lines[7] = border
		for i := 1; i < 7; i++ {
			lines[i] = "| " + a.text + " |"
		}

	default: // normal
		for i := 0; i < 8; i++ {
			lines[i] = a.text
		}
	}

	return strings.Join(lines, "\n")
}
