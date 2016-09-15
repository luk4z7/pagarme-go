// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package card

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
	"time"
)

var repositoryCard repository.Repository

const (
	endPoint = "https://api.pagar.me/1/cards"
)

type Card struct {
	Brand       string    `json:"brand"`        // Bandeira do cartão
	Country     string    `json:"country"`      // País do cartão
	Customer    string    `json:"customer"`     // Cliente associado ao cartão
	DateCreated time.Time `json:"date_created"` // Data de criação do cartão no formato ISODate
	DateUpdated time.Time `json:"date_updated"` // Data de atualização cartão no formato ISODate
	Fingerprint string    `json:"fingerprint"`  // Identificador do cartão na nossa base
	FirstDigits string    `json:"first_digits"` // Primeiros digitos do cartão
	HolderName  string    `json:"holder_name"`  // Nome do portador do cartão
	Id          string    `json:"id"`           // Id do cartão
	LastDigits  string    `json:"last_digits"`  // Últimos digitos do cartão
	Object      string    `json:"object"`       // Nome do tipo do objeto criado/modificado.
	Valid       bool      `json:"valid"`        // Indicador de cartão válido
}

func (s *Card) Create(d []byte, p url.Values, h auth.Headers) (Card, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryCard.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Card) Get(p url.Values, h auth.Headers) (Card, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryCard.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}
