// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/card"
	"github.com/luk4z7/pagarme-go/lib/customer"
	"github.com/luk4z7/pagarme-go/lib/plan"
	"github.com/luk4z7/pagarme-go/lib/transaction"
	"github.com/luk4z7/pagarme-go/request"
	"net/url"
	"time"
)

const (
	endPoint = "https://api.pagar.me/1/subscriptions"
)

// object               Nome do tipo de objeto criado ou modificado. Valor retornado: subscription
// plan                 Objeto com os dados do plano que a assinatura está associada
// id                   Número do identificador do plano
// current_transaction  Objeto com os dados da ultima transação realizada pela assinatura
// postback_url         Endpoint da aplicação integrada ao Pagar.me que irá receber os jsons de resposta a
// 			cada atualização dos processos
//
// payment_method       Método de pagamento associado a essa assinatura
// current_period_start Início do periodo de cobrança da assinatura
// current_period_end   Término do periodo de cobrança da assinatura
// charges              Numero de cobranças que foram efetuadas na assinatura, sem contar a cobrança inicial da
// 	                assinatura no caso de cartão de crédito.
//
// status               Possíveis estados da transaçãov ou assinatura. Valores possíveis: trialing, paid,
//                      pending_payment, unpaid, canceled e ended
//
// date_created         Data de criação da assinatura
// phone                Objeto com dados do telefone do cliente
// address              Objeto com dados do endereço do cliente
// custormer            Objeto com dados do cliente
// card                 Objeto com dados do cartão do cliente
// metadata             Objeto com dados adicionais do cliente ou produto ou serviço vendido
type Subscription struct {
	Object             string                  `json:"object"`
	Plan               plan.Plan               `json:"plan"`
	Id                 int                     `json:"id"`
	CurrentTransaction transaction.Transaction `json:"current_transaction"`
	PostbackUrl        string                  `json:"postback_url"`
	PaymentMethod      string                  `json:"payment_method"`
	CurrentPeriodStart string                  `json:"current_period_start"`
	CurrentPeriodEnd   string                  `json:"current_period_end"`
	Charges            int                     `json:"charges"`
	Status             string                  `json:"status"`
	DateCreated        time.Time               `json:"date_created"`
	Phone              customer.Phones         `json:"phone"`
	Address            customer.Addresses      `json:"address"`
	Custormer          customer.Customer       `json:"custormer"`
	Card               card.Card               `json:"card"`
	Metadata           Metadata                `json:"metadata"`
}

type Metadata struct{}

func (s *Subscription) Cancel(p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/cancel"
	req := request.Client{}
	_, err, errApi := req.New("POST", url.Values{"route": {route}}, []byte(`{}`), s)
	return *s, err, errApi
}

func (s *Subscription) Create(d []byte, p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	req := request.Client{}
	_, err, errApi := req.New("POST", url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Subscription) Get(p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Subscription) GetAll(p url.Values, h auth.Headers) ([]Subscription, error, liberr.ErrorsAPI) {
	res := []Subscription{}
	req := request.Client{}
	_, err, errApi := req.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}

func (s *Subscription) Update(d []byte, p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	req := request.Client{}
	_, err, errApi := req.New("UPDATE", url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}
