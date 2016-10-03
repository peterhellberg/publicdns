/*

Package publicdns is a client for Google Public DNS

Installation

Just go get the package:

go get -u github.com/peterhellberg/publicdns

Usage

    package main

    import (
    	"encoding/json"
    	"fmt"
    	"os"

    	"github.com/peterhellberg/publicdns"
    )

    func main() {
    	res, err := publicdns.Resolve("example.com", publicdns.ANY)
    	if err != nil {
    		fmt.Printf("Error: %v\n", err)
    		os.Exit(1)
    	}

    	enc := json.NewEncoder(os.Stdout)

    	enc.SetIndent("", "  ")

    	enc.Encode(res)
    }

*/
package publicdns

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// RRType is the resource record type
type RRType string

const (
	// A is the address record
	A RRType = "A"

	// AAAA is the IPv6 address record
	AAAA RRType = "AAAA"

	// CNAME is the canonical name record
	CNAME RRType = "CNAME"

	// MX is the mail exchange record
	MX RRType = "MX"

	// PTR is the pointer record
	PTR RRType = "PTR"

	// ANY is a wildcard record in Google Public DNS
	ANY RRType = "ANY"
)

// Response returned from Google Public DNS
type Response struct {
	Status    int        `json:"Status"`
	TC        bool       `json:"TC"`
	RD        bool       `json:"RD"`
	RA        bool       `json:"RA"`
	AD        bool       `json:"AD"`
	CD        bool       `json:"CD"`
	Questions []Question `json:"Question"`
	Answers   []Answer   `json:"Answer"`
	Comment   string     `json:"Comment"`
}

// Question is returned in the response
type Question struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

// Answer is returned in the response
type Answer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

// DefaultClient is the default client
var DefaultClient = &Client{http.DefaultClient}

// Client is the Google Public DNS client
type Client struct {
	httpClient *http.Client
}

// Resolve resolves the name and RR type
func (c *Client) Resolve(name string, rrType RRType) (*Response, error) {
	resp, err := c.httpClient.Get(rawurl(name, rrType))
	if err != nil {
		return nil, err
	}
	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 64)
		resp.Body.Close()
	}()

	r := &Response{Questions: []Question{}, Answers: []Answer{}}

	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}

// Resolve resolves the name and RR type
func Resolve(name string, rrType RRType) (*Response, error) {
	return DefaultClient.Resolve(name, rrType)
}

func rawurl(name string, rrType RRType) string {
	return fmt.Sprintf("https://dns.google.com/resolve?name=%s&type=%s", name, rrType)
}
