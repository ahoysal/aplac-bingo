package elements

import "github.com/gdamore/tcell/v2"

type Element interface {
	Update([]tcell.Event) error
	Draw(tcell.Screen, int, int)
}
