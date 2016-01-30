package ianasparlist

import (
	"net"
	"time"
)

func filterBadDestinations(networks []Network) []net.IPNet {
	filtered := []net.IPNet{}
	now := time.Now()
	for _, network := range networks {
		if (!network.Destination || (network.Destination && !network.Global)) && (network.Terminates.IsZero() || !now.After(network.Terminates)) {
			filtered = append(filtered, network.CIDR)
		}
	}
	return filtered
}

func InvalidPublicV4Destinations() []net.IPNet {
	return filterBadDestinations(V4)
}

func InvalidPublicV6Destinations() []net.IPNet {
	return filterBadDestinations(V6)
}

func InvalidPublicDestinations() []net.IPNet {
	return append(filterBadDestinations(V4), filterBadDestinations(V6)...)
}
