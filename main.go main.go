package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playerX float64
	playerY float64
	player  *ebiten.Image
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

	// Move up
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.playerY > 0 {
		g.playerY -= 2
	}

	// Move down
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.playerY < 360-32 {
		g.playerY += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{135, 206, 235, 255})

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
	}
	game.player.Fill(color.RGBA{220, 50, 50, 225})

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
