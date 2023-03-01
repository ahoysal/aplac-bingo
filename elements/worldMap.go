package elements

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type TileMap struct {
	grid  [][]rune
	w, h  int
	tiles map[rune]Tileable
}

func NewTileMap(pureText string, tiles map[rune]Tileable) *TileMap {
	rows := strings.Split(pureText, "\n")
	w, h := len(rows[0]), len(rows)
	grid := make([][]rune, h)
	for r := range grid {
		grid[r] = []rune(rows[r])
	}
	return &TileMap{
		grid: grid,
		w:    w, h: h,
		tiles: tiles}
}

func (tm *TileMap) Update([]tcell.Event) error {
	return nil
}

func (tm *TileMap) Draw(screen tcell.Screen, cx, cy int) {
	w, h := screen.Size()
	for sy := 0; sy < h; sy++ {
		for sx := 0; sx < w; sx++ {
			wx, wy := sx+cx, sy+cy
			if t, err := tm.Get(wx, wy); err == nil {
				if tileable, ok := tm.tiles[t]; ok {
					main, combc, style := tileable.Tile(wx, wy)
					screen.SetContent(sx, sy, main, combc, style)
				} else {
					screen.SetContent(sx, sy, rune(t), nil, tcell.StyleDefault)
				}
			}
		}
	}
}

func (tm *TileMap) InBounds(x, y int) bool {
	return 0 <= x && x < tm.w && 0 <= y && y < tm.h
}

func (tm *TileMap) Get(x, y int) (rune, error) {
	if tm.InBounds(x, y) {
		return tm.grid[y][x], nil
	}
	return '0', fmt.Errorf("%d, %d out of bounds", x, y)
}
