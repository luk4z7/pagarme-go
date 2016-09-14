// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package customer

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositoryCustomer repository.Repository

const (
	endPoint = "https://api.pagar.me/1/customers"
)

type Customer struct {
	Addresses      []Addresses `json:"addresses"`       // Lista de endereços do Customere
	BornAt         string      `json:"born_at"`         // Data de nascimento do Customere no formato ISODate
	DateCreated    time.Time   `json:"date_created"`    // Data de criação do Customere no formato ISODate
	DocumentNumber string      `json:"document_number"` // Número do CPF ou CNPJ do Customere
	DocumentType   string      `json:"document_type"`   // Tipo do documento do Customere
	Email          string      `json:"email"`           // E-mail do Customere
	Gender         string      `json:"gender"`          // Gênero do Customere
	Id             int         `json:"id"`              // Id do Customere
	Name           string      `json:"name"`            // Data de nascimento do Customere
	Object         string      `json:"object"`          // Nome do tipo do objeto criado/modificado.
	Phones         []Phones    `json:"phones"`
}

type Addresses struct {
	City 	      string `json:"city"`		    // Nome do objeto criado
	Complementary string `json:"complementary"`         // Complemento do endereço do Customere
	Country       string `json:"country"`               // País do endereço do Customere
	Id            int    `json:"id"`                    // Id do endereço
	Neighborhood  string `json:"neighborhood"`          // Bairro do Customere
	Object        string `json:"object"`                // Nome do objeto criado
	State         string `json:"state"`		    // Estado do endereço do Customere
	Street        string `json:"street"`                // Logradouro do Customere
	StreetNumber  string `json:"street_number"`         // Numero do endereço do Customere
	Zipcode       string `json:"zipcode"`               // CEP do Customere
}

type Phones struct {
	DDD    string `json:"ddd"`			    // Numero do DDD do telefone
	DDI    string `json:"ddi"`                          // Número do DDI do telefone
	Id     int     `json:"id"`			    // Id gerado pelo sistema para o telefone criado
	Number string `json:"number"`			    // Numero do telefone do Customere
	Object string `json:"object"`                       // Nome do objeto criado
}

func (s *Customer) Create(d []byte, p url.Values, h auth.Headers) (Customer, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryCustomer.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Customer) Get(p url.Values, h auth.Headers) (Customer, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryCustomer.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Customer) GetAll(p url.Values, h auth.Headers) ([]Customer, error, liberr.ErrorsAPI) {
	res := []Customer{}
	_, err, errApi := repositoryCustomer.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}