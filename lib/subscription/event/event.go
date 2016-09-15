// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package subscription

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/lib/event"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
)

var repositorySubscriptionEvent repository.Repository

type SubscriptionEvent struct {
	event event.Event
}

const (
	endPoint = "https://api.pagar.me/1/subscriptions"
)

func (s *SubscriptionEvent) Get(p url.Values, h auth.Headers) (event.Event, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("subscription_id") + "/events/" + p.Get("id")
	_, err, errApi := repositorySubscriptionEvent.Get(url.Values{"route": {route}}, &s.event, h)
	return s.event, err, errApi
}

func (s *SubscriptionEvent) GetAll(p url.Values, h auth.Headers) ([]event.Event, error, liberr.ErrorsAPI) {
	res := []event.Event{}
	route := endPoint + "/" + p.Get("subscription_id") + "/events"
	_, err, errApi := repositorySubscriptionEvent.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}
