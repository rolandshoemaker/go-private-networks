package privateNetworks

import (
	"net"
	"time"
)

type Network struct {
	CIDR net.IPNet

	Comment string
	RFC     string

	Allocated  time.Time
	Terminates time.Time

	Source             bool
	Destination        bool
	Forwardable        bool
	Global             bool
	ReservedByProtocol bool
}

func InvalidV4Destinations() []net.IPNet {
	filtered := []net.IPNet{}
	for _, network := range V4 {
		if !network.Destination {
			filtered = append(filtered, network.CIDR)
		}
	}
	return filtered
}

func InvalidV6Destinations() []net.IPNet {
	filtered := []net.IPNet{}
	for _, network := range V6 {
		if !network.Destination {
			filtered = append(filtered, network.CIDR)
		}
	}
	return filtered
}
