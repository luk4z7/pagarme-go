// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/search"
)

var searchRecord search.Search

func main() {
	get, err, errorsApi := searchRecord.Get(url.Values{}, auth.Headers{
		"type"  : "transaction",
		"query" : `{
			  "query": {
			    "filtered": {
			      "query": {"match_all": {}},
			      "filter": {
				"and": [
				  {
				    "range": {
				      "date_created": {
					"lte": "2016-01-31"
				      }
				    }
				  },
				  {
				    "or": [
				      {"term": {"status": "waiting_payment"}},
				      {"term": {"status": "paid"}}
				    ]
				  }
				]
			      }
			    }
			  },
			  "sort": [
			    {
			      "date_created": {"order": "desc"}
			    }
			  ],
			  "size": 5,
			  "from": 0
		}`,
	})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseGet, _ := json.MarshalIndent(get, "", " ")
		os.Stdout.Write(responseGet)
	}
}