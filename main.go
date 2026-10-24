package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	playerX   float64
	playerY   float64
	velocityY float64
	player    *ebiten.Image
	ground    *ebiten.Image
}

func (g *Game) Update() error {
	// Move right
	if ebiten.IsKeyPressed(ebiten.KeyD) && g.playerX < 640-32 {
		g.playerX += 2
	}
	// Move left
	if ebiten.IsKeyPressed(ebiten.KeyA) && g.playerX > 0 {
		g.playerX -= 2
	}

	// Gravity
	g.velocityY += 0.4
	g.playerY += g.velocityY

	// Floor collision
	floorY := float64(360 - 42)

	if g.playerY >= floorY {
		g.playerY = floorY
		g.velocityY = 0
	}

	// Jump
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.playerY >= floorY {
		g.velocityY = -8
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{135, 206, 235, 255})

	groundOpts := &ebiten.DrawImageOptions{}
	groundOpts.GeoM.Translate(0, 360-10)

	screen.DrawImage(g.ground, groundOpts)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.playerX, g.playerY)

	screen.DrawImage(g.player, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Platformer")

	game := &Game{
		player: ebiten.NewImage(32, 32),
		ground: ebiten.NewImage(640, 10),
	}

	game.ground.Fill(color.RGBA{70, 170, 80, 255})

	game.player.Fill(color.RGBA{220, 50, 50, 225})
	game.playerY = 360 - 32

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
