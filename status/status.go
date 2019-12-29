package status

import (
	"strings"
)

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

type Location struct {
	Alias       string
	Country     string
	Description string
}

const NEW_VERSION_INDICATOR = "A new version is available"

func ParseLocations(output string) []Location {
	skip := 2
	if strings.HasPrefix(output, NEW_VERSION_INDICATOR) {
		skip = 3
	}
	lines := strings.Split(output, "\n")
	locations := []Location{}
	for _, line := range lines[skip:] {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			break
		}
		loc := Location{Alias: fields[0]}
		locations = append(locations, loc)
	}
	return locations
}
