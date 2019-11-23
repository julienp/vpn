package status

import "strings"

type VPNStatus int

const (
	Unknown VPNStatus = iota
	Disconnected
	Connecting
	Connected
)

func (s VPNStatus) String() string {
	switch s {
	case Connected:
		return "Connected"
	case Connecting:
		return "Connecting"
	case Disconnected:
		return "Disconnected"
	default:
		return "Unknown"
	}
}

func ParseStatus(output string) VPNStatus {
	if strings.HasPrefix(output, "Connected to") {
		return Connected
	} else if strings.HasPrefix(output, "Not connected") {
		return Disconnected
	} else if strings.HasPrefix(output, "Connecting") {
		return Connecting
	}
	return Unknown
}
