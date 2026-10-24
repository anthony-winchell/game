package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Physics and dimensions
const (
	screenWidth  = 640
	screenHeight = 360

	groundHeight = 10
	playerSize   = 32

	moveSpeed = 2
	gravity   = 0.4
	jumpSpeed = 8.0

	floorY = screenHeight - groundHeight - playerSize
)

// Colors
var (
	skyColor    = color.RGBA{135, 206, 235, 255}
	groundColor = color.RGBA{70, 170, 80, 255}
	playerColor = color.RGBA{220, 50, 50, 255}
)

type Game struct {
	player    *Player
	platforms []rect
}

type Player struct {
	x        float64
	y        float64
	vx       float64
	vy       float64
	w        float64
	h        float64
	img      *ebiten.Image
	onGround bool
}

func (p *Player) rect() rect {
	return rect{p.x, p.y, p.w, p.h}
}

type rect struct {
	x, y, w, h float64
}

func (a rect) overlaps(b rect) bool {
	return (a.x < b.x+b.w) && (a.x+a.w > b.x) && (a.y < b.y+b.h) && (a.y+a.h > b.y)
}

func newPlayer(x, y float64) *Player {
	p := &Player{
		x: x, y: y, w: playerSize, h: playerSize,
	}

	p.img = ebiten.NewImage(int(p.w), int(p.h))
	p.img.Fill(playerColor)
	return p
}

func (g *Game) Update() error {
	// Move right
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.x += moveSpeed
	}
	// Move left
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.x -= moveSpeed
	}

	g.player.x = max(0, min(g.player.x, screenWidth-g.player.w))

	// Gravity
	g.player.vy += gravity
	g.player.y += g.player.vy
	g.player.onGround = false

	for _, platform := range g.platforms {
		if g.player.rect().overlaps(platform) && g.player.vy > 0 {
			g.player.y = platform.y - g.player.h
			g.player.vy = 0
			g.player.onGround = true
		}
	}

	// Jump
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.player.onGround {
		g.player.vy = -jumpSpeed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(skyColor)

	for _, p := range g.platforms {
		vector.FillRect(screen, float32(p.x), float32(p.y), float32(p.w), float32(p.h), groundColor, false)
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.x, g.player.y)

	screen.DrawImage(g.player.img, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Platformer")

	player := newPlayer(0, floorY)

	game := &Game{
		player: player,
		platforms: []rect{
			{0, screenHeight - groundHeight, screenWidth, groundHeight}, // Floor
			{100, 280, 120, 8}, // Platform
			{320, 200, 100, 8}, // Platform
			{480, 140, 80, 8},  // Platform
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
