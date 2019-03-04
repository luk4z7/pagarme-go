// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

import (
	"github.com/luk4z7/pagarme-go/auth"
	liberr "github.com/luk4z7/pagarme-go/error"
	"github.com/luk4z7/pagarme-go/repository"
	"net/url"
)

var repositorySearch repository.Repository

const (
	endPoint = "https://api.pagar.me/1/search"
)

type Search struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total    int        `json:"total"`
	MaxScore int        `json:"max_score"`
	HitsInto []HitsInto `json:"hits"`
}

type HitsInto struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	Id     string      `json:"_id"`
	Score  string      `json:"_score"`
	Source interface{} `json:"_source"`
}

func (s *Search) Get(p url.Values, h auth.Headers) (Search, error, liberr.ErrorsAPI) {
	_, err, errApi := repositorySearch.Get(url.Values{"route": {endPoint}}, s, h)
	return *s, err, errApi
}
