// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postback

import "time"

type Postback struct {
	DateCreated time.Time    `json:"date_created"`
	DateUpdated time.Time    `json:"date_updated"`
	Deliveries  []Deliveries `json:"deliveries"`
	Headers     string 	 `json:"headers"`
	Id 	    string 	 `json:"id"`
	Model 	    string 	 `json:"model"`
	ModelId     string 	 `json:"model_id"`
	NextRetry   string 	 `json:"next_retry"`
	Object 	    string 	 `json:"object"`
	Payload     string 	 `json:"payload"`
	RequestUrl  string 	 `json:"request_url"`
	Retries     int 	 `json:"retries"`
	Signature   string  	 `json:"signature"`
	Status      string 	 `json:"status"`
}

type Deliveries struct {
	DateCreated 	time.Time `json:"date_created"`
	DateUpdated 	time.Time `json:"date_updated"`
	Id 	    	string 	  `json:"id"`
	Object 		string    `json:"object"`
	ResponseBody    string    `json:"response_body"`
	ResponseHeaders string    `json:"response_headers"`
	ResponseTime 	int       `json:"response_time"`
	Status 		string    `json:"status"`
	StatusCode 	string    `json:"status_code"`
	StatusReason 	string    `json:"status_reason"`
}