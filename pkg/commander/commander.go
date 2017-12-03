package commander

import (
	"io"
	"os"
	"os/exec"
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
	shell := os.Getenv("SHELL")
	c.cmd = exec.Command(shell, "-c", c.script)
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

// Stop kills all process we created
func (c *Commander) Stop() error {
	if c.cmd != nil && c.cmd.Process != nil {
		pgid, err := syscall.Getpgid(c.cmd.Process.Pid)
		if err != nil {
			return err
		}
		syscall.Kill(-pgid, 15)
	}
	return nil
}

// New ...
func New(script string) *Commander {
	return &Commander{
		script: script,
	}
}
