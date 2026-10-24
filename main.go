package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	player *Player
	ground *ebiten.Image
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

	if g.player.y >= floorY {
		g.player.y = floorY
		g.player.vy = 0
		g.player.onGround = true
	}

	// Jump
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.player.onGround {
		g.player.vy = -jumpSpeed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(skyColor)

	groundOpts := &ebiten.DrawImageOptions{}
	groundOpts.GeoM.Translate(0, screenHeight-groundHeight)

	screen.DrawImage(g.ground, groundOpts)

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
		ground: ebiten.NewImage(screenWidth, groundHeight),
	}

	game.ground.Fill(groundColor)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
