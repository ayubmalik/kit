package vpn

import (
	"net"
)

// IsTunnelConnected returns true is there is running interface of type point to point
func IsTunnelConnected() bool {
	interfaces, err := net.Interfaces()
	if err != nil {
		return false
	}

	return IsPointToPoint(interfaces)
}

func IsPointToPoint(interfaces []net.Interface) bool {
	for _, i := range interfaces {
		if (i.Flags&net.FlagPointToPoint != 0) && (i.Flags&net.FlagUp != 0) && (i.Flags&net.FlagRunning != 0) {
			return true
		}
	}
	return false
}
