// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transfer

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/bank"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
	"time"
)

const (
	endPoint = "https://api.pagar.me/1/transfers"
)

// object                 Nome do tipo do objeto criado/modificado
// id                     Id da transferência
// amount                 Valor em centavos transferido
// type                   Tipo de transferência. Valores possíveis: ted, doc e credito_em_conta
// status                 Estado da transferência. Valores possíveis: pending_transfer, transferred, failed, processing
// 			  e canceled
//
// source_type            O tipo de origem da qual irá ser transferido o valor
// source_id              O id da origem da transferencia
// target_type            O tipo de destino da transferencia
// target_id              O id do destino da transferencia
// fee                    Taxa em centavos cobrada pela transferência
// funding_date           Data ocorrência da transferência no formato ISODate
// funding_estimated_date Data estimada para efetivação da transferência no formato ISODate
// transaction_id         Id da transação estornada no caso de estorno de boleto
// date_created           Data de criação da transferência no formato ISODate
// bank_account           Objeto da conta bancária
type Transfer struct {
	Object               string       `json:"object"`
	Id                   int          `json:"id"`
	Amount               int          `json:"amount"`
	Type                 string       `json:"type"`
	Status               string       `json:"status"`
	SourceType           string       `json:"source_type"`
	SourceId             string       `json:"source_id"`
	TargetType           string       `json:"target_type"`
	TargetId             string       `json:"target_id"`
	Fee                  int          `json:"fee"`
	FundingDate          time.Time    `json:"funding_date"`
	FundingEstimatedDate time.Time    `json:"funding_estimated_date"`
	TransactionId        int          `json:"transaction_id"`
	DateCreated          time.Time    `json:"date_created"`
	BankAccount          bank.Account `json:"bank_account"`
}

func (s *Transfer) Create(d []byte, p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	req := request.Client{}
	_, err, errApi := req.New("POST", url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Transfer) Get(p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Transfer) GetAll(p url.Values, h auth.Headers) ([]Transfer, error, liberr.ErrorsAPI) {
	res := []Transfer{}
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}

func (s *Transfer) Update(d []byte, p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.New("UPDATE", url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}

func (s *Transfer) Cancel(p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/cancel"
	req := request.Client{}
	_, err, errApi := req.New("POST", url.Values{"route": {route}}, []byte(`{}`), s)
	return *s, err, errApi
}
