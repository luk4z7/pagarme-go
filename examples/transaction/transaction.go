// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction"
	"net/url"
	"os"
)

var transactionRecord transaction.Transaction

func main() {
	// Create
	data := []byte(`{
		"amount": 100,
		"card_id": "card_cisb66pu000po4x6ez7hbim8s",
		"customer": {
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
				"zipcode":"47850000",
				"country":"Lordaeron"
			},
			"phone": {
				"ddi":"88",
				"ddd":"88",
				"number":"808080808"
			}
		},
		"postback_url": "http://requestb.in/pkt7pgpk",
		"metadata": {
			"idProduto": "10"
		}
	}`)
	create, err, errorsApi := transactionRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}

	// Retornando o saldo de um recebedor
	get, err, errorsApi := transactionRecord.Get(url.Values{"id": {"700636"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}

	// Retornando todas as transações com paginação
	getall, err, errorsApi := transactionRecord.GetAll(url.Values{}, auth.Headers{
		"page":  "1",
		"count": "10",
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAll, _ := json.MarshalIndent(getall, "", " ")
		os.Stdout.Write(responseGetAll)
	}
}
