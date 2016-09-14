// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operation

import (
	"time"
	"net/url"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositoryBalance repository.Repository

const (
	endPoint = "https://api.pagar.me/1/balance/operations"
)

type BalanceOperation struct {
	Amount 		int 		  `json:"amount"`	   // Valor em centavos transacionado para a conta
	DateCreated 	time.Time 	  `json:"date_created"`    // Data de criação da operação de saldo no formato ISODate
	Fee 		int 		  `json:"fee"`	 	   // Valor em centavos que foi cobrado pela transação (taxa)
	Id 		int 		  `json:"id"`		   // Id da operação de saldo
	BalanceMovement BalanceMovement   `json:"movement_object"` // Objeto da origem da operação de saldo
	Object 		string 		  `json:"object"`	   // Nome do tipo do objeto criado/modificado
	Status 		string 		  `json:"status"`	   // Estado do saldo da conta. Valores possíveis:
	   							   // waiting_funds ,  available e  transferred
	Type 		string 		  `json:"type"`		   // Tipo de objeto que gerou a operação de saldo
}

type BalanceMovement struct {
	Amount 		    int       `json:"amount"`		     // Valor em centados do que foi pago
	AnticipationFee     int       `json:"anticipation_fee"`      // Valor em centavos que foi cobrado pela antecipação (taxa)
	BulkAnticipationId  int       `json:"bulk_anticipation_id"`  //
	DateCreated         time.Time `json:"date_created"`	     // Data de criação da operação no formato ISODate
	Fee 		    int       `json:"fee"`		     // Valor em centavos que foi cobrado pela transação (taxa)
	Id 		    int       `json:"id"`		     // Id do recebível
	Installment 	    int       `json:"installment"`	     // Número da parcela
	Object 		    string    `json:"object"` 		     // Nome do tipo do objeto criado/modificado
	OriginalPaymentDate string    `json:"original_payment_date"` // Dia e hora da primeira data de pagamento no formato ISODate
	PaymentDate 	    time.Time `json:"payment_date"`	     // Dia e hora do pagamento no formato ISODate
	PaymentMethod	    string    `json:"payment_method"`	     // Forma de pagamento usada
	RecipientId	    string    `json:"recipient_id"`          // Id do recebedor ao qual esse recebível pertence
	SplitRuleId 	    string    `json:"split_rule_id"`	     // Id da regra de split, se houver alguma
	Status 		    string    `json:"status"`		     // Estado atual do recebível. Valores possíveis:
								     // waiting_funds ,  paid e  suspended
	TransactionId 	    int       `json:"transaction_id"`	     // Id da transação desse recebível
	Type 		    string    `json:"type"`		     // Tipo do recebível. Valores possíveis:
	  							     // credit , refund ,  chargeback e  chargeback_refund
}

func (s *BalanceOperation) Get(p url.Values, h auth.Headers) (BalanceOperation, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryBalance.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *BalanceOperation) GetAll(p url.Values, h auth.Headers) ([]*BalanceOperation, error, liberr.ErrorsAPI) {
	res := []*BalanceOperation{}
	_, err, errApi := repositoryBalance.Get(url.Values{"route": {endPoint}}, res, h)
	return res, err, errApi
}