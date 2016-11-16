package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsevents"
)

func main() {
	path := "."
	dev, err := fsevents.DeviceForPath(path)
	if err != nil {
		log.Fatal(err)
	}
	es := &fsevents.EventStream{
		Paths:   []string{path},
		Latency: 500 * time.Millisecond,
		Device:  dev,
		Flags:   fsevents.FileEvents | fsevents.WatchRoot,
	}
	es.Start()
	go func() {
		for msg := range es.Events {
			for _, event := range msg {
				fmt.Println(event.Flags)
			}
		}
	}()
	time.Sleep(time.Minute)
}
