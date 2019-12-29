package controller

import (
	"fmt"
	"os/exec"

	"github.com/julienp/vpn/status"
)

type Controller struct {
	command   string
	extraArgs []string
	Status    status.VPNStatus
	Location  *status.Location
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

func (e *Controller) ListLocations() ([]status.Location, error) {
	args := append(e.extraArgs, "list", "all")
	cmd := exec.Command(e.command, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return []status.Location{}, err
	}
	locations := status.ParseLocations(string(stdoutStderr))
	return locations, nil
}

func (e *Controller) SetLocation(location string) error {
	return fmt.Errorf("Not implemented")
}

func NewController() *Controller {
	return &Controller{
		command: "expressvpn",
	}
}
