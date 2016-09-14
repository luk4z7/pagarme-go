// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/plan"
)

var planRecord plan.Plan

func main() {
	// Criando um plano
	data := []byte(`{
		"amount": "100",
		"days": "30",
		"name": "Plano de teste de 14 dias",
		"payment_methods": [
			"credit_card"
		],
		"trial_days": "14"
	}`)
	create, err, errorsApi := planRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}


	// Criando um plano anual válido por 3 anos com boleto
	data2 := []byte(`{
		"amount":"8000",
		"days": 365,
		"name": "Plano de três anos",
		"payment_methods": [
			"boleto"
		],
		"installments": 1,
		"charges": 3
	}`)
	create2, err, errorsApi := planRecord.Create(data2, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate2, _ := json.MarshalIndent(create2, "", " ")
		os.Stdout.Write(responseCreate2)
	}


	// Criando um plano anual parcelado válido por 3 anos com boleto
	data3 := []byte(`{
		"amount":"8000",
		"days": 365,
		"name": "Plano Anual Parcelado por três anos com boleto",
		"payment_methods": [
			"boleto"
		],
		"installments": 12,
		"charges": 3
	}`)
	create3, err, errorsApi := planRecord.Create(data3, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate3, _ := json.MarshalIndent(create3, "", " ")
		os.Stdout.Write(responseCreate3)
	}


	get, err, errorsApi := planRecord.Get(url.Values{"id": {"62535"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}


	getall, err, errorsApi := planRecord.GetAll(url.Values{}, auth.Headers{
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


	// Atualizando um plano
	data4 := []byte(`{
		"name": "Plano de Teste de Teste",
		"trial_days": "8"
	}`)
	update, err, errorsApi := planRecord.Update(data4, url.Values{"id" : {"40648"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseUpdate, _ := json.MarshalIndent(update, "", " ")
		os.Stdout.Write(responseUpdate)
	}
}