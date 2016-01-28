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
	Comment            string
	RFC                string
	Block              string
	IP                 string
	Mask               string
	Allocated          string
	Terminates         string
	Source             string
	Destination        string
	Forwardable        string
	Global             string
	ReservedByProtocol string
}

func bytesToString(bytes []byte) string {
	intStrs := []string{}
	for _, b := range bytes {
		intStrs = append(intStrs, strconv.Itoa(int(b)))
	}
	return strings.Join(intStrs, ",")
}

var commentStrip = regexp.MustCompile(`\[\d+\]$`)

// deal with some annoyances of the format
func stripString(input string) string {
	input = commentStrip.ReplaceAllLiteralString(input, "")
	input = strings.Replace(input, `"`, "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func parseBool(input string) (bool, error) {
	switch strings.ToLower(input) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("String is not a bool")
	}
}

var dateFormat = "2006-01"

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
		// (sometimes, maybe once, there is an entry with two seperate blocks)
		for _, block := range strings.Split(stripString(r[0]), ",") {
			block = strings.Trim(block, " ")
			_, cidr, err := net.ParseCIDR(block)
			if err != nil {
				fmt.Printf("Couldn't parse CIDR from %s: %s\n", block, err)
				continue
			}
			ip, mask := bytesToString(cidr.IP), bytesToString(cidr.Mask)
			allocated, terminates := "", ""
			if r[3] != "" {
				if t, err := time.Parse(dateFormat, r[3]); err == nil {
					allocated = fmt.Sprintf("%d, %d", t.Unix(), t.UnixNano())
				}
			}
			if r[4] != "" {
				if t, err := time.Parse(dateFormat, r[4]); err == nil {
					terminates = fmt.Sprintf("%d, %d", t.Unix(), t.UnixNano())
				}
			}
			attrs := make([]string, 5)
			for i, attr := range r[5:] {
				attrs[i] = strings.ToLower(stripString(attr))
				if attrs[i] == "" || attrs[i] == "n/a" {
					attrs[i] = "false"
				}
			}
			networks = append(networks, network{
				Comment:            stripString(r[1]),
				RFC:                stripString(r[2]),
				Block:              block,
				IP:                 ip,
				Mask:               mask,
				Allocated:          allocated,
				Terminates:         terminates,
				Source:             attrs[0],
				Destination:        attrs[1],
				Forwardable:        attrs[2],
				Global:             attrs[3],
				ReservedByProtocol: attrs[4],
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
