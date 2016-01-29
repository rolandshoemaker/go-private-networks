package specialNetworks

import (
	"net"
	"time"
)

// Network describes a special v4/6 network
// according to the IANA Special-Purpose Address
// Registries
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

// InvalidV4Destinations returns a slice of private IPv4 networks
// that shouldn't be used as destinations
func InvalidV4Destinations() []net.IPNet {
	return filterDestination(false, V4)
}

// InvalidV6Destinations returns a slice of private IPv6 networks
// that shouldn't be used as destinations
func InvalidV6Destinations() []net.IPNet {
	return filterDestination(false, V6)
}

// InvalidDestinations returns a slice of both private IPv4 and IPv6
// networks that shouldn't be used as destinations
func InvalidDestinations() []net.IPNet {
	return append(InvalidV4Destinations(), InvalidV6Destinations()...)
}

// ValidV4Destinations returns a slice of private IPv4 networks
// that can be used as destinations
func ValidV4Destinations() []net.IPNet {
	return filterDestination(true, V4)
}

// ValidV6Destinations returns a slice of private IPv6 networks
// that can be used as destinations
func ValidV6Destinations() []net.IPNet {
	return filterDestination(true, V6)
}

// ValidDestinations returns a slice of both private IPv4 and IPv6
// networks that can be used as destinations
func ValidDestinations() []net.IPNet {
	return append(ValidV4Destinations(), ValidV6Destinations()...)
}
