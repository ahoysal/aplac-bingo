package elements

import "github.com/gdamore/tcell/v2"

type Sign struct {
	say  []string
	tb   *Textbox
	p    *Player
	x, y int
}

func NewSign(x, y int, say []string, textbox *Textbox, player *Player) *Sign {
	return &Sign{
		say: say,
		tb:  textbox,
		p:   player,
		x:   x, y: y,
	}
}

func (s *Sign) Update(events []tcell.Event) error {
	if s.InPlayerVicinity() && (s.tb.IsFinished() || s.tb.IsLowPriority()) {
		s.tb.SetLowPriorityText([]string{"press space to interact"})
		for _, ev := range events {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyRune && ev.Rune() == ' ' {
					s.tb.SetText(s.say)
					break
				}
			}
		}
	}
	return nil
}

func (s *Sign) Draw(screen tcell.Screen, cx int, cy int) {
	screen.SetContent(s.x-cx, s.y-cy, '!', nil, tcell.StyleDefault.Foreground(tcell.ColorYellow))
}

const vicinity = 3

func (s *Sign) InPlayerVicinity() bool {
	return s.p.x-vicinity <= s.x && s.x <= s.p.x+vicinity && s.p.y-vicinity <= s.y && s.y <= s.p.y+vicinity
}
