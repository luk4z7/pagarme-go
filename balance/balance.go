// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
)

var repositoryBalance repository.Repository

const (
	endPoint = "https://api.pagar.me/1/balance"
)

// amount  Valor em centavos que você tem disponível em sua conta Pagar.me
// object  Nome do tipo do objeto criado/modificado.
// amount  Valor em centavos que você já transferiu para sua conta bancária (quanto já recebeu efetivamente)
// amount  Valor em centavos que você tem a receber do Pagar.me
type BalanceAmount struct {
	Available    int    `json:"amount"`
	Object       string `json:"object"`
	Transferred  int    `json:"amount"`
	WaitingFunds int    `json:"amount"`
}

func (s *BalanceAmount) GetAll(p url.Values, h auth.Headers) (BalanceAmount, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryBalance.Get(url.Values{"route": {endPoint}}, s, h)
	return *s, err, errApi
}
