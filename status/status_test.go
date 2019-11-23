package status

import (
	"testing"
)

func TestParseStatus(t *testing.T) {
	cases := []struct {
		output   string
		expected VPNStatus
	}{
		{
			"",
			Unknown,
		},
		{
			"",
			Unknown,
		},
		{
			"Connecting ...",
			Connecting,
		},
		{
			"Not connected",
			Disconnected,
		},
		{
			`Connected to Japan - Tokyo - 1
			- If your VPN connection unexpectedly drops, internet traffic will be blocked to protect your privacy.
			- To disable Network Lock, disconnect ExpressVPN then type 'expressvpn preferences set network_lock off'.`,
			Connected,
		},
	}
	for _, test := range cases {
		got := ParseStatus(test.output)
		if test.expected != got {
			t.Errorf("Expected %q, got %q", test.expected, got)
		}
	}
}
