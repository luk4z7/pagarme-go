// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"net/url"
	liberr "github.com/luk4z7/pagarme-go/error"
)

// Check the values the map passed
func MustBeNotEmpty(values url.Values, require func() []string) (bool, error) {

	required := require()

	valid := make(map[int]string)
	for k1, v := range required {
		for k, v2 := range values {
			if k == v {
				if len(v2[0]) == 0 {
					return false, &liberr.Err{Name:"Parametros invalidos"}
				}
				valid[k1] = k
			}
		}
	}

	if len(valid) != len(required) {
		return false, &liberr.Err{Name:"Parametros invalidos"}
	}

	return true, nil
}

func IsEmpty(values url.Values) error {
	_, err := MustBeNotEmpty(values, func() []string {
		return []string{""}
	})
	return err
}