// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bank

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositoryBank repository.Repository

const (
	endPoint = "https://api.pagar.me/1/bank_accounts"
)

type Account struct {
	Agencia        string    `json:"agencia"`	  // Agência bancária
	AgenciaDv      string    `json:"agencia_dv"`      // Dígito verificador da agência bancária
	BankCode       string    `json:"bank_code"`       // Código do banco
	Conta	       string    `json:"conta"` 	  // Número da conta bancária
	ContaDv	       string    `json:"conta_dv"`        // Dígito verificador da conta bancária
	DateCreated    time.Time `json:"date_created"`    // Data de criação da conta no formato ISODate
	DocumentNumber string    `json:"document_number"` // Documento identificador do titular da conta (CPF ou CNPJ)
	DocumentType   string    `json:"document_type"`   // Tipo do documento identificador do titular da conta
	Id	       int       `json:"id"`	          // Id da conta bancária
	LegalName      string    `json:"legal_name"`      // Nome completo (se pessoa física) ou razão social (se pessoa jurídica)
	Object	       string    `json:"object"`          // Nome do tipo do objeto criado/modificado
}

func (s *Account) Create(d []byte, p url.Values, h auth.Headers) (Account, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryBank.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Account) Get(p url.Values, h auth.Headers) (Account, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryBank.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Account) GetAll(p url.Values, h auth.Headers) ([]Account, error, liberr.ErrorsAPI) {
	res := []Account{}
	_, err, errApi := repositoryBank.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}