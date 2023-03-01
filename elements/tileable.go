package elements

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Tileable interface {
	Tile(x, y int) (rune, []rune, tcell.Style)
}

type SimpleTile struct {
	toTile [][]rune
	w, h   int
	style  tcell.Style
}

func NewSimpleTile(toTile string, style tcell.Style) *SimpleTile {
	rows := strings.Split(toTile, "\n")
	arr := make([][]rune, len(rows))
	for r := range rows {
		arr[r] = []rune(rows[r])
	}

	return &SimpleTile{
		toTile: arr,
		w:      len(arr[0]), h: len(rows),
		style: style,
	}
}

func (st *SimpleTile) Tile(x, y int) (rune, []rune, tcell.Style) {
	return st.toTile[y%st.h][x%st.w], nil, st.style
}
