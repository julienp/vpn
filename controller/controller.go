package controller

import (
	"os/exec"

	"github.com/julienp/vpn/status"
)

type Controller struct {
	command   string
	extraArgs []string
	Status    status.VPNStatus
}

func (e *Controller) RefreshStatus() error {
	args := append(e.extraArgs, "status")
	cmd := exec.Command(e.command, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	e.Status = status.ParseStatus(string(stdoutStderr))
	return nil
}

func NewController() *Controller {
	return &Controller{
		command:   "sh",
		extraArgs: []string{"-c", "sleep 2 && echo 'Connected to lala'"},
	}
}
