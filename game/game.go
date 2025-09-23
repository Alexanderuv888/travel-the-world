package game

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"travel-the-world/assets"
	"travel-the-world/camera"
	"travel-the-world/common"
	"travel-the-world/hud"
	"travel-the-world/level"
	"travel-the-world/save"
	"travel-the-world/tiles"
	"travel-the-world/ui"
	"travel-the-world/unit"
	"travel-the-world/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	state         GameState
	Menu          *ui.Menu
	AssetsManager *assets.Manager
	CurrentLevel  *level.Level
	WorldCtx      *world.Context
	Player        *unit.Unit
	hud           *hud.HUD

	Camera *camera.Camera

	mousePanX, mousePanY int
}

func NewGame() (*Game, error) {
	menu := ui.NewMenu()
	assetsManager := assets.NewManager()

	l, err := level.NewLevel("level_2", assetsManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create new level: %s", err)
	}
	camera := camera.NewCamera(l.Width()/5, l.Height()/6)
	unit := unit.NewUnit(1600, 800, assetsManager, &unit.PlayerBehavior{})

	l.CreateUnits(10)
	l.Units = append(l.Units, unit)
	ilist := &common.InteractableList{}
	for _, npc := range l.Units {
		ilist.Add(npc)
	}

	worldCtx := &world.Context{InteractableList: *ilist, Camera: camera}

	g := &Game{
		state:         StateMenu,
		Menu:          menu,
		hud:           hud.NewHUD(unit),
		AssetsManager: assetsManager,
		CurrentLevel:  l,
		WorldCtx:      worldCtx,
		Player:        unit,
		Camera:        camera,
		mousePanX:     math.MinInt32,
		mousePanY:     math.MinInt32,
	}
	menu.SetStartCallback(func() {
		g.state = StatePlaying
	})

	menu.SetExitCallback(func() {
		g.state = StateQuitGame
	})
	menu.SetLoadGameCallback(func() {
		menu.SaveLoadMenu.SetFiles(ui.LoadSaveFiles())
		menu.SaveLoadMenu.SetVisible(true)
		menu.SaveLoadMenu.SetButtons(ui.LoadMenuButtons)
		//menu.SaveMenu.Visible = false
	})
	menu.SetSaveGameCallback(func() {
		menu.SaveLoadMenu.SetFiles(ui.LoadSaveFiles())
		menu.SaveLoadMenu.SetVisible(true)
		menu.SaveLoadMenu.SetButtons(ui.SaveMenuButtons)
		//menu.SaveMenu.LoadFiles()
		//menu.SaveMenu.Visible = true
		//menu.SaveLoadMenu.SetVisible(false)
	})
	ui.SetLoadHandler(func(name string) {
		saveData, err := save.LoadGame(name)
		if err != nil {
			fmt.Println("Error loading game:", err)
			return
		}
		// Копируем состояние загруженной игры в текущую игру
		g.Player.LoadStats(saveData)
		g.state = StatePlaying
	})
	ui.SetSaveGameHandler(func(name string) {
		// формируем путь, например save/saves/<name>.json
		path := "save/saves/" + name
		// если хотите добавлять расширение:
		// if !strings.HasSuffix(path, ".json") { path += ".json" }
		g.SaveProgress(path)
		// обновить список файлов в меню, если меню открыто:
		menu.SaveMenu.LoadFiles()
	})
	return g, nil
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	g.CurrentLevel.DrawLevel(screen, g.Camera)
	dq := &tiles.DrawQueue{}

	for _, unit := range g.CurrentLevel.Units {
		unit.Render(dq)
	}

	for _, obj := range g.CurrentLevel.Objects {
		dq.Add(obj)
	}

	dq.DrawAll(screen, g.Camera, g.Player.Point(), int(g.Player.Stats.Vision/2.5))
	dq.Clear()
	g.hud.Draw(screen)
	g.drawDebugInfo(screen)

}

func (g *Game) drawDebugInfo(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("KEYS WASD EC R\nFPS  %0.0f\nTPS  %0.0f\nangle  %0.2f\ntopAngle  %0.2f\nbottomAngle  %0.2f\nUnitPOS  %0.0f,%0.0f\nmousePOS  %0.0f,%0.0f", ebiten.ActualFPS(), ebiten.ActualTPS(), g.Player.Angle, g.Player.TopAngle, g.Player.BottomAngle, g.Player.X, g.Player.Y, float64(mouseX)+g.Camera.X, float64(mouseY)+g.Camera.Y))

	x1, y1 := g.Camera.WorldToScreen(g.Player.X, g.Player.Y)             //WorldToIso(g.Player.X, g.Player.Y, 64, 32, g.Camera.X, g.Camera.Y)
	x2, y2 := g.Camera.WorldToScreen(g.Player.GoalX(), g.Player.GoalY()) //WorldToIso(g.Player.GoalX, g.Player.GoalY, 64, 32, g.Camera.X, g.Camera.Y)
	vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, color.RGBA{255, 0, 0, 255}, false)
}

func (g *Game) Layout(screenWidth, screenHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) UpdateGame() error {
	g.listenKeyBoardAndMouse()
	ilist := unitToIntarectableList(g.CurrentLevel.Units)
	g.WorldCtx.Update(ilist)
	levelDimentions := image.Point{g.CurrentLevel.WidthInt(), g.CurrentLevel.HeightInt()}

	for _, unit := range g.CurrentLevel.Units {
		unit.Update(g.WorldCtx, levelDimentions)
	}
	return nil
}

func (g *Game) listenKeyBoardAndMouse() {
	c := g.Camera
	c.Update()
	x, y := ebiten.CursorPosition()
	p := c.ScreenToWorldPoint(x, y)
	for _, obj := range g.WorldCtx.InteractableList.Items {
		if p.In(obj.Rect()) {
			obj.Highlight()
		}
	}
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

func unitToIntarectableList(units []*unit.Unit) common.InteractableList {
	ilist := &common.InteractableList{}
	for _, npc := range units {
		ilist.Add(npc)
	}
	return *ilist
}
