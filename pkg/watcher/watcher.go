package watcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Watcher ..
type Watcher struct {
	dirs     []string
	interval time.Duration
	files    map[string]os.FileInfo
	isStop   bool
}

// Watch ...
func (w *Watcher) Watch(fn func([]string)) {
	for {
		if w.isStop {
			break
		}
		var paths []string
		for _, dir := range w.dirs {
			filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
				if !f.IsDir() {
					if pf, ok := w.files[path]; !ok {
						w.files[path] = f
						paths = append(paths, path)
					} else if pf.ModTime().UnixNano() != f.ModTime().UnixNano() {
						w.files[path] = f
						paths = append(paths, path)
					} else {
						// fmt.Println(path, pf.ModTime().UnixNano(), f.ModTime().UnixNano())
					}
				}
				return nil
			})
			if len(paths) > 0 {
				fn(paths)
				paths = nil
			}
		}
		time.Sleep(w.interval)
	}
	log.Println("stopping watch")
}

func (w *Watcher) Stop() {
	w.isStop = true
}

// New ...
func New(dirs string, interval time.Duration) *Watcher {
	ds := strings.Split(dirs, ",")
	files := make(map[string]os.FileInfo)
	for _, dir := range ds {
		filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				files[path] = f
			}
			return nil
		})
	}
	return &Watcher{
		dirs:     ds,
		interval: interval,
		files:    files,
		isStop:   false,
	}
}
