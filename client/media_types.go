//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/Tamrin007/cellar/design
// --out=$(GOPATH)/src/github.com/Tamrin007/cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "net/http"

// A bottle of wine (default view)
//
// Identifier: application/vnd.goa.example.bottle+json; view=default
type GoaExampleBottle struct {
	// API href for making requests on the bottle
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// Unique bottle ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of wine
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// DecodeGoaExampleBottle decodes the GoaExampleBottle instance encoded in resp body.
func (c *Client) DecodeGoaExampleBottle(resp *http.Response) (*GoaExampleBottle, error) {
	var decoded GoaExampleBottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
