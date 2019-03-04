// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plan

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
	"time"
)

var repositoryRecipient repository.Repository

const (
	endPoint = "https://api.pagar.me/1/plans"
)

// object           Nome do tipo do objeto criado/modificado
// id               Id do plano
// amount           Valor do plano em centavos
// days             Dias para efetuação da próxima cobrança da assinatura atrelada ao plano
// name             Nome do plano
// trial_days       Dias que o usuário poderá testar o serviço gratuitamente
// date_created     Data da criação do plano (ISODate)
// payment_methods  Formas de pagamento aceitas no plano
// charges          Número de cobranças que podem ser feitas em uma assinatura
// installments     Informa em quantas vezes o pagamento será parcelado entre cada cobrança
type Plan struct {
	Object         string      `json:"object"`
	Id             int         `json:"id"`
	Amount         int         `json:"amount"`
	Days           int         `json:"days"`
	Name           string      `json:"name"`
	TrialDays      int         `json:"trial_days"`
	DateCreated    time.Time   `json:"date_created"`
	PaymentMethods interface{} `json:"payment_methods"`
	Charges        int         `json:"charges"`
	Installments   int         `json:"installments"`
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
