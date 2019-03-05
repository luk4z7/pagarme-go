// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operation

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
	"time"
)

const (
	endPoint = "https://api.pagar.me/1/balance/operations"
)

// amount          Valor em centavos transacionado para a conta
// date_created    Data de criação da operação de saldo no formato ISODate
// fee             Valor em centavos que foi cobrado pela transação (taxa)
// id              Id da operação de saldo
// movement_object Objeto da origem da operação de saldo
// object          Nome do tipo do objeto criado/modificado
// status          Estado do saldo da conta. Valores possíveis: waiting_funds, available e transferred
// type            Tipo de objeto que gerou a operação de saldo
type BalanceOperation struct {
	Amount          int             `json:"amount"`
	DateCreated     time.Time       `json:"date_created"`
	Fee             int             `json:"fee"`
	Id              int             `json:"id"`
	BalanceMovement BalanceMovement `json:"movement_object"`
	Object          string          `json:"object"`
	Status          string          `json:"status"`
	Type            string          `json:"type"`
}

// amount                Valor em centados do que foi pago
// anticipation_fee      Valor em centavos que foi cobrado pela antecipação (taxa)
// bulk_anticipation_id
// date_created          Data de criação da operação no formato ISODate
// fee                   Valor em centavos que foi cobrado pela transação (taxa)
// id                    Id do recebível
// installment           Número da parcela
// object                Nome do tipo do objeto criado/modificado
// original_payment_date Dia e hora da primeira data de pagamento no formato ISODate
// payment_date          Dia e hora do pagamento no formato ISODate
// payment_method        Forma de pagamento usada
// recipient_id          Id do recebedor ao qual esse recebível pertence
// split_rule_id         Id da regra de split, se houver alguma
// status                Estado atual do recebível. Valores possíveis: waiting_funds, paid e suspended
// transaction_id        Id da transação desse recebível
// type                  Tipo do recebível. Valores possíveis: credit, refund, chargeback e chargeback_refund
type BalanceMovement struct {
	Amount              int       `json:"amount"`
	AnticipationFee     int       `json:"anticipation_fee"`
	BulkAnticipationId  int       `json:"bulk_anticipation_id"`
	DateCreated         time.Time `json:"date_created"`
	Fee                 int       `json:"fee"`
	Id                  int       `json:"id"`
	Installment         int       `json:"installment"`
	Object              string    `json:"object"`
	OriginalPaymentDate string    `json:"original_payment_date"`
	PaymentDate         time.Time `json:"payment_date"`
	PaymentMethod       string    `json:"payment_method"`
	RecipientId         string    `json:"recipient_id"`
	SplitRuleId         string    `json:"split_rule_id"`
	Status              string    `json:"status"`
	TransactionId       int       `json:"transaction_id"`
	Type                string    `json:"type"`
}

func (s *BalanceOperation) Get(p url.Values, h auth.Headers) (BalanceOperation, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *BalanceOperation) GetAll(p url.Values, h auth.Headers) ([]*BalanceOperation, error, liberr.ErrorsAPI) {
	res := []*BalanceOperation{}
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {endPoint}}, res, h)
	return res, err, errApi
}
