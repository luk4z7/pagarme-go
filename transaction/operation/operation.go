// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operation

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/operation"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
)

type TransactionOperation struct {
	operation operation.Operation
}

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

func (s *TransactionOperation) Get(p url.Values, h auth.Headers) (operation.Operation, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/operations/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, &s.operation, h)
	return s.operation, err, errApi
}

func (s *TransactionOperation) GetAll(p url.Values, h auth.Headers) ([]operation.Operation, error, liberr.ErrorsAPI) {
	res := []operation.Operation{}
	route := endPoint + "/" + p.Get("transaction_id") + "/operations"
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}
