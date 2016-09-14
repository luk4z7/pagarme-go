// Copyright 2016 The Lucas Alves Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package error

type ErrorsAPI struct {
	Errors []Errors `json:"errors"`
	Url    string   `json:"url"`
	Method string   `json:"method"`
}

type Errors struct {
	ParameterName string `json:"parameter_name"`
	Type 	      string `json:"type"`
	Message       string `json:"message"`
}

type Err struct {
	Name string
}

func (e *Err) Error() string {
	return e.Name
}

// Return panic when error != nil
func Check(e error, m string) {
	if e != nil {
		if m == "" {
			panic(m)
		} else {
			panic(e)
		}
	}
}

