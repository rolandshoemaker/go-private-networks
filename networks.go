// DO NOT EDIT BY HAND
// CODE AUTOMATICALLY GENERATED USING cmd/iana-parser.go
// Generated 2016-01-28 23:24:52.766585666 +0000 UTC

package privateNetworks

import (
	"net"
)

var (
  V4 = []Network{
    {
      CIDR: net.IPNet{
        IP: []byte{ 0,0,0,0 },
        Mask: []byte{ 255,0,0,0 },
      },
      Comment: "This host on this network",
      RFC: "[RFC1122], section 3.2.1.3",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 10,0,0,0 },
        Mask: []byte{ 255,0,0,0 },
      },
      Comment: "Private-Use",
      RFC: "[RFC1918]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 100,64,0,0 },
        Mask: []byte{ 255,192,0,0 },
      },
      Comment: "Shared Address Space",
      RFC: "[RFC6598]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 127,0,0,0 },
        Mask: []byte{ 255,0,0,0 },
      },
      Comment: "Loopback",
      RFC: "[RFC1122], section 3.2.1.3",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 169,254,0,0 },
        Mask: []byte{ 255,255,0,0 },
      },
      Comment: "Link Local",
      RFC: "[RFC3927]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 172,16,0,0 },
        Mask: []byte{ 255,240,0,0 },
      },
      Comment: "Private-Use",
      RFC: "[RFC1918]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,0,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "IETF Protocol Assignments",
      RFC: "[RFC6890], section 2.1",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,0,0 },
        Mask: []byte{ 255,255,255,248 },
      },
      Comment: "IPv4 Service Continuity Prefix",
      RFC: "[RFC7335]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,0,8 },
        Mask: []byte{ 255,255,255,255 },
      },
      Comment: "IPv4 dummy address",
      RFC: "[RFC7600]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,0,9 },
        Mask: []byte{ 255,255,255,255 },
      },
      Comment: "Port Control Protocol Anycast",
      RFC: "[RFC7723]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,0,170 },
        Mask: []byte{ 255,255,255,255 },
      },
      Comment: "NAT64/DNS64 Discovery",
      RFC: "[RFC7050], section 2.2",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,0,171 },
        Mask: []byte{ 255,255,255,255 },
      },
      Comment: "NAT64/DNS64 Discovery",
      RFC: "[RFC7050], section 2.2",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,0,2,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "Documentation (TEST-NET-1)",
      RFC: "[RFC5737]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,31,196,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "AS112-v4",
      RFC: "[RFC7535]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,52,193,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "AMT",
      RFC: "[RFC7450]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,88,99,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "Deprecated (6to4 Relay Anycast)",
      RFC: "[RFC7526]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,168,0,0 },
        Mask: []byte{ 255,255,0,0 },
      },
      Comment: "Private-Use",
      RFC: "[RFC1918]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 192,175,48,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "Direct Delegation AS112 Service",
      RFC: "[RFC7534]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 198,18,0,0 },
        Mask: []byte{ 255,254,0,0 },
      },
      Comment: "Benchmarking",
      RFC: "[RFC2544]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 198,51,100,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "Documentation (TEST-NET-2)",
      RFC: "[RFC5737]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 203,0,113,0 },
        Mask: []byte{ 255,255,255,0 },
      },
      Comment: "Documentation (TEST-NET-3)",
      RFC: "[RFC5737]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 240,0,0,0 },
        Mask: []byte{ 240,0,0,0 },
      },
      Comment: "Reserved",
      RFC: "[RFC1112], section 4",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 255,255,255,255 },
        Mask: []byte{ 255,255,255,255 },
      },
      Comment: "Limited Broadcast",
      RFC: "[RFC919], section 7",
    },

  }
  V6 = []net.IPNet{
    {
      CIDR: net.IPNet{
        IP: []byte{ 0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1 },
        Mask: []byte{ 255,255,255,255,255,255,255,255,255,255,255,255,255,255,255,255 },
      },
      Comment: "Loopback Address",
      RFC: "[RFC4291]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,255,255,255,255,255,255,255,255,255,255 },
      },
      Comment: "Unspecified Address",
      RFC: "[RFC4291]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 0,0,0,0,0,0,0,0,0,0,255,255,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,255,255,255,255,255,255,0,0,0,0 },
      },
      Comment: "IPv4-mapped Address",
      RFC: "[RFC4291]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 0,100,255,155,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,255,255,255,255,255,255,0,0,0,0 },
      },
      Comment: "IPv4-IPv6 Translat.",
      RFC: "[RFC6052]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,255,255,0,0,0,0,0,0,0,0 },
      },
      Comment: "Discard-Only Address Block",
      RFC: "[RFC6666]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,254,0,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "IETF Protocol Assignments",
      RFC: "[RFC2928]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "TEREDO",
      RFC: "[RFC4380]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,1,0,0,0,0,0,0,0,0,0,0,0,1 },
        Mask: []byte{ 255,255,255,255,255,255,255,255,255,255,255,255,255,255,255,255 },
      },
      Comment: "Port Control Protocol Anycast",
      RFC: "[RFC7723]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,2,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "Benchmarking",
      RFC: "[RFC5180]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,3,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "AMT",
      RFC: "[RFC7450]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,4,1,18,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "AS112-v6",
      RFC: "[RFC7535]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,5,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "EID Space for LISP (TEMPORARY - registered 2015-10-22, expires 
      2016-10-22)",
      RFC: "[draft-ietf-lisp-eid-block]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,16,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,240,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "Deprecated (previously ORCHID)",
      RFC: "[RFC4843]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,0,32,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,240,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "ORCHIDv2",
      RFC: "[RFC7343]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,1,13,184,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "Documentation",
      RFC: "[RFC3849]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 32,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "6to4",
      RFC: "[RFC3056]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 38,32,0,79,128,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,255,255,255,255,255,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "Direct Delegation AS112 Service",
      RFC: "[RFC7534]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 252,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 254,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "Unique-Local",
      RFC: "[RFC4193]",
    },
    {
      CIDR: net.IPNet{
        IP: []byte{ 254,128,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
        Mask: []byte{ 255,192,0,0,0,0,0,0,0,0,0,0,0,0,0,0 },
      },
      Comment: "Linked-Scoped Unicast",
      RFC: "[RFC4291]",
    },

  }
)

