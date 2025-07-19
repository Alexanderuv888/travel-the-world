package assets

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Manager struct {
	animationSet map[string]*AnimationSet
	images       map[string]*ebiten.Image
	sounds       map[string]*audio.Player
	fonts        map[string]font.Face

	audioCtx *audio.Context
}

func NewManager() *Manager {
	audioCtx, err := audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		animationSet: make(map[string]*AnimationSet),
		images:       make(map[string]*ebiten.Image),
		sounds:       make(map[string]*audio.Player),
		fonts:        make(map[string]font.Face),
		audioCtx:     audioCtx,
	}
}

func (a *Manager) LoadAnimationSet(folder string, tileSetName string) (*AnimationSet, error) {
	animationSetName := fmt.Sprintf("%s/%s", folder, tileSetName)
	if as, ok := a.animationSet[animationSetName]; ok {
		return as, nil
	}
	as, err := LoadAnimationSet(folder, tileSetName, a)
	if err != nil {
		return nil, err
	}
	a.animationSet[animationSetName] = as
	return as, nil
}

func (a *Manager) LoadImage(path string) (*ebiten.Image, error) {
	if img, ok := a.images[path]; ok {
		return img, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	img := ebiten.NewImageFromImage(src)

	a.images[path] = img
	return img, nil
}

func (a *Manager) LoadSound(path string) (*audio.Player, error) {
	if s, ok := a.sounds[path]; ok {
		return s, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stream, err := wav.Decode(a.audioCtx, f)
	if err != nil {
		return nil, err
	}
	player, err := audio.NewPlayer(a.audioCtx, stream)
	if err != nil {
		return nil, err
	}
	a.sounds[path] = player
	return player, nil
}

func (a *Manager) LoadFont(path string, size float64) (font.Face, error) {
	key := path + ":" + strconv.FormatFloat(size, 'f', 2, 64)
	if f, ok := a.fonts[key]; ok {
		return f, nil
	}

	fdata, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	tt, err := opentype.Parse(fdata)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	a.fonts[key] = face
	return face, nil
}
