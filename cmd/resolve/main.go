package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/peterhellberg/publicdns"
)

func main() {
	flag.Parse()

	name := flag.Arg(0)

	if name == "" {
		fmt.Println("Usage: resolve hostname [A,AAAA,CNAME,MX,PTR,ANY]")
		os.Exit(1)
	}

	var rrType publicdns.RRType

	switch strings.ToUpper(flag.Arg(1)) {
	case "A":
		rrType = publicdns.A
	case "AAAA":
		rrType = publicdns.AAAA
	case "CNAME":
		rrType = publicdns.CNAME
	case "MX":
		rrType = publicdns.MX
	case "PTR":
		rrType = publicdns.PTR
	default:
		rrType = publicdns.ANY
	}

	res, err := publicdns.Resolve(name, rrType)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)

	enc.SetIndent("", "  ")

	enc.Encode(res)
}
