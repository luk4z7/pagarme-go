// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/balance/operation"
	"net/url"
	"os"
)

var balanceOperation operation.BalanceOperation

func main() {
	// Get a Balance Operation
	get, err, errorsApi := balanceOperation.Get(url.Values{"id": {"0"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}

	// Get all Balance Operation
	getall, err, errorsApi := balanceOperation.GetAll(url.Values{}, auth.Headers{
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
