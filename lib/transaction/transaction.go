// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/card"
	"github.com/luk4z7/pagarme-go/lib/customer"
	"github.com/luk4z7/pagarme-go/lib/transaction/splitrule"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
	"time"
)

var repositoryTransaction repository.Repository

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

// object                   Nome do tipo do objeto criado/modificado. Valor retornado: transaction
// status		    Para cada atualização no processamento da transação, esta propriedade será alterada, e o
// 			    objeto transaction retornado como resposta através da sua URL de postback ou após o término
// 			    do processamento da ação atual. Valores possíveis:  processing, authorized, paid, refunded,
// 			    waiting_payment, pending_refund, refused
//
// refuse_reason	    Motivo/agente responsável pela validação ou anulação da transação. Valores possíveis:
// 			    acquirer, antifraud, internal_error, no_acquirer, acquirer_timeout
//
// status_reason	    Adquirente responsável pelo processamento da transação. Valores possíveis:  development
// 			    (em ambiente de testes),  pagarme (adquirente Pagar.me), stone, cielo, rede, mundipagg
//
// acquirer_response_code   Mensagem de resposta do adquirente referente ao status da transação.
// acquirer_name
// authorization_code	    Código de autorização retornado pela bandeira.
// soft_descriptor	    Texto que irá aparecer na fatura do cliente depois do nome da loja. OBS: Limite de 13
// 			    caracteres, apenas letras e números.
//
// tid		    	    Código que identifica a transação no adquirente.
// nsu		    	    Código que identifica a transação no adquirente.
// date_created	    	    Data de criação da transação no formato ISODate
// date_updated	    	    Data de última atualização da transação no formato ISODate
// amount		    Valor em centados do que foi pago
// authorized_amount
// paid_amount
// refunded_amount
// installments	    	    Número de parcelas/prestações a serem cobradas
// id			    Código de identificação da transação
// cost
// card_holder_name	    Nome do portador do cartão. Usado quando o cartão a ser configurado na assinatura ainda não
// 			    está salvo no nosso banco de dados
//
// card_last_digits
// card_first_digits
// card_brand
// postback_url	            URL para onde são enviadas as notificações de alteração de status
// payment_method	    Método de pagamento possíveis: credit_card e boleto
// capture_method
// antifraud_score
// boleto_url		    URL do boleto para ser impresso
// boleto_barcode	    Código de barras do boleto gerado na transação
// boleto_expiration_date   Data de vencimento do boleto no formato ISODate
// referer		    Mostra de onde a transação foi criada. Valores : api_key ou  encryption_key
// ip			    IP de origem que criou a transção, podendo ser ou do seu cliente (quando criado via checkout
// 			    ou utilizando card_hash) ou do servidor.
//
// subscription_id	    Código da assinatura
// phone		    Objeto do tipo phone
// address		    Objeto do tipo address
// customer		    Objeto do tipo customer
// card		    	    Objeto do tipo card
// split_rules
// metadata
// antifraud_metadata
type Transaction struct {
	Object               string                `json:"object"`
	Status               string                `json:"status"`
	RefuseReason         string                `json:"refuse_reason"`
	StatusReason         string                `json:"status_reason"`
	AcquirerResponseCode string                `json:"acquirer_response_code"`
	AcquirerName         string                `json:"acquirer_name"`
	AuthorizationCode    string                `json:"authorization_code"`
	SoftDescriptor       string                `json:"soft_descriptor"`
	Tid                  interface{}           `json:"tid"`
	Nsu                  interface{}           `json:"nsu"`
	DateCreated          time.Time             `json:"date_created"`
	DateUpdated          time.Time             `json:"date_updated"`
	Amount               int                   `json:"amount"`
	AuthorizedAmount     int                   `json:"authorized_amount"`
	PaidAmount           int                   `json:"paid_amount"`
	RefundedAmount       int                   `json:"refunded_amount"`
	Installments         int                   `json:"installments"`
	Id                   int                   `json:"id"`
	Cost                 float64               `json:"cost"`
	CardHolderName       string                `json:"card_holder_name"`
	CardLastDigits       string                `json:"card_last_digits"`
	CardFirstDigits      string                `json:"card_first_digits"`
	CardBrand            string                `json:"card_brand"`
	PostbackUrl          string                `json:"postback_url"`
	PaymentMethod        string                `json:"payment_method"`
	CaptureMethod        string                `json:"capture_method"`
	AntifraudScore       string                `json:"antifraud_score"`
	BoletoUrl            string                `json:"boleto_url"`
	BoletoBarcode        string                `json:"boleto_barcode"`
	BoletoExpirationDate string                `json:"boleto_expiration_date"`
	Referer              string                `json:"referer"`
	Ip                   string                `json:"ip"`
	SubscriptionId       int                   `json:"subscription_id"`
	Phones               customer.Phones       `json:"phone"`
	Addresses            customer.Addresses    `json:"address"`
	Customer             customer.Customer     `json:"customer"`
	Card                 card.Card             `json:"card"`
	SplitRules           []splitrule.SplitRule `json:"split_rules"`
	Metadata             Metadata              `json:"metadata"`
	AntifraudMetadata    AntifraudMetadata     `json:"antifraud_metadata"`
}

type Metadata struct {
	IdProduto string `json:"id_produto"`
}

type AntifraudMetadata struct{}

func (s *Transaction) Create(d []byte, p url.Values, h auth.Headers) (Transaction, error, liberr.ErrorsAPI) {
	_, err, errApi := repositoryTransaction.Create(url.Values{"route": {endPoint}}, d, s)
	return *s, err, errApi
}

func (s *Transaction) Get(p url.Values, h auth.Headers) (Transaction, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("id")
	_, err, errApi := repositoryTransaction.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *Transaction) GetAll(p url.Values, h auth.Headers) ([]Transaction, error, liberr.ErrorsAPI) {
	res := []Transaction{}
	_, err, errApi := repositoryTransaction.Get(url.Values{"route": {endPoint}}, &res, h)
	return res, err, errApi
}
