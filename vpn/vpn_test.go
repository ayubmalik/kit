package vpn_test

import (
	"github.com/ayubmalik/kit/vpn"
	"net"
	"testing"
)

func TestIsPointToPoint(t *testing.T) {
	testCases := []struct {
		name       string
		interfaces []net.Interface
		expected   bool
	}{
		{
			name:       "Empty interfaces",
			interfaces: []net.Interface{},
			expected:   false,
		},
		{
			name: "No Point-to-Point interfaces",
			interfaces: []net.Interface{
				{Name: "eth0", Flags: net.FlagUp | net.FlagRunning},
				{Name: "wlan0", Flags: net.FlagUp | net.FlagRunning},
			},
			expected: false,
		},
		{
			name: "Point-to-Point interface exists but not up",
			interfaces: []net.Interface{
				{Name: "eth0", Flags: net.FlagUp | net.FlagRunning},
				{Name: "ppp0", Flags: net.FlagPointToPoint},
			},
			expected: false,
		},
		{
			name: "Point-to-Point interface exists, is up but not running",
			interfaces: []net.Interface{
				{Name: "eth0", Flags: net.FlagUp | net.FlagRunning},
				{Name: "ppp0", Flags: net.FlagPointToPoint | net.FlagUp},
			},
			expected: false,
		},
		{
			name: "Point-to-Point interface exists and is up and running",
			interfaces: []net.Interface{
				{Name: "eth0", Flags: net.FlagUp | net.FlagRunning},
				{Name: "ppp0", Flags: net.FlagPointToPoint | net.FlagUp | net.FlagRunning},
			},
			expected: true,
		},
		{
			name: "Multiple Point-to-Point interfaces with one up and running",
			interfaces: []net.Interface{
				{Name: "ppp0", Flags: net.FlagPointToPoint},
				{Name: "ppp1", Flags: net.FlagPointToPoint | net.FlagUp | net.FlagRunning},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := vpn.IsPointToPoint(tc.interfaces)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for test case: %s", tc.expected, result, tc.name)
			}
		})
	}
}
