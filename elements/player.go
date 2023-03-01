package elements

import "github.com/gdamore/tcell/v2"

type Player struct {
	x, y int

	cameraCallback func(int, int)
}

func NewPlayer(x, y int, cameraCallback func(int, int)) *Player {
	return &Player{x, y, cameraCallback}
}

func (p *Player) Update(events []tcell.Event) error {
	for _, event := range events {
		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case 'w':
					p.y -= 1
				case 's':
					p.y += 1
				case 'a':
					p.x -= 1
				case 'd':
					p.x += 1
				}
			case tcell.KeyUp:
				p.y -= 1
			case tcell.KeyDown:
				p.y += 1
			case tcell.KeyLeft:
				p.x -= 1
			case tcell.KeyRight:
				p.x += 1
			}
		}
	}
	p.cameraCallback(p.x, p.y)
	return nil
}

func (p *Player) Draw(screen tcell.Screen, cx int, cy int) {
	w, h := p.x-cx, p.y-cy
	screen.SetContent(w, h-1, 'Õ', nil, tcell.StyleDefault)
	screen.SetContent(w, h, '|', nil, tcell.StyleDefault)
	screen.SetContent(w-1, h, '-', nil, tcell.StyleDefault)
	screen.SetContent(w-1, h+1, '(', nil, tcell.StyleDefault)
	screen.SetContent(w+1, h, '-', nil, tcell.StyleDefault)
	screen.SetContent(w+1, h+1, ')', nil, tcell.StyleDefault)
	// O
	//-|-
	//( )
}
