// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/postback"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
)

var repositoryPostback repository.Repository

type SubscriptionPostback struct {
	postback postback.Postback
}

const (
	endPoint = "https://api.pagar.me/1/subscriptions"
)

func (s *SubscriptionPostback) Get(p url.Values, h auth.Headers) (postback.Postback, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/postbacks/" + p.Get("id")
	_, err, errApi := repositoryPostback.Get(url.Values{"route": {route}}, &s.postback, h)
	return s.postback, err, errApi
}

func (s *SubscriptionPostback) GetAll(p url.Values, h auth.Headers) ([]postback.Postback, error, liberr.ErrorsAPI) {
	res := []postback.Postback{}
	route := endPoint + "/" + p.Get("transaction_id") + "/postbacks"
	_, err, errApi := repositoryPostback.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}

func (s *SubscriptionPostback) Redeliver(p url.Values, h auth.Headers) (postback.Postback, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/postbacks/" + p.Get("id") + "/redeliver"
	_, err, errApi := repositoryPostback.Create(url.Values{"route": {route}}, []byte(`{}`), s)
	return s.postback, err, errApi
}
