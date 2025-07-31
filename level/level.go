package level

import (
	"errors"
	"image"
	"log"
	"math/rand/v2"
	"path/filepath"
	"travel-the-world/assets"
	"travel-the-world/tiles"
	"travel-the-world/unit"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

const (
	levelsPath = "assets/levels"
	padding    = 0
)

type Level struct {
	name         string
	W, H         int
	TileW, TileH int
	needUpdate   bool
	am           *assets.Manager
	background   *ebiten.Image
	tmap         *tiled.Map
	Objects      []*tiles.ObjectTile
	Units        []*unit.Unit
}

func NewLevel(name string, am *assets.Manager) (*Level, error) {
	fileName := name + ".tmx"

	levelPath := filepath.Join(levelsPath, name, fileName)

	tmap, err := tiled.LoadFile(levelPath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	background := ebiten.NewImage(tmap.Width*tmap.TileWidth+padding, tmap.Height*tmap.TileHeight+padding)

	l := &Level{
		name:       name,
		W:          tmap.Width,
		H:          tmap.Height,
		TileW:      tmap.TileWidth,
		TileH:      tmap.TileHeight,
		needUpdate: true,
		am:         am,
		background: background,
		tmap:       tmap,
	}
	return l, nil
}

func (l *Level) Size() (width, height int) {
	return l.W, l.H
}

func (l *Level) AddObject(tx, ty float64, w, h int, img *ebiten.Image) {
	l.Objects = append(l.Objects, tiles.NewObjectTile(tx, ty, w, h, img))
}

func (l *Level) CreateUnits(amount int) {
	for range make([]struct{}, amount) {
		x, y := l.getRandomCoordinates()
		unit := unit.NewUnit(x, y, l.am, &unit.IdleBehavior{})
		l.Units = append(l.Units, unit)
	}
}

func (l *Level) getRandomCoordinates() (X, Y float64) {
	x := rand.IntN(l.W - 5)
	y := rand.IntN(l.H - 5)
	X = float64((x-y)*(l.TileW/2)) + float64(l.W*l.TileW/2)
	Y = float64((x+y)*(l.TileH/2)) + padding
	return X, Y
}

func (l *Level) drawIsoLayer(background *ebiten.Image, layer *tiled.Layer, ts *tiles.TilesetXML, firstGID uint32) {
	tiles := ts.GetSlices()
	offsetX := ts.GetIntProperty("offsetX")
	offsetY := ts.GetIntProperty("offsetY")
	for y := 0; y < l.H; y++ {
		for x := 0; x < l.W; x++ {
			tile := layer.Tiles[y*l.W+x]
			if tile.Nil {
				continue
			}

			// Перевод координат в изометрию
			screenX := float64((x-y)*(l.TileW/2)+offsetX) + float64(l.W*l.TileW/2-l.TileW/2)
			screenY := float64((x+y)*(l.TileH/2)+offsetY) + padding
			if layer.Name == "tree" {
				l.AddObject(screenX, screenY, l.TileW, l.TileH, tiles[tile.ID])
			} else {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(screenX, screenY)
				background.DrawImage(tiles[tile.ID], op)
			}
		}
	}
}

func (l *Level) DrawLevel(screen *ebiten.Image, cameraPos image.Point) {
	if l.needUpdate {
		l.Update(cameraPos)
		l.needUpdate = false
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(cameraPos.X), -float64(cameraPos.Y))
	screen.DrawImage(l.background, op)
}

func (l *Level) Update(cameraPos image.Point) {
	for _, layer := range l.tmap.Layers {
		if !layer.Visible {
			continue
		}
		for _, tilesetName := range layer.Properties.Get("TilesetName") {
			firstGID, err := findFirstGID(l.tmap, tilesetName)
			if err != nil {
				log.Fatal(err)
				continue
			}
			tsxPath := filepath.Join(levelsPath, l.name, tilesetName)
			ts, err := tiles.LoadTilesetFromTSX(tsxPath)
			if err != nil {
				log.Fatal(err)
			}
			l.drawIsoLayer(l.background, layer, ts, firstGID)
		}
	}
}

func findFirstGID(m *tiled.Map, fileName string) (uint32, error) {
	for _, tileset := range m.Tilesets {
		if tileset.Source == fileName {
			return tileset.FirstGID, nil
		}
	}
	return 0, errors.New("Tileset " + fileName + " not found.")
}

func (l *Level) Width() float64 {
	return float64(l.W * l.TileW)
}

func (l *Level) Height() float64 {
	return float64(l.H * l.TileH)
}

func (l *Level) WidthInt() int {
	return l.W * l.TileW
}

func (l *Level) HeightInt() int {
	return l.H * l.TileH
}
