// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"net/url"
	"time"
	"github.com/luk4z7/pagarme-go/lib/plan"
	"github.com/luk4z7/pagarme-go/lib/transaction"
	"github.com/luk4z7/pagarme-go/lib/customer"
	"github.com/luk4z7/pagarme-go/lib/card"
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositorySubscription repository.Repository

const (
	endPoint = "https://api.pagar.me/1/subscriptions"
)

type Subscription struct {
	Object 		   string                  `json:"object"`               // Nome do tipo de objeto criado ou
								                 // modificado. Valor retornado: subscription

	Plan 		   plan.Plan               `json:"plan"`                 // Objeto com os dados do plano que a
										 // assinatura está associada

	Id 		   int                     `json:"id"`                   // Número do identificador do plano
	CurrentTransaction transaction.Transaction `json:"current_transaction"`  // Objeto com os dados da ultima
							                         // transação realizada pela assinatura

	PostbackUrl        string                  `json:"postback_url"`         // Endpoint da aplicação integrada ao
								                 // Pagar.me que irá receber os jsons de
                                                                                 // resposta a cada atualização dos processos

	PaymentMethod      string                  `json:"payment_method"`       // Método de pagamento associado a
                                                                                 // essa assinatura

	CurrentPeriodStart string                  `json:"current_period_start"` // Início do periodo de cobrança da assinatura
	CurrentPeriodEnd   string                  `json:"current_period_end"`   // Término do periodo de cobrança da assinatura
	Charges            int                     `json:"charges"`              // Numero de cobranças que foram
                                                                                 // efetuadas na assinatura, sem contar
                                                                                 // a cobrança inicial da assinatura no
                                                                                 // caso de cartão de crédito.

	Status 	           string                  `json:"status"`               // Possíveis estados da transaçãov ou
                                                                                 // assinatura. Valores possíveis: trialing,
                                                                                 // paid, pending_payment,  unpaid,
                                                                                 // canceled e ended

	DateCreated 	   time.Time               `json:"date_created"`         // Data de criação da assinatura
	Phone 		   customer.Phones         `json:"phone"`                // Objeto com dados do telefone do cliente
	Address 	   customer.Addresses      `json:"address"`              // Objeto com dados do endereço do cliente
	Custormer 	   customer.Customer       `json:"custormer"`            // Objeto com dados do cliente
	Card 		   card.Card               `json:"card"`                 // Objeto com dados do cartão do cliente
	Metadata 	   Metadata                `json:"metadata"`             // Objeto com dados adicionais do
                                                                                 // cliente ou produto ou serviço vendido
}

type Metadata struct {

}

func (s *Subscription) Cancel(p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id") + "/cancel"
	_, err, errApi := repositorySubscription.Create(url.Values{"route": {route}}, []byte(`{}`), s)
	return *s, err, errApi
}

func (s *Subscription) Create(d []byte, p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	_, err, errApi := repositorySubscription.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Subscription) Get(p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositorySubscription.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Subscription) GetAll(p url.Values, h auth.Headers) ([]Subscription, error, liberr.ErrorsAPI) {
	res := []Subscription{}
	_, err, errApi := repositorySubscription.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}

func (s *Subscription) Update(d []byte, p url.Values, h auth.Headers) (Subscription, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositorySubscription.Update(url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}