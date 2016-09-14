// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/subscription"
	subscriptionTransaction "github.com/luk4z7/pagarme-go/lib/subscription/transaction"
	subscriptionEvent "github.com/luk4z7/pagarme-go/lib/subscription/event"
)

var subscriptionRecord subscription.Subscription

var subscriptionTransactions subscriptionTransaction.SubscriptionTransaction

var subscriptionEvents subscriptionEvent.SubscriptionEvent

func main() {
	// Criando uma subscription
	data := []byte(`{
		"customer": {
			"email": "api@test.com"
		},
		"plan_id": "62535",
		"card_id": "card_cisp3at4s00fowm6egw1kgit1",
		"postback_url": "http://requestb.in/zyn5obzy",
		"payment_method": "boleto"
	}`)
	create, err, errorsApi := subscriptionRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}


	// Retornando uma assinatura
	get, err, errorsApi := subscriptionRecord.Get(url.Values{"id": {"99815"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}


	// Retornando todas as assinaturas
	getall, err, errorsApi := subscriptionRecord.GetAll(url.Values{}, auth.Headers{
		"page"  : "1",
		"count" : "10",
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAll, _ := json.MarshalIndent(getall, "", " ")
		os.Stdout.Write(responseGetAll)
	}


	// Atualizando um recebedor com uma outra conta bancária existente
	data4 := []byte(`{
		"plan_id": "62544",
		"card_id": "card_cisp3at4s00fowm6egw1kgit1"
	}`)
	update, err, errorsApi := subscriptionRecord.Update(data4, url.Values{"id" : {"99815"}}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseUpdate, _ := json.MarshalIndent(update, "", "   ")
		os.Stdout.Write(responseUpdate)
	}


	// Cancelando uma assinatura
	cancel, err, errorsApi := subscriptionRecord.Cancel(url.Values{
		"id": {"99815"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCancel, _ := json.MarshalIndent(cancel, "", " ")
		os.Stdout.Write(responseCancel)
	}


	// Transações de uma assinatura
	getTransactions, err, errorsApi := subscriptionTransactions.Get(url.Values{
		"id": {"99815"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseTransactions, _ := json.MarshalIndent(getTransactions, "", " ")
		os.Stdout.Write(responseTransactions)
	}


	// Retornar um evento de uma assinatura
	getEvent, err, errorsApi := subscriptionEvents.Get(url.Values{
		"subscription_id": {"99815"},
		"id": {"99815"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetEvent, _ := json.MarshalIndent(getEvent, "", " ")
		os.Stdout.Write(responseGetEvent)
	}


	// Retornar todos os eventos de uma assinatura
	getallEvents, err, errorsApi := subscriptionEvents.GetAll(url.Values{
		"subscription_id": {"99815"},
	}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGetAllEvents, _ := json.MarshalIndent(getallEvents, "", " ")
		os.Stdout.Write(responseGetAllEvents)
	}
}