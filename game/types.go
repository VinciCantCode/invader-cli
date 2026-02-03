package game

const (
	playerChar = 'A'
	alienChar  = 'V'
	bulletChar = '|'
)

type Position struct {
	X, Y int
}

type Player struct {
	Pos Position
	HP  int
}

type Alien struct {
	Pos Position
}

type Bullet struct {
	Pos  Position
	DirY int // -1 for up, 1 for down
}

type Game struct {
	Player          Player
	Aliens          []Alien
	Bullets         []Bullet
	Running         bool
	AlienDir        int
	AlienMoveCounter int
	Score           int
	HighScore       int
	PlayerHit       bool
	HitCounter      int
	GameOver        bool
	Width           int
	Height          int
	Difficulty      int // 1 easy, 2 medium, 3 hard
	Paused          bool
}