// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/transaction"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
)

type SubscriptionTransaction struct {
	transaction transaction.Transaction
}

const (
	endPoint = "https://api.pagar.me/1/subscriptions"
)

func (s *SubscriptionTransaction) Get(p url.Values, h auth.Headers) ([]transaction.Transaction, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/transactions"
	res := []transaction.Transaction{}
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}
