// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/bank"
	"net/url"
	"os"
)

var bankAccount bank.Account

func main() {
	// Create a Card
	data := []byte(`{
		"bank_code": "184",
		"agencia": "0808",
		"agencia_dv": "8",
		"conta": "08808",
		"conta_dv": "8",
		"document_number": "80802694594",
		"legal_name": "Lucas Alves"
	}`)
	create, err, errorsApi := bankAccount.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	}
	responseCreate, _ := json.MarshalIndent(create, "", " ")
	os.Stdout.Write(responseCreate)

	// Get a Card
	get, err, errorsApi := bankAccount.Get(url.Values{"id": {"15897336"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	}
	responseGet, _ := json.MarshalIndent(get, "", " ")
	os.Stdout.Write(responseGet)

	// Get a Balance Operation
	getall, err, errorsApi := bankAccount.GetAll(url.Values{}, auth.Headers{
		"page":  "1",
		"count": "10",
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	}
	responseGetAll, _ := json.MarshalIndent(getall, "", " ")
	os.Stdout.Write(responseGetAll)
}
