// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package splitrule

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
	"time"
)

var repositorySplitRule repository.Repository

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

type SplitRule struct {
	Object              string    `json:"object"`
	Id                  int       `json:"id"`
	RecipientId         string    `json:"recipient_id"`
	ChargeProcessingFee bool      `json:"charge_processing_fee"`
	Liable              bool      `json:"liable"`
	Percentage          int       `json:"percentage"`
	Amount              int       `json:"amount"`
	DateCreated         time.Time `json:"date_created"`
	DateUpdated         time.Time `json:"date_updated"`
	Route               string    `json:"route"`
}

func (s *SplitRule) Get(p url.Values, h auth.Headers) (SplitRule, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/split_rules/" + p.Get("id")
	_, err, errApi := repositorySplitRule.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *SplitRule) GetAll(p url.Values, h auth.Headers) ([]SplitRule, error, liberr.ErrorsAPI) {
	res := []SplitRule{}
	route := endPoint + "/" + p.Get("transaction_id") + "/split_rules"
	_, err, errAPi := repositorySplitRule.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errAPi
}
