// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recipient

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/bank"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositoryRecipient repository.Repository

const (
	endPoint = "https://api.pagar.me/1/recipients"
)

type Recipient struct {
	AnticipatableVolumePercentage int          `json:"anticipatable_volume_percentage"`
	AutomaticAnticipationEnabled  bool         `json:"automatic_anticipation_enabled"`
	BankAccount 		      bank.Account `json:"bank_account"`
	DateCreated 		      time.Time    `json:"date_created"`
	DateUpdated 		      time.Time    `json:"date_updated"`
	Id 			      string       `json:"id"`
	LastTransfer 		      string       `json:"last_transfer"`
	Object			      string       `json:"object"`
	TransferDay 		      int          `json:"transfer_day"`
	TransferEnabled 	      bool         `json:"transfer_enabled"`
	TransferInnterval 	      string       `json:"transfer_interval"`
}

func (s *Recipient) Create(d []byte, p url.Values, h auth.Headers) (Recipient, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryRecipient.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Recipient) Get(p url.Values, h auth.Headers) (Recipient, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryRecipient.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Recipient) GetAll(p url.Values, h auth.Headers) ([]Recipient, error, liberr.ErrorsAPI) {
	res := []Recipient{}
	_, err, errApi := repositoryRecipient.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}

func (s *Recipient) Update(d []byte, p url.Values, h auth.Headers) (Recipient, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryRecipient.Update(url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}