// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/recipient/balance/operation"
	"net/url"
	"os"
)

var balanceRecipient operation.OperationRecipient

func main() {
	// Get a Balance Amount
	get, err, errorsApi := balanceRecipient.Get(url.Values{
		"id_recipient": {"re_ciflm3dq9008r116ds3o8afvt"},
		"id_operation": {"10"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}

	getAll, err, errorsApi := balanceRecipient.GetAll(url.Values{
		"id_recipient": {"re_ciflm3dq9008r116ds3o8afvt"},
	}, auth.Headers{
		"page":  "1",
		"count": "10",
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAll, _ := json.MarshalIndent(getAll, "", " ")
		os.Stdout.Write(responseGetAll)
	}
}
