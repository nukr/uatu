package commander

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Commander ...
type Commander struct {
	script string
	cmd    *exec.Cmd
}

// Run ...
func (c *Commander) Run() {
	if c.cmd != nil {
		c.cmd.Process.Kill()
	}
	script := strings.Split(c.script, " ")
	c.cmd = exec.Command(script[0], script[1:]...)
	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = c.cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		io.Copy(os.Stdout, stdout)
	}()
}

// New ...
func New(script string) *Commander {
	return &Commander{
		script: script,
	}
}
