// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

import (
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/repository"
	liberr "github.com/luk4z7/pagarme-go/error"
)

var repositoryBalance repository.Repository

const (
	endPoint = "https://api.pagar.me/1/balance"
)

type BalanceAmount struct {
	Available    int    `json:"amount"` // Valor em centavos que você tem disponível em sua conta Pagar.me
	Object       string `json:"object"` // Nome do tipo do objeto criado/modificado.
	Transferred  int    `json:"amount"` // Valor em centavos que você já transferiu para sua conta bancária
					    // (quanto já recebeu efetivamente)
	WaitingFunds int    `json:"amount"` // Valor em centavos que você tem a receber do Pagar.me
}

func (s *BalanceAmount) GetAll(p url.Values, h auth.Headers) (BalanceAmount, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryBalance.Get(url.Values{"route": {endPoint}}, s, h)
	return *s, err, errApi
}