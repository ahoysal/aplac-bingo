package renderer

import (
	"errors"
	"image"

	"github.com/gdamore/tcell/v2"
)

func DrawBoxAround(rect image.Rectangle, style tcell.Style, screen tcell.Screen) {
	for x := rect.Max.X + 1; x >= rect.Min.X-1; x-- {
		screen.SetContent(x, rect.Min.Y-1, '-', nil, style)
		screen.SetContent(x, rect.Max.Y+1, '-', nil, style)
	}

	for y := rect.Max.Y; y >= rect.Min.Y; y-- {
		screen.SetContent(rect.Min.X-1, y, '|', nil, style)
		screen.SetContent(rect.Max.X+1, y, '|', nil, style)
	}
}

func FillRect(rect image.Rectangle, primary rune, comb []rune, style tcell.Style, screen tcell.Screen) {
	for x := rect.Max.X; x >= rect.Min.X; x-- {
		for y := rect.Max.Y; y >= rect.Min.Y; y-- {
			screen.SetContent(x, y, primary, comb, style)
		}
	}
}

func DrawText(text []rune, rect image.Rectangle, style tcell.Style, screen tcell.Screen) error {
	offset := 0
	dx, dy := rect.Dx(), rect.Dy()
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '\n':
			offset += dx - (offset + i) - 1
			continue
		}
		x, y := (offset+i)%dx, (offset+i)/dx
		if y > dy {
			return errors.New("text overflows")
		}
		screen.SetContent(x+rect.Min.X, y+rect.Min.Y, text[i], nil, style)
	}
	return nil
}
