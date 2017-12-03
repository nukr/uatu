package commander

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

// Commander ...
type Commander struct {
	script string
	cmd    *exec.Cmd
}

// Run ...
func (c *Commander) Run() error {
	c.Stop()
	script := strings.Split(c.script, " ")
	c.cmd = exec.Command(script[0], script[1:]...)
	c.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "c.cmd.StdoutPipe() error")
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

func (c *Commander) Stop() error {
	if c.cmd != nil && c.cmd.Process != nil {
		pgid, err := syscall.Getpgid(c.cmd.Process.Pid)
		if err != nil {
			return err
		} else {
			syscall.Kill(-pgid, 15)
		}
	}
	return nil
}

// New ...
func New(script string) *Commander {
	return &Commander{
		script: script,
	}
}
