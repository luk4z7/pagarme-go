// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operation

import "time"

// date_created             Data de criação da operação no formato ISODate
// date_updated             Data de atualização da operação no formato ISODate
// ended_at                 Data de termíno do processamento da operação no formato Unix Timestamp
// fail_reason              Motivo da falha da operação
// group_id                 Id do grupo de operações que essa operação pertence
// id                       Id da operação
// metadata                 Informações sobre o ambiente da operação
// model                    Modelo da operação
// model_id                 Id do modelo da transação
// next_group_id            Id do próximo grupo de operações
// processor                Adquirente que processou essa operação
// processor_response_code  Código de resposta retornado pelo adquirente
// request_id 		    Id da requisição interna que disparou essa operação
// rollbacked 		    Indicador de operação desfeita
// started_at 		    Data de início do processamento da operação no formato Unix Timestamp
// status 		    Status da operação. Valores possíveis: waiting, processing, deferred, failed, success e dropped
// type 		    Tipo da operação. Valores possíveis: analyze, authorize, capture, issue, conciliate e refund
type Operation struct {
	DateCreated           time.Time `json:"date_created"`
	DateUpdated           time.Time `json:"date_updated"`
	EndedAt               int       `json:"ended_at"`
	FailReason            string    `json:"fail_reason"`
	GroupId               string    `json:"group_id"`
	Id                    string    `json:"id"`
	Metadata              Metadata  `json:"metadata"`
	Model                 string    `json:"model"`
	ModelId               string    `json:"model_id"`
	NextGroupId           string    `json:"next_group_id"`
	Processor             string    `json:"processor"`
	ProcessorResponseCode string    `json:"processor_response_code"`
	RequestId             string    `json:"request_id"`
	Rollbacked            bool      `json:"rollbacked"`
	StartedAt             int       `json:"started_at"`
	Status                string    `json:"status"`
	Type                  string    `json:"type"`
}

type Metadata struct {
	Environment Environment `json:"environment"`
}

type Environment struct {
	CapturedAmount    int    `json:"captured_amount"`
	AuthorizationCode string `json:"authorization_code"`
	AuthorizedAmount  int    `json:"authorized_amount"`
	Nsu               int    `json:"nsu"`
	ResponseCode      string `json:"response_code"`
	Tid               int    `json:"tid"`
}
