// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/customer"
)

var customerRecord customer.Customer

func main() {
	data := []byte(`{
		"name":"Name",
		"email":"example@example.com",
		"document_number":"80802694594",
		"gender":"M",
		"born_at":"09-22-2015",
		"address": {
			"street":"Rua de exemplo",
			"street_number":"808",
			"neighborhood":"Bairro de exemplo",
			"complementary":"Apartamento 8",
			"city":"Cidade",
			"state":"Lordaeron",
			"zipcode":"80808080",
			"country":"Lordaeron"
		},
		"phone": {
			"ddi":"88",
			"ddd":"88",
			"number":"808080808"
		}
	}`)
	create, err, errorsApi := customerRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}


	get, err, errorsApi := customerRecord.Get(url.Values{"id": {"90456"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}


	getall, err, errorsApi := customerRecord.GetAll(url.Values{}, auth.Headers{
		"page"  : "1",
		"count" : "10",
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAll, _ := json.MarshalIndent(getall, "", " ")
		os.Stdout.Write(responseGetAll)
	}
}