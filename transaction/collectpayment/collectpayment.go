// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collectpayment

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
)

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

type CollectPayment struct{}

func (s *CollectPayment) Create(d []byte, p url.Values, h auth.Headers) (CollectPayment, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/collect_payment"
	req := request.Client{}
	_, err, errApi := req.Create(url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}
