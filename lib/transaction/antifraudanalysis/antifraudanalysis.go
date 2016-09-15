// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package antifraudanalysis

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
	"time"
)

var repositoryAntifraudAnalysis repository.Repository

const (
	endPoint = "https://api.pagar.me/1/transactions"
)

type AntifraudAnalysis struct {
	Object      string    `json:"object"`
	Name        string    `json:"name"`
	Score       string    `json:"score"`
	Cost        string    `json:"cost"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Id          int       `json:"id"`
}

func (s *AntifraudAnalysis) Get(p url.Values, h auth.Headers) (AntifraudAnalysis, error, liberr.ErrorsAPI) {
	route := endPoint + "/" + p.Get("transaction_id") + "/antifraud_analyses/" + p.Get("id")
	_, err, errApi := repositoryAntifraudAnalysis.Get(url.Values{"route": {route}}, s, h)
	return *s, err, errApi
}

func (s *AntifraudAnalysis) GetAll(p url.Values, h auth.Headers) ([]AntifraudAnalysis, error, liberr.ErrorsAPI) {
	res := []AntifraudAnalysis{}
	route := endPoint + "/" + p.Get("transaction_id") + "/antifraud_analyses"
	_, err, errApi := repositoryAntifraudAnalysis.Get(url.Values{"route": {route}}, &res, h)
	return res, err, errApi
}
