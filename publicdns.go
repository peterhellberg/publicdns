package publicdns

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RRType string

const (
	A     RRType = "A"
	AAAA  RRType = "AAAA"
	CNAME RRType = "CNAME"
	MX    RRType = "MX"
	PTR   RRType = "PTR"
	ANY   RRType = "ANY"
)

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

type Question struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type Answer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

var DefaultClient = &Client{http.DefaultClient}

type Client struct {
	httpClient *http.Client
}

func (c *Client) Resolve(name string, rrType RRType) (*Response, error) {
	rawurl := fmt.Sprintf("https://dns.google.com/resolve?name=%s&type=%s", name, rrType)

	resp, err := c.httpClient.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 64)
		resp.Body.Close()
	}()

	r := &Response{}

	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}

func Resolve(name string, rrType RRType) (*Response, error) {
	return DefaultClient.Resolve(name, rrType)
}
