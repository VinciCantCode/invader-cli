package main

import (
	"bufio"
	"fmt"
	"github.com/nsf/termbox-go"
	"invader-cli/game"
	"os"
	"strconv"
	"time"
)

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, fg, bg)
	}
}

func loadHighScore() int {
	file, err := os.Open("highscore.txt")
	if err != nil {
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		score, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0
		}
		return score
	}
	return 0
}

func saveHighScore(score int) {
	file, err := os.Create("highscore.txt")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%d", score)
}

func showMenu(highScore int) int {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw borders (assume 80x24 for menu)
	width, height := 80, 24
	for x := 0; x < width; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, height-1, '-', termbox.ColorWhite, termbox.ColorBlack)
	}
	for y := 0; y < height; y++ {
		termbox.SetCell(0, y, '|', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(width-1, y, '|', termbox.ColorWhite, termbox.ColorBlack)
	}

	drawText(10, 5, "Welcome to Invader CLI!", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 6, fmt.Sprintf("High Score: %d", highScore), termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 8, "Choose board size:", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 10, "1. Small (60x20)", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 12, "2. Medium (80x24)", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 14, "3. Large (100x30)", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 16, "Press 1, 2, or 3 to select, Q to quit", termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Ch == '1' {
				return 1
			} else if ev.Ch == '2' {
				return 2
			} else if ev.Ch == '3' {
				return 3
			} else if ev.Ch == 'q' || ev.Ch == 'Q' || ev.Key == termbox.KeyEsc {
				return 0
			}
		}
	}
}

func getSize(choice int) (int, int) {
	switch choice {
	case 1:
		return 60, 20
	case 2:
		return 80, 24
	case 3:
		return 100, 30
	default:
		return 80, 24
	}
}

func showDifficultyMenu() int {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw borders
	width, height := 80, 24
	for x := 0; x < width; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, height-1, '-', termbox.ColorWhite, termbox.ColorBlack)
	}
	for y := 0; y < height; y++ {
		termbox.SetCell(0, y, '|', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(width-1, y, '|', termbox.ColorWhite, termbox.ColorBlack)
	}

	drawText(10, 5, "Choose Difficulty:", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 7, "1. Easy (Slow enemy descent)", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 9, "2. Medium (Normal speed)", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 11, "3. Hard (Fast enemy descent)", termbox.ColorWhite, termbox.ColorBlack)
	drawText(10, 13, "Press 1, 2, or 3 to select", termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Ch == '1' {
				return 1
			} else if ev.Ch == '2' {
				return 2
			} else if ev.Ch == '3' {
				return 3
			}
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	highScore := loadHighScore()

	for {
		choice := showMenu(highScore)
		if choice == 0 {
			break
		}
		width, height := getSize(choice)

		difficulty := showDifficultyMenu()

		g := &game.Game{}
		g.Init(width, height, highScore, difficulty)

		inputChan := make(chan termbox.Event)
		go func() {
			for {
				ev := termbox.PollEvent()
				inputChan <- ev
			}
		}()

		running := true
		for running {
			select {
			case ev := <-inputChan:
				if ev.Type == termbox.EventKey {
					if ev.Ch == 'p' || ev.Ch == 'P' {
						g.Paused = !g.Paused
					} else if ev.Ch == 'f' || ev.Ch == 'F' {
						g.Init(width, height, highScore, g.Difficulty)
					} else if ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Ch == 'Q' {
						running = false
					} else if !g.GameOver && !g.Paused {
						g.ProcessKey(ev.Key)
					}
				}
			default:
				// No input, continue
			}

			if !g.GameOver && !g.Paused {
				g.Update()
				if g.Score > highScore {
					highScore = g.Score
					saveHighScore(highScore)
				}
			}
			g.Draw()
			time.Sleep(100 * time.Millisecond)
		}
	}
}