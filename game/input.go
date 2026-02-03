package game

import "github.com/nsf/termbox-go"

func (g *Game) ProcessKey(key termbox.Key) {
	playerMinY := g.Height - 4
	playerMaxY := g.Height - 2
	switch key {
	case termbox.KeyArrowLeft:
		if g.Player.Pos.X > 1 {
			g.Player.Pos.X--
		}
	case termbox.KeyArrowRight:
		if g.Player.Pos.X < g.Width-2 {
			g.Player.Pos.X++
		}
	case termbox.KeyArrowUp:
		if g.Player.Pos.Y > playerMinY {
			g.Player.Pos.Y--
		}
	case termbox.KeyArrowDown:
		if g.Player.Pos.Y < playerMaxY {
			g.Player.Pos.Y++
		}
	case termbox.KeySpace:
		g.Bullets = append(g.Bullets, Bullet{Pos: Position{g.Player.Pos.X, g.Player.Pos.Y - 1}, DirY: -1})
	}
}