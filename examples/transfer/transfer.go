// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transfer"
	"net/url"
	"os"
)

var transferRecord transfer.Transfer

func main() {
	// Criando uma transferência a partir de um recebedor específico
	data := []byte(`{
		"amount":"100",
		"bank_account_id": "15897336",
		"recipient_id": "re_cisv78zak000rbp6dx6rvyp6v"
	}`)
	create, err, errorsApi := transferRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", "   ")
		os.Stdout.Write(responseCreate)
	}

	// Criando uma transferência a partir do recebedor padrão
	data2 := []byte(`{
		"amount":"13000",
		"source_id": "re_cisv78zak000rbp6dx6rvyp6v",
		"target_id": "re_cisv78zak000rbp6dx6rvyp6v"
	}`)
	create2, err, errorsApi := transferRecord.Create(data2, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate2, _ := json.MarshalIndent(create2, "", " ")
		os.Stdout.Write(responseCreate2)
	}

	// Retornando uma transferência
	get, err, errorsApi := transferRecord.Get(url.Values{"id": {"7118"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", "   ")
		os.Stdout.Write(responseGet)
	}

	// Retornando todas as transferências
	getall, err, errorsApi := transferRecord.GetAll(url.Values{}, auth.Headers{
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

	// Cancelando uma transferência
	cancel, err, errorsApi := transferRecord.Cancel(url.Values{
		"id": {"7118"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCancel, _ := json.MarshalIndent(cancel, "", " ")
		os.Stdout.Write(responseCancel)
	}
}
