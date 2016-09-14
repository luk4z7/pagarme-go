// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plan

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/repository"
	liberr "github.com/luk4z7/pagarme-go/error"
)

var repositoryRecipient repository.Repository

const (
	endPoint = "https://api.pagar.me/1/plans"
)

type Plan struct {
	Object         string            `json:"object"`          // Nome do tipo do objeto criado/modificado
	Id             int               `json:"id"`              // Id do plano
	Amount         int               `json:"amount"`          // Valor do plano em centavos
	Days           int               `json:"days"`            // Dias para efetuação da próxima cobrança da
								  // assinatura atrelada ao plano

	Name 	       string            `json:"name"`            // Nome do plano
	TrialDays      int               `json:"trial_days"`      // Dias que o usuário poderá testar o serviço
	                                                          // gratuitamente

	DateCreated    time.Time         `json:"date_created"`    // Data da criação do plano (ISODate)
	PaymentMethods interface{}       `json:"payment_methods"` // Formas de pagamento aceitas no plano
	Charges        int               `json:"charges"`         // Número de cobranças que podem ser feitas em uma
						                  // assinatura

	Installments   int               `json:"installments"`    // Informa em quantas vezes o pagamento será parcelado
						                  // entre cada cobrança
}

func (s *Plan) Create(d []byte, p url.Values, headers auth.Headers) (Plan, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryRecipient.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Plan) Get(p url.Values, h auth.Headers) (Plan, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryRecipient.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Plan) GetAll(p url.Values, h auth.Headers) ([]Plan, error, liberr.ErrorsAPI) {
	res := []Plan{}
	_, err, errApi := repositoryRecipient.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}

func (s *Plan) Update(d []byte, p url.Values, h auth.Headers) (Plan, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryRecipient.Update(url.Values{"route": {route}}, d, s)
	return *s, err, errApi
}