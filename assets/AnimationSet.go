package assets

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationSet struct {
	Animations map[string]*Animation // idle, walk, attack, die
}

type Action string
type Direction string

type Animation struct {
	Frames      map[string][]*ebiten.Image // direction → frames
	FrameRate   float64
	FrameWidth  int
	FrameHeight int
}

func GetFrame(baseFolder, tileSetName string, action Action, dir Direction, am *Manager, frameIndex int) (*ebiten.Image, error) {
	animationSet, err := LoadAnimationSet(baseFolder, tileSetName, am)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	frame := animationSet.GetFrame(fmt.Sprintf("%s/%s", tileSetName, action), dir, frameIndex)
	if frame == nil {
		log.Fatal("Не удалось сформировать фрэйм")
	}
	return frame, nil
}

func LoadAnimationSet(folder string, tileSetName string, am *Manager) (*AnimationSet, error) {
	img, err := am.LoadImage(fmt.Sprintf("%s/%s.png", folder, tileSetName))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	metaData, _ := os.ReadFile(folder + "/meta.json")

	var raw map[string]struct {
		FrameWidth  int      `json:"frame_width"`
		FrameHeight int      `json:"frame_height"`
		Directions  []string `json:"directions"`
		StartFrame  int      `json:"start_frame"`
		FrameCount  int      `json:"frame_count"`
		FPS         float64  `json:"fps"`
	}

	json.Unmarshal(metaData, &raw)

	animSet := &AnimationSet{Animations: make(map[string]*Animation)}

	for name, conf := range raw {

		anim := &Animation{
			Frames:      make(map[string][]*ebiten.Image),
			FrameRate:   conf.FPS,
			FrameWidth:  conf.FrameWidth,
			FrameHeight: conf.FrameHeight,
		}

		for i, dir := range conf.Directions {
			for f := 0; f < conf.FrameCount; f++ {
				x := (f + conf.StartFrame) * conf.FrameWidth
				y := i * conf.FrameHeight
				sub := img.SubImage(image.Rect(x, y, x+conf.FrameWidth, y+conf.FrameHeight)).(*ebiten.Image)
				anim.Frames[dir] = append(anim.Frames[dir], sub)
			}
		}

		animSet.Animations[tileSetName+"/"+name] = anim
	}
	return animSet, nil
}

func (a *AnimationSet) GetFrame(animName string, dir Direction, frameIndex int) *ebiten.Image {
	anim, ok := a.Animations[animName]
	if !ok {
		return nil
	}

	frames, ok := anim.Frames[string(dir)]
	if !ok || frameIndex < 0 || frameIndex >= len(frames) {
		return nil
	}

	return frames[frameIndex]
}

func (a *AnimationSet) GetAnimationLength(animName string, dir Direction) int {
	anim, ok := a.Animations[animName]
	if !ok {
		return 0
	}
	return len(anim.Frames[string(dir)])
}

func (a *AnimationSet) GetAnimationFrameRate(animName string) int {
	anim, ok := a.Animations[animName]
	if !ok {
		return 0
	}
	return int(anim.FrameRate)
}

func ASName(folder, tileSetName string) string {
	return fmt.Sprintf("%s/%s", folder, tileSetName)
}
