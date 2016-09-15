// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction/postback"
	"net/url"
	"os"
)

var transactionPostback postback.TransactionPostback

func main() {
	get, err, errorsApi := transactionPostback.Get(url.Values{
		"transaction_id": {"700975"},
		"id":             {"po_cissza9lk05h49r73lf4uww2z"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}

	getall, err, errorsApi := transactionPostback.GetAll(url.Values{"transaction_id": {"700975"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAll, _ := json.MarshalIndent(getall, "", " ")
		os.Stdout.Write(responseGetAll)
	}

	// Reenviar um POSTback de uma transação
	redeliver, err, errorsApi := transactionPostback.Redeliver(url.Values{
		"transaction_id": {"700975"},
		"id":             {"po_cissza9lk05h49r73lf4uww2z"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseRedeliver, _ := json.MarshalIndent(redeliver, "", " ")
		os.Stdout.Write(responseRedeliver)
	}
}
