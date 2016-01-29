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

func filterDestination(destination bool, networks []Network) []net.IPNet {
	filtered := []net.IPNet{}
	now := time.Now()
	for _, network := range V4 {
		if network.Destination == destination && (network.Terminates.IsZero() || !now.After(network.Terminates)) {
			filtered = append(filtered, network.CIDR)
		}
	}
	return filtered
}

func InvalidV4Destinations() []net.IPNet {
	return filterDestination(false, V4)
}

func InvalidV6Destinations() []net.IPNet {
	return filterDestination(false, V6)
}

func InvalidDestinations() []net.IPNet {
	return append(InvalidV4Destinations(), InvalidV6Destinations()...)
}

func ValidV4Destinations() []net.IPNet {
	return filterDestination(true, V4)
}

func ValidV6Destinations() []net.IPNet {
	return filterDestination(true, V6)
}

func ValidDestinations() []net.IPNet {
	return append(ValidV4Destinations(), ValidV6Destinations()...)
}
