// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction/collectpayment"
	"net/url"
	"os"
)

var transactionCollectPayment collectpayment.CollectPayment

func main() {
	data := []byte(`{"email": "seu@email.com"}`)
	create, err, errorsApi := transactionCollectPayment.Create(data, url.Values{"transaction_id": {"700975"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}
}
