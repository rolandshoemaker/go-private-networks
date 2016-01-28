package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type network struct {
	Comment string
	RFC     string
	Block   string
	IP      string
	Mask    string
}

func bytesToString(bytes []byte) string {
	intStrs := []string{}
	for _, b := range bytes {
		intStrs = append(intStrs, strconv.Itoa(int(b)))
	}
	return strings.Join(intStrs, ",")
}

var commentStrip = regexp.MustCompile(`\[\d+\]$`)

func readCSV(reader *csv.Reader) []network {
	networks := []network{}
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, r := range records {
		if i == 0 {
			continue // skip header
		}
		// record structure
		// Address Block, Name, RFC, Allocation Date, Termination Date, Source, Destination, Forwardable, Global, Reserved-by-Protocol
		for _, block := range strings.Split(r[0], ",") {
			block = commentStrip.ReplaceAllLiteralString(block, "")
			block = strings.Trim(block, " ")
			_, cidr, err := net.ParseCIDR(block)
			if err != nil {
				fmt.Printf("Couldn't parse CIDR from %s: %s\n", block, err)
				continue
			}
			ip, mask := bytesToString(cidr.IP), bytesToString(cidr.Mask)
			networks = append(networks, network{
				Comment: r[1],
				RFC:     r[2],
				Block:   block,
				IP:      ip,
				Mask:    mask,
			})
		}
	}
	return networks
}

const generatedCodeWarning = `// DO NOT EDIT BY HAND
// CODE AUTOMATICALLY GENERATED USING cmd/iana-parser.go`

func main() {
	v6File := "iana-ipv6-special-registry-1.csv"
	v4File := "iana-ipv4-special-registry-1.csv"

	f, err := os.Open(v4File)
	if err != nil {
		fmt.Println(err)
		return
	}
	v4Networks := readCSV(csv.NewReader(f))
	f, err = os.Open(v6File)
	if err != nil {
		fmt.Println(err)
		return
	}
	v6Networks := readCSV(csv.NewReader(f))

	t, err := template.ParseFiles("networks.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bytes.NewBuffer(nil)
	data := struct {
		Warning    string
		Now        time.Time
		V4Networks []network
		V6Networks []network
	}{
		Warning:    generatedCodeWarning,
		Now:        time.Now().UTC(),
		V4Networks: v4Networks,
		V6Networks: v6Networks,
	}
	err = t.Execute(buf, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf.String())
}
