package elements

import (
	"errors"
	"image"

	"example.com/m/renderer"
	"github.com/gdamore/tcell/v2"
)

type Textbox struct {
	text                  [][]rune
	slideIndex, charIndex int
	lowPriority           bool
}

func NewTextbox() *Textbox {
	return &Textbox{
		text:       nil,
		slideIndex: -1, charIndex: -1,
	}
}

func (t *Textbox) SetText(slides []string) error {
	t.text = make([][]rune, len(slides))
	for i, v := range slides {
		t.text[i] = []rune(v)
	}

	t.slideIndex = 0
	t.charIndex = 0
	t.lowPriority = false
	return nil
}

func (t *Textbox) SetLowPriorityText(slides []string) error {
	if t.lowPriority {
		return errors.New("conflicting text")
	}
	t.SetText(slides)
	t.lowPriority = true
	return nil
}

func (t *Textbox) IsFinished() bool {
	return t.slideIndex == -1
}

func (t *Textbox) IsLowPriority() bool {
	return t.lowPriority
}

func (t *Textbox) Update(events []tcell.Event) error {
	if t.slideIndex == -1 {
		return nil
	}

	clicked := false
	for _, ev := range events {
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune && ev.Rune() == ' ' {
				clicked = true
				break
			}
		}
	}

	slidelen := len(t.text[t.slideIndex])

	if clicked {
		if t.charIndex < slidelen {
			t.charIndex = slidelen
		} else {
			t.charIndex = 0
			t.slideIndex++
			if t.slideIndex == len(t.text) {
				t.slideIndex = -1
				t.lowPriority = false
			}
		}
	} else if t.charIndex < slidelen {
		t.charIndex++
	}

	return nil
}

func (t *Textbox) Draw(screen tcell.Screen, cx int, cy int) {
	w, h := screen.Size()
	trect := image.Rect(3, 2, w-4, h/3-1)
	renderer.DrawBoxAround(trect, tcell.StyleDefault, screen)
	renderer.FillRect(trect, ' ', nil, tcell.StyleDefault, screen)
	if t.slideIndex != -1 {
		renderer.DrawText(t.text[t.slideIndex][:t.charIndex], trect, tcell.StyleDefault, screen)
	}
}
