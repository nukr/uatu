package commander

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// Commander ...
type Commander struct {
	script string
	cmd    *exec.Cmd
}

// Run ...
func (c *Commander) Run() error {
	if c.cmd != nil {
		pgid, err := syscall.Getpgid(c.cmd.Process.Pid)
		if err != nil {
			return err
		}
		syscall.Kill(-pgid, 15)
	}
	script := strings.Split(c.script, " ")
	c.cmd = exec.Command(script[0], script[1:]...)
	c.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = c.cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		io.Copy(os.Stdout, stdout)
	}()
	go func() {
		io.Copy(os.Stderr, stderr)
	}()
	return nil
}

// New ...
func New(script string) *Commander {
	return &Commander{
		script: script,
	}
}
