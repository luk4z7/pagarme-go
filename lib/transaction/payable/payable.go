// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package payable

import (
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/repository"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/payable"
)

var repositoryPayable repository.Repository

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

type TransactionPayable struct {
	payable payable.Payable
}

func (s *TransactionPayable) Get(p url.Values, h auth.Headers) (payable.Payable, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/payables/" + p.Get("id")
	_, err, errApi := repositoryPayable.Get(url.Values{"route": {route}}, &s.payable, h)
	return s.payable, err, errApi
}

func (s *TransactionPayable) GetAll(p url.Values, h auth.Headers) ([]payable.Payable, error, liberr.ErrorsAPI) {
	res := []payable.Payable{}
	route := endPoint + "/" + p.Get("transaction_id") + "/payables"
	_, err, errApi := repositoryPayable.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}