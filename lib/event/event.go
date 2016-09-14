// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

type Event struct {
	Id      string    `json:"id"`       // Id do evento
	Model   string    `json:"model"`    // Objeto associado a esse evento
	ModelId string    `json:"model_id"` // Id do objeto associado a esse evento
	Name    string    `json:"name"`     // Nome do evento
	Object  string    `json:"object"`   // Nome do tipo do objeto criado/modificado
	Payload Payload `json:"payload"`  // Objeto com status dos eventos
}

type Payload struct {
	CurrentStatus string `json:"current_status"`
	DesiredStatus string `json:"desired_status"`
	OldStatus     string `json:"old_status"`
}