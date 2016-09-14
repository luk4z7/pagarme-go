// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transfer

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/repository"
	"github.com/luk4z7/pagarme-go/lib/bank"
	liberr "github.com/luk4z7/pagarme-go/error"
)

var repositoryTransfer repository.Repository

const (
	endPoint = "https://api.pagar.me/1/transfers"
)

type Transfer struct {
	Object               string       `json:"object"`                 // Nome do tipo do objeto criado/modificado
	Id                   int          `json:"id"`                     // Id da transferência
	Amount               int          `json:"amount"`                 // Valor em centavos transferido
	Type                 string       `json:"type"`                   // Tipo de transferência. Valores possíveis:
									  // ted, doc e credito_em_conta

	Status               string       `json:"status"`                 // Estado da transferência. Valores possíveis:
									  // pending_transfer, transferred, failed,
				                                          //  processing e  canceled

	SourceType           string       `json:"source_type"`            // O tipo de origem da qual irá ser transferido
									  // o valor

	SourceId             string       `json:"source_id"`              // O id da origem da transferencia
	TargetType           string       `json:"target_type"`            // O tipo de destino da transferencia
	TargetId             string       `json:"target_id"`              // O id do destino da transferencia
	Fee                  int          `json:"fee"`                    // Taxa em centavos cobrada pela transferência
	FundingDate          time.Time    `json:"funding_date"`           // Data ocorrência da transferência no formato
									  // ISODate

	FundingEstimatedDate time.Time    `json:"funding_estimated_date"` // Data estimada para efetivação da transferência
									  // no formato ISODate

	TransactionId        int          `json:"transaction_id"`         // Id da transação estornada no caso de estorno
									  // de boleto

	DateCreated          time.Time    `json:"date_created"`           // Data de criação da transferência no formato
									  // ISODate

	BankAccount          bank.Account `json:"bank_account"`           // Objeto da conta bancária
}

func (s *Transfer) Create(d []byte, p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryTransfer.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Transfer) Get(p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryTransfer.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Transfer) GetAll(p url.Values, h auth.Headers) ([]Transfer, error, liberr.ErrorsAPI) {
	res := []Transfer{}
	_, err, errApi := repositoryTransfer.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}

func (s *Transfer) Update(d []byte, p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryTransfer.Update(url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}

func (s *Transfer) Cancel(p url.Values, h auth.Headers) (Transfer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/cancel"
	_, err, errApi := repositoryTransfer.Create(url.Values{"route": {route}}, []byte(`{}`), s)
	return *s, err, errApi
}