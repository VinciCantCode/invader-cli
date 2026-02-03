# invader-cli

A simple CLI-based Space Invaders game written in Go.

## How to Play

- Start the game and choose a board size: Small (60x20), Medium (80x24), or Large (100x30).
- Then choose difficulty: Easy (slow enemy descent), Medium (normal), Hard (fast descent).
- Use arrow keys to move the spaceship (A) left, right, up, down within the boundaries.
- The spaceship can float up and down in a small area at the bottom.
- Press spacebar to shoot.
- Defend against incoming aliens (V) that move automatically and fire frequently.
- Destroy all aliens to win.
- Avoid alien bullets and don't let aliens reach the bottom.
- You have 3 HP; each hit by an alien bullet reduces HP by 1 and flashes the screen red.
- Game ends when HP reaches 0 or aliens reach the bottom, displaying "Game Over! Press F to restart".
- The game area is enclosed by borders.

## Controls

- Left Arrow: Move left
- Right Arrow: Move right
- Up Arrow: Move up
- Down Arrow: Move down
- Space: Shoot (single-shot projectiles that move forward automatically)
- P: Pause/Resume game
- F: Restart game
- Q or Esc: Quit

## Game Interface

- Score: Displays current score
- High: Displays the all-time high score
- HP: Displays current health points (starts at 3)
- Screen flashes red when hit by alien bullets

## Building and Running

```bash
go build -o invader
./invader
```

## Project Structure

- `main.go`: Entry point
- `game/`: Game logic package
  - `types.go`: Constants and structs
  - `game.go`: Game initialization and update logic
  - `input.go`: Input handling
  - `rendering.go`: Drawing and rendering

## Dependencies

- github.com/nsf/termbox-go
