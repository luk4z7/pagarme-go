// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zipcode

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
)

const (
	endPoint = "https://api.pagar.me/1/zipcodes"
)

type Zipcode struct {
	Neighborhood string `json:"object"`  // Bairro
	Street       string `json:"street"`  // Rua
	City         string `json:"city"`    // Cidade
	State        string `json:"state"`   // Estado
	Zipcode      string `json:"zipcode"` // Código postal (CEP)
}

func (s *Zipcode) Get(p url.Values, h auth.Headers) (Zipcode, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}
