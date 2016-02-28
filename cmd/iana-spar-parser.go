package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
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

var dateFormat = "2006-01"

func parseDate(date string) (output string) {
	if t, err := time.Parse(dateFormat, date); err == nil {
		output = fmt.Sprintf("%d, 0", t.Unix())
	}
	return
}

func readCSV(filename string) []network {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()
	reader := csv.NewReader(f)
	networks := []network{}
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, r := range records {
		if i == 0 && r[0] == "Address Block" {
			continue // skip header
		}
		// record structure
		// Address Block, Name, RFC, Allocation Date, Termination Date, Source, Destination, Forwardable, Global, Reserved-by-Protocol
		// (sometimes, maybe once, there is an entry with two seperate address blocks)
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
				allocated = parseDate(r[3])
			}
			if r[4] != "" {
				terminates = parseDate(r[4])
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
// CODE AUTOMATICALLY GENERATED USING cmd/iana-spar-parser.go`

func main() {
	v6Filename := flag.String("v6Filename", "iana-ipv6-special-registry-1.csv", "")
	v4Filename := flag.String("v4Filename", "iana-ipv4-special-registry-1.csv", "")
	outputFilename := flag.String("output", "networks.go", "")
	flag.Parse()

	v4Networks := readCSV(*v4Filename)
	v6Networks := readCSV(*v6Filename)

	t, err := template.ParseFiles("networks.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	o, err := os.OpenFile(*outputFilename, os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer o.Close()
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
	err = t.Execute(o, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// format output with `go fmt`
	cmd := exec.Command("gofmt", "-w", *outputFilename)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
