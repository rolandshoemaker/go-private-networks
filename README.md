# Golang IANA SPAR IPv4/6 Lists

This library provides two lists of IPv4/6 networks described by/generated from
the IANA Special Purpose Address Registry lists. These lists contain the defined
`net.IPNet`s as well as other metadeta provided by IANA.

The library also provides a number of basic methods to filter the lists and return
only `net.IPNet`s with certain properties (i.e. invalid public destinations).

In order to generate/update `networks.go` you need CSV copies of the SPAR lists
(v4 available [here](https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry-1.csv) and v6 [here](https://www.iana.org/assignments/iana-ipv6-special-registry/iana-ipv6-special-registry-1.csv)), the
parser/generator can then be run using

```
$ go run cmd/iana-spar-parser.go
```
