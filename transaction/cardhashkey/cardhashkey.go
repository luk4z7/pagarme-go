// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cardhashkey

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
	"time"
)

const (
	endPoint = "https://api.pagar.me/1/transactions/card_hash_key"
)

type CardHashKey struct {
	DateCreated time.Time `json:"date_created"`
	Id          int       `json:"id"`
	Ip          string    `json:"ip"`
	PublicKey   string    `json:"public_key"`
}

func (s *CardHashKey) Get(p url.Values, h auth.Headers) (CardHashKey, error, liberr.ErrorsAPI) {
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {endPoint}}, s, h)
	return *s, err, errApi
}
