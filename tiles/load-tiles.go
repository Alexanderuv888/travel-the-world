package tiles

import (
	"encoding/xml"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

// /Users/vols/Projects/travel-the-world/load-tiles.go
// TilesetXML — структура для парсинга .tsx
type TilesetXML struct {
	tsxPath string
	XMLName xml.Name `xml:"tileset"`
	TileW   int      `xml:"tilewidth,attr"`
	TileH   int      `xml:"tileheight,attr"`
	Img     Image    `xml:"image"`
	Tiles   []struct {
		id  int   `xml:"id,attr"`
		Img Image `xml:"image"`
	} `xml:"tile"`
	Properties []Property `xml:"properties>property"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type Image struct {
	Source string `xml:"source,attr"`
	Trans  string `xml:"trans,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

func LoadTilesetFromTSX(tsxPath string) (*TilesetXML, error) {
	// 1. Открываем .tsx
	data, err := os.ReadFile(tsxPath)
	if err != nil {
		return nil, err
	}

	// 2. Парсим XML
	var ts TilesetXML
	if err := xml.Unmarshal(data, &ts); err != nil {
		return nil, err
	}
	ts.tsxPath = tsxPath

	return &ts, nil
}

func (ts *TilesetXML) GetSlices() []*ebiten.Image {
	if ts.GetProperty("tilesetType") == "images" {
		return getImages(ts.tsxPath, ts)
	} else {
		return sliceTSX(ts.tsxPath, ts)
	}
}

func (ts *TilesetXML) GetProperty(name string) string {
	for _, p := range ts.Properties {
		if p.Name == name {
			return p.Value
		}
	}
	return ""
}

func (ts *TilesetXML) GetIntProperty(name string) int {
	for _, p := range ts.Properties {
		if p.Name == name {
			value, err := strconv.Atoi(p.Value)
			if err != nil {
				log.Fatal(err)
				return 0
			}
			return value
		}
	}
	return 0
}

func getImages(tsxPath string, ts *TilesetXML) []*ebiten.Image {
	tilesetDir := filepath.Dir(tsxPath)

	var tiles []*ebiten.Image
	for _, tile := range ts.Tiles {
		pngPath := filepath.Join(tilesetDir, tile.Img.Source)
		img, err := getEbitenImage(pngPath, tile.Img.Trans)
		if err != nil {
			log.Fatal(err)
			continue
		}
		tiles = append(tiles, img)
	}
	return tiles
}

func sliceTSX(tsxPath string, ts *TilesetXML) []*ebiten.Image {
	tilesetDir := filepath.Dir(tsxPath)
	pngPath := filepath.Join(tilesetDir, ts.Img.Source)
	img, err := getEbitenImage(pngPath, ts.Img.Trans)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return sliceTiles(ebiten.NewImageFromImage(img), ts.TileW, ts.TileH)
}

func getEbitenImage(path string, transColor string) (*ebiten.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	if transColor != "" {
		transparentColor := parseHexColor(transColor)
		img = makeColorTransparent(img, transparentColor)
	}

	return ebiten.NewImageFromImage(img), nil
}

// Удаляет указанный цвет, делая его прозрачным
func makeColorTransparent(src image.Image, transparent color.Color) *ebiten.Image {
	bounds := src.Bounds()
	dst := image.NewNRGBA(bounds)

	tr, tg, tb, _ := transparent.RGBA()
	tr >>= 8
	tg >>= 8
	tb >>= 8

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			r >>= 8
			g >>= 8
			b >>= 8
			a >>= 8

			if r == tr && g == tg && b == tb {
				dst.Set(x, y, color.NRGBA{0, 0, 0, 0})
			} else {
				dst.Set(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
			}
		}
	}
	return ebiten.NewImageFromImage(dst)
}

// Преобразует строку "00ff00" в color.Color
func parseHexColor(s string) color.Color {
	if len(s) != 6 {
		return color.RGBA{0, 255, 0, 255}
	}
	r, _ := strconv.ParseUint(s[0:2], 16, 8)
	g, _ := strconv.ParseUint(s[2:4], 16, 8)
	b, _ := strconv.ParseUint(s[4:6], 16, 8)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

// Разрезает изображение на тайлы
func sliceTiles(img *ebiten.Image, tileW, tileH int) []*ebiten.Image {
	var tiles []*ebiten.Image
	w, h := img.Size()
	for y := 0; y < h; y += tileH {
		for x := 0; x < w; x += tileW {
			tile := img.SubImage(image.Rect(x, y, x+tileW, y+tileH)).(*ebiten.Image)
			tiles = append(tiles, tile)
		}
	}
	return tiles
}
