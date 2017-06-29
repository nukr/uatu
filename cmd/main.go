package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nukr/uatu/pkg/commander"
	"github.com/nukr/uatu/pkg/watcher"
)

// uatu -s 'go run cmd/main.go' -d 'cmd,pkg'
func main() {
	dirs := flag.String("d", "", "")
	script := flag.String("s", "", "")
	flag.Parse()
	if *dirs == "" {
		log.Fatal("-d is required")
	}
	if *script == "" {
		log.Fatal("-s is required")
	}
	w := watcher.New(*dirs, time.Second*2)
	c := commander.New(*script)
	c.Run()
	w.Watch(func(paths []string) {
		fmt.Println(paths, " is changed")
		c.Run()
	})
}
