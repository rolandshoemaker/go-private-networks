package privateNetworks

import (
	"net"
	"time"
)

type Network struct {
	CIDR net.IPNet

	Comment string
	RFC     string

	Allocated  *time.Time
	Terminates *time.Time

	Source             bool
	Destination        bool
	Forwardable        bool
	Global             bool
	ReservedbyProtocol bool
}
