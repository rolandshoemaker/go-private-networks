package ianasparlist

import (
	"time"
)

func filterBadDestinations(networks []Network) []Network {
	filtered := []Network{}
	now := time.Now()
	for _, network := range networks {
		if network.Comment != "IPv4-mapped Address" && (!network.Destination || (network.Destination && !network.Global)) && (network.Terminates.IsZero() || !now.After(network.Terminates)) {
			filtered = append(filtered, network)
		}
	}
	return filtered
}

func InvalidPublicV4Destinations() []Network {
	return filterBadDestinations(V4)
}

func InvalidPublicV6Destinations() []Network {
	return filterBadDestinations(V6)
}

func InvalidPublicDestinations() []Network {
	return append(filterBadDestinations(V4), filterBadDestinations(V6)...)
}
