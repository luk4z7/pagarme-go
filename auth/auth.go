// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"encoding/base64"
	"github.com/luk4z7/pagarme-go/config"
	liberr "github.com/luk4z7/pagarme-go/error"
	"net/http"
)

var key config.ApiKey

type Headers map[string]string

// encoded username + password for base64 for http basic access authentRequestication
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Authentication with API and return Response
func Init(uri string, h Headers) *http.Response {
	// Create new request with endPoint passed
	req, err := http.NewRequest("GET", uri, nil)
	liberr.Check(err, "Return error of the new Request given a method of NewRequest - method: Init")
	// add more headers configurations
	if len(h) > 0 {
		q := req.URL.Query()
		for k, v := range h {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	// Authentication using basic http
	req.Header.Add("Authorization", "Basic "+BasicAuth(key.GetApiKey(), "x"))
	cli := &http.Client{}
	resp, err := cli.Do(req)
	liberr.Check(err, "Return error HTTP response of Do - method: Init")

	return resp
}
