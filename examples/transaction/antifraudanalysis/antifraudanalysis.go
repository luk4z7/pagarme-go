// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction/antifraudanalysis"
	"net/url"
	"os"
)

var transactionAntifraudAnalysis antifraudanalysis.AntifraudAnalysis

func main() {
	get, err, errorsApi := transactionAntifraudAnalysis.Get(url.Values{
		"transaction_id": {"314578"},
		"id":             {"189164"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}

	getall, err, errorsApi := transactionAntifraudAnalysis.GetAll(url.Values{"transaction_id": {"314578"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAll, _ := json.MarshalIndent(getall, "", " ")
		os.Stdout.Write(responseGetAll)
	}
}
