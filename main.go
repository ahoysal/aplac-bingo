package main

import (
	"io"
	"os"
	"time"

	"example.com/m/elements"
	"github.com/gdamore/tcell/v2"
)

const spawnx, spawny = 38, 19

func main() {
	s := initializeScreen()

	w, h := s.Size()
	camerax, cameray := spawnx-(w>>1), spawny-(h>>1)
	elms := getElements(spawnx, spawny, func(x, y int) {
		camerax = x - (w >> 1)
		cameray = y - (h >> 1)
	})

	evChan := make(chan tcell.Event)
	qChan := make(chan struct{})

	quit := func() {
		close(qChan)
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	go s.ChannelEvents(evChan, qChan)

	ticker := time.NewTicker(time.Millisecond * 100)
	events := []tcell.Event{}

	for {
		select {
		case <-ticker.C:
			s.Clear()

			for _, elm := range elms {
				elm.Update(events)
			}

			events = nil

			for _, elm := range elms {
				elm.Draw(s, camerax, cameray)
			}

			s.Show()
		case ev := <-evChan:
			events = append(events, ev)
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return
				}
			}
		}
	}
}

func ReadMap(mapFile string, tilemap map[rune]elements.Tileable) (*elements.TileMap, error) {
	req, err := os.Open(mapFile)

	if err != nil {
		return nil, err
	}

	f, err := io.ReadAll(req)
	if err != nil {
		return nil, err
	}

	return elements.NewTileMap(string(f), tilemap), nil
}
