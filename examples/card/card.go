// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/card"
	"net/url"
	"os"
)

var creditCard card.Card

func main() {

	// Create a Card with card hash
	data := []byte(`{
		"card_hash": "your_card_hash"
	}`)
	create, err, errorsApi := creditCard.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}

	// Create a Card
	data2 := []byte(`{
		"card_number": "4242424242424242",
		"card_holder_name": "Marcos Mendes Teste API Create",
		"card_expiration_date": "1018"
	}`)
	create2, err, errorsApi := creditCard.Create(data2, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate2, _ := json.MarshalIndent(create2, "", " ")
		os.Stdout.Write(responseCreate2)
	}

	// Get a Card

	get, err, errorsApi := creditCard.Get(url.Values{
		"id": {"card_cisb66pu000po4x6ez7hbim8s"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}
}
