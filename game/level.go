package game

import (
	"errors"
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
	objects      []*tiles.ObjectTile
	units        []*unit.Unit
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
	l.createUnits(50)
	return l, nil
}

func (l *Level) Size() (width, height int) {
	return l.W, l.H
}

func (l *Level) addObject(tx, ty float64, w, h int, img *ebiten.Image) {
	l.objects = append(l.objects, tiles.NewObjectTile(tx, ty, w, h, img))
}

func (l *Level) createUnits(amount int) {
	for range make([]struct{}, amount) {
		x, y := l.getRandomCoordinates()
		unit := unit.NewUnit(x, y, l.am)
		l.units = append(l.units, unit)
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
				l.addObject(screenX, screenY, l.TileW, l.TileH, tiles[tile.ID])
			} else {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(screenX, screenY)
				background.DrawImage(tiles[tile.ID], op)
			}
		}
	}
}

func (l *Level) drawLevel(screen *ebiten.Image, camera *Camera) {
	if l.needUpdate {
		l.update(camera)
		l.needUpdate = false
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-camera.X, -camera.Y)
	screen.DrawImage(l.background, op)
}

func (l *Level) update(camera *Camera) {
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
