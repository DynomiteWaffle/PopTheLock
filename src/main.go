package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	score         int
	pickPositon   int //rotation
	pinPosition   int // 8 states - 45 deg appart
	pinMaxCount   int
	degreeFreedom float32 // amount of error you can have on input
	debugText     string
	debugshow     bool
	debugAuto     bool
	// window size
	windowWidth        int
	windowCenterWidth  int
	windowHeight       int
	windowCenterHeight int
	windowUnit         int

	lockBodyColor    color.Color
	lockShackleColor color.Color
	lockDialColor    color.Color
}

func (g *Game) Update() error {
	// Debug / settings
	//
	//
	g.debugText = ""
	// window size
	g.debugText += "Window:\nSize " + strconv.Itoa(g.windowHeight) + "x" + strconv.Itoa(g.windowWidth) + "\nUnit " + strconv.Itoa(g.windowUnit) + "\n"

	g.debugText += "\nMisc:\n"
	g.debugText += "Score " + strconv.Itoa(g.score) + "\n"
	g.debugText += "Difficulty " + strconv.Itoa(g.pinMaxCount) + "\n"
	g.debugText += "Accuracy " + strconv.FormatFloat(float64(g.degreeFreedom), 'f', 1, 32) + "\n"
	g.debugText += "Auto " + strconv.FormatBool(g.debugAuto) + "\n"

	// adjust bool settings
	toggleBool(inpututil.IsKeyJustPressed(ebiten.KeyD), &g.debugshow)
	toggleBool(inpututil.IsKeyJustPressed(ebiten.KeyA), &g.debugAuto)

	// acuracy controll
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			g.degreeFreedom += 0.1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			g.degreeFreedom -= 0.1
		}
		// accuracy cap
		if g.degreeFreedom < 0 {
			g.degreeFreedom = 0
		}
	} else {
		// adjust pin count

		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			if ebiten.IsKeyPressed(ebiten.KeyShift) {
				g.pinMaxCount += 10

			} else {

				g.pinMaxCount++
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			if ebiten.IsKeyPressed(ebiten.KeyShift) {
				g.pinMaxCount -= 10

			} else {

				g.pinMaxCount--
			}
		}
		// cap pin count
		if g.pinMaxCount > 360 {
			g.pinMaxCount = 360
		} else if g.pinMaxCount < 1 {
			g.pinMaxCount = 1
		}
	}

	// Game
	//
	//
	// TODO prosses click
	// activate pin - mouse click or keyboard space
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		g.debugText += "Clicked\n"
	}

	return nil

}
func (g *Game) Draw(screen *ebiten.Image) {
	if g.debugshow {
		ebitenutil.DebugPrint(screen, g.debugText)
	}

	lockCenter := g.windowUnit * 2

	// TODO draw lock
	// shackle
	vector.StrokeCircle(screen, float32(g.windowCenterWidth), float32(lockCenter)*0.8, float32(g.windowUnit)/2, float32(g.windowUnit)/8, g.lockShackleColor, true)

	// lock body
	vector.DrawFilledCircle(screen, float32(g.windowCenterWidth), float32(lockCenter), float32(g.windowUnit)/2, g.lockBodyColor, true)
	vector.DrawFilledCircle(screen, float32(g.windowCenterWidth), float32(lockCenter), float32(g.windowUnit)/2.7, g.lockDialColor, true)
	vector.DrawFilledCircle(screen, float32(g.windowCenterWidth), float32(lockCenter), float32(g.windowUnit)/6, g.lockBodyColor, true)

	// TODO draw pick
	// TOOD draw pin
	// TODO draw score - with text.draw() - needs font
}
func (g *Game) Layout(outWidth, outHeight int) (screenWidth, screenHeight int) {
	g.windowWidth = outWidth
	g.windowHeight = outHeight

	g.windowCenterWidth = outWidth / 2
	g.windowCenterHeight = outHeight / 2

	// find smallest window unit
	if g.windowHeight/3 < g.windowWidth {
		g.windowUnit = g.windowHeight / 3
	} else {
		g.windowUnit = g.windowWidth
	}

	return outWidth, outHeight
}

func main() {

	ebiten.SetWindowTitle("Pop The Lock")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := &Game{}
	game.pinMaxCount = 8
	game.degreeFreedom = 1
	// colors
	game.lockBodyColor = color.RGBA{A: 255}
	game.lockShackleColor = color.RGBA{R: 100, B: 100, G: 100, A: 255}
	game.lockDialColor = color.RGBA{R: 255, B: 255, G: 255, A: 255}

	game.debugshow = true //TODO remove
	ebiten.RunGame(game)
}

func toggleBool(rule bool, item *bool) {
	if rule {
		if *item {
			*item = false
		} else {
			*item = true
		}
	}
}
