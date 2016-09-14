// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operation

import "time"

type Operation struct {
	DateCreated           time.Time `json:"date_created"`            // Data de criação da operação no formato ISODate
	DateUpdated           time.Time `json:"date_updated"`            // Data de atualização da operação no formato ISODate
	EndedAt               int       `json:"ended_at"`                // Data de termíno do processamento da operação
									 // no formato Unix Timestamp

	FailReason            string    `json:"fail_reason"`             // Motivo da falha da operação
	GroupId               string    `json:"group_id"`                // Id do grupo de operações que essa operação pertence
	Id                    string    `json:"id"`                      // Id da operação
	Metadata              Metadata  `json:"metadata"`                // Informações sobre o ambiente da operação
	Model                 string    `json:"model"`                   // Modelo da operação
	ModelId               string    `json:"model_id"`                // Id do modelo da transação
	NextGroupId           string    `json:"next_group_id"`           // Id do próximo grupo de operações
	Processor   	      string    `json:"processor"`               // Adquirente que processou essa operação
	ProcessorResponseCode string    `json:"processor_response_code"` // Código de resposta retornado pelo adquirente
	RequestId             string    `json:"request_id"` 		 // Id da requisição interna que disparou essa operação
	Rollbacked            bool      `json:"rollbacked"` 		 // Indicador de operação desfeita
	StartedAt             int       `json:"started_at"` 		 // Data de início do processamento da operação
								         // no formato Unix Timestamp

	Status                string    `json:"status"` 		 // Status da operação. Valores possíveis:
									 // waiting, processing, deferred,
									 // failed, success e dropped

	Type                  string    `json:"type"` 			 // Tipo da operação. Valores possíveis:
									 // analyze, authorize, capture, issue,
									 // conciliate e refund
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