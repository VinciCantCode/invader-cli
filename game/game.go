package game

import (
	"math/rand"
	"time"
)

func (g *Game) Init(width, height, highScore, difficulty int) {
	rand.Seed(time.Now().UnixNano())
	g.Width = width
	g.Height = height
	playerMaxY := height - 2
	g.Player = Player{Pos: Position{width / 2, playerMaxY}, HP: 3}
	g.Aliens = []Alien{}
	// Initialize aliens in 5x8 formation, adjusted for width
	startX := 10
	endX := startX + 35  // 8 columns * 5 -5 =35
	if endX > width-10 {
		endX = width - 10
		startX = 5
	}
	for y := 2; y <= 6; y++ {
		for x := startX; x <= endX; x += 5 {
			g.Aliens = append(g.Aliens, Alien{Pos: Position{x, y}})
		}
	}
	g.Bullets = []Bullet{}
	g.Running = true
	g.AlienDir = 1
	g.AlienMoveCounter = 0
	g.Score = 0
	g.HighScore = highScore
	g.PlayerHit = false
	g.HitCounter = 0
	g.GameOver = false
	g.Difficulty = difficulty
	g.Paused = false
}

func (g *Game) Update() {
	// Move aliens
	g.AlienMoveCounter++
	if g.AlienMoveCounter >= 10 { // Move every 10 frames
		g.AlienMoveCounter = 0
		edge := false
		for i := range g.Aliens {
			alien := &g.Aliens[i]
			alien.Pos.X += g.AlienDir
			if alien.Pos.X <= 1 || alien.Pos.X >= g.Width-2 {
				edge = true
			}
		}
		if edge {
			g.AlienDir *= -1
			descent := 1
			if g.Difficulty == 2 {
				descent = 2
			} else if g.Difficulty == 3 {
				descent = 3
			}
			for i := range g.Aliens {
				g.Aliens[i].Pos.Y += descent
			}
		}

		// Alien shooting
		if len(g.Aliens) > 0 && rand.Intn(100) < 40 {
			shooter := g.Aliens[rand.Intn(len(g.Aliens))]
			g.Bullets = append(g.Bullets, Bullet{Pos: Position{shooter.Pos.X, shooter.Pos.Y + 1}, DirY: 1})
		}
	}

	// Move bullets
	for i := len(g.Bullets) - 1; i >= 0; i-- {
		bullet := &g.Bullets[i]
		bullet.Pos.Y += bullet.DirY
		if bullet.Pos.Y < 0 || bullet.Pos.Y >= g.Height {
			g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
		}
	}

	// Check collisions
	for i := len(g.Bullets) - 1; i >= 0; i-- {
		bullet := g.Bullets[i]
		if bullet.DirY == -1 { // Player bullet
			for j := len(g.Aliens) - 1; j >= 0; j-- {
				alien := g.Aliens[j]
				if bullet.Pos.X == alien.Pos.X && bullet.Pos.Y == alien.Pos.Y {
					g.Aliens = append(g.Aliens[:j], g.Aliens[j+1:]...)
					g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
					g.Score += 10
					break
				}
			}
		} else { // Alien bullet
			if bullet.Pos.X == g.Player.Pos.X && bullet.Pos.Y == g.Player.Pos.Y {
				g.Player.HP--
				g.PlayerHit = true
				g.HitCounter = 10
				if g.Player.HP <= 0 {
					g.GameOver = true
					g.Running = false
				}
				g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
				break
			}
		}
	}

	// Check if aliens reach bottom or collide with player
	for _, alien := range g.Aliens {
		if alien.Pos.Y >= g.Height-2 || (alien.Pos.X == g.Player.Pos.X && alien.Pos.Y == g.Player.Pos.Y) {
			g.GameOver = true
			g.Running = false
		}
	}

	// Handle hit counter
	if g.HitCounter > 0 {
		g.HitCounter--
		if g.HitCounter == 0 {
			g.PlayerHit = false
		}
	}

	// Check for high score update
	if g.Score > g.HighScore {
		g.HighScore = g.Score
		// Save to file (will do in main)
	}

	// If no aliens left, win or restart
	if len(g.Aliens) == 0 {
		g.GameOver = true
		g.Running = false
	}
}