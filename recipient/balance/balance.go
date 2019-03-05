// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/balance"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
)

type BalanceRecipient struct {
	recipient balance.BalanceAmount
}

const (
	endPoint = "https://api.pagar.me/1/recipients"
)

func (s *BalanceRecipient) Get(p url.Values, h auth.Headers) (balance.BalanceAmount, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/balance"
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, &s.recipient, h)
	return s.recipient, err, errApi
}
