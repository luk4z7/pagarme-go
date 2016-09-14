// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/recipient"
)

var recipientRecord recipient.Recipient

func main() {
	// Criando um recebedor com uma conta bancária existente
	data := []byte(`{
		"transfer_interval":"monthly",
		"transfer_day": 8,
		"transfer_enabled": true,
		"bank_account_id": 15897336
	}`)
	create, err, errorsApi := recipientRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}


	// Criando um recebedor com antecipação automática
	data2 := []byte(`{
		"transfer_interval":"monthly",
		"transfer_day": 8,
		"transfer_enabled": true,
		"bank_account_id": 15897336,
		"automatic_anticipation_enabled": true,
		"anticipatable_volume_percentage": 88
	}`)
	create2, err, errorsApi := recipientRecord.Create(data2, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate2, _ := json.MarshalIndent(create2, "", " ")
		os.Stdout.Write(responseCreate2)
	}


	// Criando um recebedor com uma conta bancária nova
	data3 := []byte(`{
		"transfer_interval":"weekly",
		"transfer_day": 1,
		"transfer_enabled": true,
		"bank_account": {
			"bank_code":"184",
			"agencia":"8",
			"conta":"08808",
			"conta_dv":"8",
			"document_number":"80802694594",
			"legal_name":"Lucas Alves"
		}
	}`)
	create3, err, errorsApi := recipientRecord.Create(data3, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate3, _ := json.MarshalIndent(create3, "", " ")
		os.Stdout.Write(responseCreate3)
	}


	// Retornando o saldo de um recebedor
	get, err, errorsApi := recipientRecord.Get(url.Values{"id": {"re_cishgtigt012zv86e859ajse4"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}


	getall, err, errorsApi := recipientRecord.GetAll(url.Values{}, auth.Headers{
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


	// Atualizando um recebedor com uma outra conta bancária existente
	data4 := []byte(`{
		"bank_account_id": 6626431
	}`)
	update, err, errorsApi := recipientRecord.Update(data4, url.Values{"id" : {"re_ciflm3dq9008r116ds3o8afvt"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseUpdate, _ := json.MarshalIndent(update, "", " ")
		os.Stdout.Write(responseUpdate)
	}


	// Atualizando um recebedor para usar antecipação automática
	data5 := []byte(`{
		"automatic_anticipation_enabled": true,
		"anticipatable_volume_percentage": 40
	}`)
	update2, err, errorsApi := recipientRecord.Update(data5, url.Values{"id" : {"re_ciflm3dq9008r116ds3o8afvt"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseUpdate2, _ := json.MarshalIndent(update2, "", " ")
		os.Stdout.Write(responseUpdate2)
	}


	// Atualizando um recebedor com novo dia para transferência
	data6 := []byte(`{
		"transfer_day": 5,
		"transfer_interval": "weekly"
	}`)
	update3, err, errorsApi := recipientRecord.Update(data6, url.Values{"id" : {"re_ciflm3dq9008r116ds3o8afvt"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseUpdate3, _ := json.MarshalIndent(update3, "", " ")
		os.Stdout.Write(responseUpdate3)
	}
}