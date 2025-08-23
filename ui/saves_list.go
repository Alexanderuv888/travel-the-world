package ui

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// SaveFile хранит данные о сохранении
type SaveFile struct {
	Name    string
	ModTime time.Time
}

func LoadSaveFiles() (files []SaveFile) {
	entries, _ := os.ReadDir("save/saves")
	for _, e := range entries {
		if !e.IsDir() {
			info, _ := e.Info()
			name := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
			files = append(files, SaveFile{
				Name:    name,
				ModTime: info.ModTime(),
			})
		}
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime.After(files[j].ModTime)
	})
	return
}
