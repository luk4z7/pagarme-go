// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package capture

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/transaction"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
)

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

type Capture struct {
	capture transaction.Transaction
}

func (s *Capture) Create(d []byte, p url.Values, h auth.Headers) (transaction.Transaction, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/capture"
	req := request.Client{}
	_, err, errApi := req.New("POST", url.Values{"route": {route}}, d, &s.capture)
	return s.capture, err, errApi
}
