package main

import (
	"strings"
)

type Animation struct {
	text   string
	frames []string
	count  int
}

func NewAnimation(text string, count int) *Animation {
	return &Animation{
		text:  text,
		count: count,
	}
}

func (a *Animation) GenerateSpinFrames() {
	a.frames = make([]string, a.count)

	base := a.text

	for i := 0; i < a.count; i++ {
		shift := i % len(base)
		rotated := base[shift:] + base[:shift]

		frame := make([]string, 10)
		for j := 0; j < 10; j++ {
			frame[j] = rotated
		}

		a.frames[i] = strings.Join(frame, "\n")
	}
}

func (a *Animation) GenerateWaveFrames() {
	a.frames = make([]string, a.count)

	for i := 0; i < a.count; i++ {
		frame := make([]string, 10)

		for j := 0; j < 10; j++ {
			indent := (i + j) % 5
			frame[j] = strings.Repeat(" ", indent) + a.text
		}

		a.frames[i] = strings.Join(frame, "\n")
	}
}

func (a *Animation) GenerateZoomFrames() {
	a.frames = make([]string, a.count)

	for i := 0; i < a.count; i++ {
		spaces := strings.Repeat(" ", i)

		var line string
		for _, r := range a.text {
			line += string(r) + spaces
		}

		frame := make([]string, 10)
		for j := 0; j < 10; j++ {
			frame[j] = line
		}

		a.frames[i] = strings.Join(frame, "\n")
	}
}

func (a *Animation) GetFrame(i int) string {
	if len(a.frames) == 0 {
		return ""
	}
	return a.frames[i%len(a.frames)]
}

func (a *Animation) Play() string {
	var out strings.Builder

	for i := 0; i < len(a.frames); i++ {
		out.WriteString("=== Frame ===\n")
		out.WriteString(a.frames[i])
		out.WriteString("\n\n")
	}

	return out.String()
}
