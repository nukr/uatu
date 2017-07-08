package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/nukr/uatu/pkg/commander"
	"github.com/nukr/uatu/pkg/watcher"
)

// uatu -s 'go run cmd/main.go' -d 'cmd,pkg'
func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	dirs := flag.String("d", "", "")
	script := flag.String("s", "", "")
	version := flag.Bool("v", false, "version")
	flag.Parse()
	if *dirs == "" {
		logger.Log("message", "-d is required")
		os.Exit(1)
	}
	if *script == "" {
		logger.Log("message", "-s is required")
		os.Exit(1)
	}
	if *version {
		logger.Log("message", "0.0.1")
		os.Exit(0)
	}
	w := watcher.New(*dirs, time.Second*2)
	c := commander.New(*script)
	err := c.Run()
	if err != nil {
		logger.Log("error", err)
	}
	w.Watch(func(paths []string) {
		fmt.Println(paths, " is changed")
		err := c.Run()
		if err != nil {
			logger.Log("error", err)
		}
	})
}
