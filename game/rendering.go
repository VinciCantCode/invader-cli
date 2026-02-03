package game

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}

func (g *Game) Draw() {
	if g.PlayerHit {
		termbox.Clear(termbox.ColorDefault, termbox.ColorRed)
	} else {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}

	// Draw borders
	for x := 0; x < g.Width; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorWhite, termbox.ColorBlack) // top
		termbox.SetCell(x, g.Height-1, '-', termbox.ColorWhite, termbox.ColorBlack) // bottom
	}
	for y := 0; y < g.Height; y++ {
		termbox.SetCell(0, y, '|', termbox.ColorWhite, termbox.ColorBlack) // left
		termbox.SetCell(g.Width-1, y, '|', termbox.ColorWhite, termbox.ColorBlack) // right
	}

	// Draw score and HP (inside borders)
	drawText(2, 1, "Score: "+fmt.Sprintf("%d", g.Score)+" High: "+fmt.Sprintf("%d", g.HighScore)+" HP: "+fmt.Sprintf("%d", g.Player.HP), termbox.ColorWhite, termbox.ColorBlack)

	// Draw player
	termbox.SetCell(g.Player.Pos.X, g.Player.Pos.Y, playerChar, termbox.ColorWhite, termbox.ColorBlack)

	// Draw aliens
	for _, alien := range g.Aliens {
		termbox.SetCell(alien.Pos.X, alien.Pos.Y, alienChar, termbox.ColorRed, termbox.ColorBlack)
	}

	// Draw bullets
	for _, bullet := range g.Bullets {
		termbox.SetCell(bullet.Pos.X, bullet.Pos.Y, bulletChar, termbox.ColorYellow, termbox.ColorBlack)
	}

	// Draw Game Over if applicable
	if g.GameOver {
		message := "Game Over! Press F to restart"
		x := (g.Width - len(message)) / 2
		y := g.Height / 2
		drawText(x, y, message, termbox.ColorRed, termbox.ColorBlack)
	} else if g.Paused {
		message := "Paused - Press P to resume"
		x := (g.Width - len(message)) / 2
		y := g.Height / 2
		drawText(x, y, message, termbox.ColorYellow, termbox.ColorBlack)
	}

	termbox.Flush()
}