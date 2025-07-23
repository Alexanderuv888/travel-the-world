package game

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"travel-the-world/assets"
	"travel-the-world/tiles"
	"travel-the-world/unit"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	w, h, int,
	AssetsManager *assets.Manager
	CurrentLevel *Level
	Unit         *unit.Unit

	Camera *Camera

	camX, camY float64
	camScale   float64
	camScaleTo float64

	mousePanX, mousePanY int

	offscreen *ebiten.Image
}

// NewGame returns a new isometric demo Game.
func NewGame() (*Game, error) {
	assetsManager := assets.NewManager()

	l, err := NewLevel("level_2", assetsManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create new level: %s", err)
	}
	//camera := NewCamera(0, float64(l.H*l.TileH/2))
	camera := NewCamera(l.Width()/5, l.Height()/6)
	unit := unit.NewUnit(0, l.Height()/2, assetsManager)

	g := &Game{
		AssetsManager: assetsManager,
		CurrentLevel:  l,
		Unit:          unit,
		Camera:        camera,
		camScale:      1,
		camScaleTo:    1,
		mousePanX:     math.MinInt32,
		mousePanY:     math.MinInt32,
	}
	return g, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentLevel.drawLevel(screen, g.Camera)
	dq := &tiles.DrawQueue{}

	g.Unit.Render(dq)

	for _, unit := range g.CurrentLevel.units {
		unit.Render(dq)
	}

	for _, obj := range g.CurrentLevel.objects {
		dq.Add(obj)
	}

	// потом:
	dq.DrawAll(screen, g.Camera.X, g.Camera.Y)
	dq.Clear()
	g.drawDebugInfo(screen)

}

func (g *Game) drawDebugInfo(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("KEYS WASD EC R\nFPS  %0.0f\nTPS  %0.0f\nangle  %0.2f\ntopAngle  %0.2f\nbottomAngle  %0.2f\nUnitPOS  %0.0f,%0.0f\nmousePOS  %0.0f,%0.0f", ebiten.ActualFPS(), ebiten.ActualTPS(), g.Unit.Angle, g.Unit.TopAngle, g.Unit.BottomAngle, g.Unit.X, g.Unit.Y, float64(mouseX)+g.Camera.X, float64(mouseY)+g.Camera.Y))

	x1, y1 := g.Unit.X-g.Camera.X, g.Unit.Y-g.Camera.Y             //WorldToIso(g.Unit.X, g.Unit.Y, 64, 32, g.Camera.X, g.Camera.Y)
	x2, y2 := g.Unit.GoalX()-g.Camera.X, g.Unit.GoalY()-g.Camera.Y //WorldToIso(g.Unit.GoalX, g.Unit.GoalY, 64, 32, g.Camera.X, g.Camera.Y)
	vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, color.RGBA{255, 0, 0, 255}, false)
}

func (g *Game) Layout(screenWidth, screenHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.listenKeyBoardAndMouse()
	levelDimentions := image.Point{g.CurrentLevel.WidthInt(), g.CurrentLevel.HeightInt()}

	g.Unit.Update(g.CurrentLevel.objects, g.CurrentLevel.units, levelDimentions)
	for _, npc := range g.CurrentLevel.units {
		npc.Update(g.CurrentLevel.objects, g.CurrentLevel.units, levelDimentions)
	}
	return nil
}

func (g *Game) listenKeyBoardAndMouse() {
	g.Camera.Update()
	g.Unit.ListenKeyBoard(g.Camera.Pos())
}

func WorldToScreenIso(x, y float64, tileW, tileH int, cameraX, cameraY float64) (float64, float64) {
	sx := (x - y)
	sy := (x + y)
	return sx - cameraX, sy - cameraY
}

func WorldToIso(x, y float64, tileW, tileH int, cameraX, cameraY float64) (isoX, isoY float64) {
	isoX = (x - y) + float64(50*tileW/2) - cameraX
	isoY = (x + y) - cameraY

	return
}
