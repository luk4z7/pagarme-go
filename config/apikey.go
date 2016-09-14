// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"syscall"
)

type Config interface {
	GetApiKey() string
}

type ApiKey struct {
	Key string
}

func (c *ApiKey) GetApiKey() string {
	v, f := syscall.Getenv("PAGARME_APIKEY")
	if f != true {
		panic("configure sua variavel de ambiente")
	}
	return v
}