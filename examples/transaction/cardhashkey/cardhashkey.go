// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction/cardhashkey"
	"net/url"
	"os"
)

var transactionCardHashKey cardhashkey.CardHashKey

func main() {
	get, err, errorsApi := transactionCardHashKey.Get(url.Values{}, auth.Headers{
		"encryption_key": "ek_test_XZUwx9tUMr0l2GL1H6XjIsnc8X5nYg",
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}
}
