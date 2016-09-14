// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositorySubscription repository.Repository

type SubscriptionTransaction struct {
	transaction transaction.Transaction
}

const (
	endPoint = "https://api.pagar.me/1/subscriptions"
)

func (s *SubscriptionTransaction) Get(p url.Values, h auth.Headers) ([]transaction.Transaction, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/transactions"
	res := []transaction.Transaction{}
	_, err, errApi := repositorySubscription.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}