// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"time"
	"net/url"
	"github.com/luk4z7/pagarme-go/lib/customer"
	"github.com/luk4z7/pagarme-go/lib/card"
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/transaction/splitrule"
	"github.com/luk4z7/pagarme-go/repository"
)

var repositoryTransaction repository.Repository

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

type Transaction struct {
	Object               string    		    `json:"object"`                 // Nome do tipo do objeto criado/modificado.
										    // Valor retornado: transaction

	Status               string    		    `json:"status"`		    // Para cada atualização no processamento
										    // da transação, esta propriedade será
										    // alterada, e o objeto transaction
										    // retornado como resposta através da
										    // sua URL de postback ou após o término
	                                                                            // do processamento da ação atual.
										    // Valores possíveis:  processing,
										    // authorized, paid, refunded,
										    // waiting_payment, pending_refund, refused

	RefuseReason         string    		    `json:"refuse_reason"`	    // Motivo/agente responsável pela validação
										    // ou anulação da transação.
										    // Valores possíveis:  acquirer, antifraud,
										    // internal_error, no_acquirer,
										    // acquirer_timeout

	StatusReason 	     string    		    `json:"status_reason"`	    // Adquirente responsável pelo processamento
										    // da transação.
										    // Valores possíveis:  development
										    // (em ambiente de testes),  pagarme
										    // (adquirente Pagar.me), stone, cielo,
										    // rede, mundipagg

	AcquirerResponseCode string    		    `json:"acquirer_response_code"` // Mensagem de resposta do adquirente
										    // referente ao status da transação.

	AcquirerName         string    		    `json:"acquirer_name"`
	AuthorizationCode    string    		    `json:"authorization_code"`	    // Código de autorização retornado pela
	                                                                            // bandeira.

	SoftDescriptor       string    		    `json:"soft_descriptor"`	    // Texto que irá aparecer na fatura do
										    // cliente depois do nome da loja.
	  		                                                            // OBS: Limite de 13 caracteres,
										    // apenas letras e números.

	Tid 		     interface{}            `json:"tid"`		    // Código que identifica a transação no
										    // adquirente.

	Nsu 		     interface{}     	    `json:"nsu"`		    // Código que identifica a transação no
										    // adquirente.

	DateCreated 	     time.Time 		    `json:"date_created"`	    // Data de criação da transação no formato
										    // ISODate

	DateUpdated 	     time.Time 		    `json:"date_updated"`	    // Data de última atualização da transação
										    // no formato ISODate

	Amount 		     int       		    `json:"amount"`		    // Valor em centados do que foi pago
	AuthorizedAmount     int       		    `json:"authorized_amount"`
	PaidAmount 	     int       		    `json:"paid_amount"`
	RefundedAmount 	     int       		    `json:"refunded_amount"`
	Installments 	     int       		    `json:"installments"`	    // Número de parcelas/prestações a serem
										    // cobradas

	Id 		     int       		    `json:"id"`			    // Código de identificação da transação
	Cost 		     float64       	    `json:"cost"`
	CardHolderName       string    		    `json:"card_holder_name"`	    // Nome do portador do cartão.
										    // Usado quando o cartão a ser configurado
										    // na assinatura ainda não está salvo no
										    // nosso banco de dados

	CardLastDigits 	     string    		    `json:"card_last_digits"`
	CardFirstDigits      string    		    `json:"card_first_digits"`
	CardBrand 	     string    		    `json:"card_brand"`
	PostbackUrl	     string    		    `json:"postback_url"`	    // URL para onde são enviadas as
							                            // notificações de alteração de status

	PaymentMethod 	     string    		    `json:"payment_method"`	    // Método de pagamento possíveis:
										    // credit_card e boleto

	CaptureMethod 	     string    		    `json:"capture_method"`
	AntifraudScore 	     string    		    `json:"antifraud_score"`
	BoletoUrl 	     string    		    `json:"boleto_url"`		    // URL do boleto para ser impresso
	BoletoBarcode        string    		    `json:"boleto_barcode"`	    // Código de barras do boleto gerado na
										    // transação

	BoletoExpirationDate string    		    `json:"boleto_expiration_date"` // Data de vencimento do boleto no formato
										    // ISODate

	Referer              string    	            `json:"referer"`		    // Mostra de onde a transação foi criada.
										    // Valores : api_key ou  encryption_key

	Ip                   string    		    `json:"ip"`			    // IP de origem que criou a transção,
										    // podendo ser ou do seu cliente (quando
										    // criado via checkout ou utilizando
										    // card_hash) ou do servidor.

	SubscriptionId       int       		    `json:"subscription_id"`	    // Código da assinatura
	Phones	     	     customer.Phones	    `json:"phone"`		    // Objeto do tipo phone
	Addresses	     customer.Addresses     `json:"address"`		    // Objeto do tipo address
	Customer	     customer.Customer      `json:"customer"`		    // Objeto do tipo customer
	Card		     card.Card		    `json:"card"`		    // Objeto do tipo card
	SplitRules	     []splitrule.SplitRule `json:"split_rules"`
	Metadata	     Metadata		    `json:"metadata"`
	AntifraudMetadata    AntifraudMetadata      `json:"antifraud_metadata"`
}

type Metadata struct {
	idProduto string `json:"id_produto"`
}

type AntifraudMetadata struct {
	
}

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