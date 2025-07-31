package assets

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Manager struct {
	animationSet map[string]*AnimationSet
	images       map[string]*ebiten.Image
	soundsFiles  map[string]*[]byte
	fonts        map[string]font.Face

	audioCtx *audio.Context
}

const (
	sr48000 int = 48000
	sr44100 int = 44100
)

func NewManager() *Manager {
	audioCtx := audio.NewContext(sr48000)

	return &Manager{
		animationSet: make(map[string]*AnimationSet),
		images:       make(map[string]*ebiten.Image),
		soundsFiles:  make(map[string]*[]byte),
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

func (a *Manager) loadSound(path string) (*[]byte, error) {
	if s, ok := a.soundsFiles[path]; ok {
		return s, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	a.soundsFiles[path] = &data
	return &data, nil
}

func (a *Manager) CreatePlyerFor(path string) (*audio.Player, error) {
	data, err := a.loadSound(path)
	if err != nil {
		return nil, err
	}

	stream, err := wav.DecodeWithSampleRate(a.audioCtx.SampleRate(), bytes.NewReader(*data))
	if err != nil {
		return nil, err
	}

	player, err := a.audioCtx.NewPlayer(stream)
	if err != nil {
		return nil, err
	}

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
