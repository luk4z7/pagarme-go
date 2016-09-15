// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"bytes"
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/config"
	liberr "github.com/luk4z7/pagarme-go/error"
	"io/ioutil"
	"net/http"
	"net/url"
)

var key config.ApiKey

var errApi = liberr.ErrorsAPI{}

// Repository is a object type for abstract methods for API
// Get, Create, Getall, etc ...
type Repository struct {
}

// params expected data for route settings and access should be as follows:
// 	_, err := repositoryCard.Get(url.Values{"route": {route}}, struct, headers)
//
// structure expected by default a struct for unmarshall data eg:
// type Card struct {
// 	Brand       string    `json:"brand"`
// 	Country     string    `json:"country"`
//
// headers expected data type for the header configuration http headers, eg:
// 	Accept:        application/json
//   	Content-Type:  application/json
// 	Authorization: Basic
//
func (r *Repository) Get(p url.Values, i interface{}, h auth.Headers) (interface{}, error, liberr.ErrorsAPI) {
	// mount the route
	resp := auth.Init(string(p.Get("route")), h)
	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	liberr.Check(err, "Cannot read this resource - method: Get")

	if resp.StatusCode == 400 {
		err, errApi := checkError([]byte(body))
		if err != nil {
			return i, err, errApi
		}
	}
	err = json.Unmarshal([]byte(body), &i)
	liberr.Check(err, "Cannot unmarshal data - method: Get")

	return i, nil, errApi
}

// params and structure also works to the method described above
// data expected []byte for create/update any services
func (r *Repository) Create(p url.Values, d []byte, i interface{}) (interface{}, error, liberr.ErrorsAPI) {
	req, err, errApi := request("POST", p.Get("route"), d, i)
	return req, err, errApi
}

func (s *Repository) Update(p url.Values, d []byte, i interface{}) (interface{}, error, liberr.ErrorsAPI) {
	req, err, errApi := request("PUT", p.Get("route"), d, i)
	return req, err, errApi
}

// Create a new connection with the endPoint last and return the response of the server
func request(m, p string, d []byte, i interface{}) (interface{}, error, liberr.ErrorsAPI) {
	req, err := http.NewRequest(m, p, bytes.NewBuffer(d))
	// Format the header
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+auth.BasicAuth(key.GetApiKey(), "x"))

	client := &http.Client{}
	resp, err := client.Do(req)
	liberr.Check(err, "Return error HTTP response of Do - method: request")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	liberr.Check(err, "Cannot read this resource - method: request")

	if resp.StatusCode == 400 {
		err, errApi := checkError([]byte(body))
		if err != nil {
			return i, err, errApi
		}
	}
	err = json.Unmarshal([]byte(body), &i)
	liberr.Check(err, "Cannot unmarshal data - method: request")

	return i, nil, errApi
}

// Check the error of API return and return the json and liberr.Err
func checkError(d []byte) (error, liberr.ErrorsAPI) {
	err := json.Unmarshal(d, &errApi)
	liberr.Check(err, "Cannot unmarshal data - method: checkError")

	if len(errApi.Errors) == 1 {
		return &liberr.Err{Name: "errosapi"}, errApi
	}
	return nil, errApi
}
