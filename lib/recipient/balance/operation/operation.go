// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operation

import (
	"net/url"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/balance/operation"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositoryBalance repository.Repository

const (
	endPoint = "https://api.pagar.me/1/recipients"
)

type OperationRecipient struct {
	recipient operation.BalanceOperation
}

func (s *OperationRecipient) Get(p url.Values, h auth.Headers) (operation.BalanceOperation, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id_recipient") + "/balance/operations/" + p.Get("id_operation")
	_, err, errApi := repositoryBalance.Get(url.Values{"route": {route}}, &s.recipient, h)
	return s.recipient, err, errApi
}

func (s *OperationRecipient) GetAll(p url.Values, h auth.Headers) ([]*operation.BalanceOperation, error, liberr.ErrorsAPI) {
	res := []*operation.BalanceOperation{}
	route := endPoint + "/" + p.Get("id_recipient") + "/balance/operations"
	_, err, errApi := repositoryBalance.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}