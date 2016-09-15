// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/recipient"
	"net/url"
	"os"
)

var payableRecord recipient.Recipient

func main() {
	get, err, errorsApi := payableRecord.Get(url.Values{"id": {"25786"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}

	getall, err, errorsApi := payableRecord.GetAll(url.Values{}, auth.Headers{
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
