// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package payable

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/repository"
	liberr "github.com/luk4z7/pagarme-go/error"
)

var repositoryRecipient repository.Repository

const (
	endPoint = "https://api.pagar.me/1/payables"
)

type Payable struct {
	Object        string    `json:"object"`
	Id            int       `json:"id"`
	Status        string    `json:"status"`
	Amount        int       `json:"amount"`
	Fee           int       `json:"fee"`
	Installment   string    `json:"installment"`
	TransactionId int       `json:"transaction_id"`
	SplitRuleId   string    `json:"split_rule_id"`
	PaymentDate   time.Time `json:"payment_date"`
	Type 	      string 	`json:"type"`
	PaymentMethod string    `json:"payment_method"`
	DateCreated   time.Time `json:"date_created"`
}

func (s *Payable) Get(p url.Values, h auth.Headers) (Payable, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryRecipient.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Payable) GetAll(p url.Values, h auth.Headers) ([]Payable, error, liberr.ErrorsAPI) {
	res := []Payable{}
	_, err, errApi := repositoryRecipient.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}