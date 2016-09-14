# pagarme-go

## Installation

```bash
go get github.com/luk4z7/pagarme-go
```

```bash
export PAGARME_APIKEY=YOUR_APYKEY
```

## Usage


Create a Bank:

```go
package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/bank"
)

var bankAccount bank.Account

func main() {
	data := []byte(`{
		"bank_code": "184",
		"agencia": "0809",
		"agencia_dv": "9",
		"conta": "08809",
		"conta_dv": "8",
		"document_number": "80802694594",
		"legal_name": "JORGER MENDES"
	}`)
	create, err, errorsApi := bankAccount.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	}
	responseCreate, _ := json.MarshalIndent(create, "", " ")
	os.Stdout.Write(responseCreate)
}
```

Example Response

```json
{
     "agencia": "0809",
     "agencia_dv": "9",
     "bank_code": "184",
     "conta": "08809",
     "conta_dv": "8",
     "date_created": "2016-09-14T06:45:19.395Z",
     "document_number": "80802694594",
     "document_type": "cpf",
     "id": 16740182,
     "legal_name": "JORGER MENDES",
     "object": "bank_account"
}
```


Create a Card:

```go
package main

import (
	"net/url"
	"encoding/json"
	"os"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/card"
)

var creditCard card.Card

func main() {
	data2 := []byte(`{
		"card_number": "4242424242424242",
		"card_holder_name": "Marcos Mendes Teste API Create",
		"card_expiration_date": "1018"
	}`)
	create2, err, errorsApi := creditCard.Create(data2, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate2, _ := json.MarshalIndent(create2, "", " ")
		os.Stdout.Write(responseCreate2)
	}
}
```

Example Response

```json
{
     "brand": "visa",
     "country": "US",
     "customer": "",
     "date_created": "2016-09-04T20:47:36.701Z",
     "date_updated": "2016-09-04T20:47:36.83Z",
     "fingerprint": "1fSoeUfMRR/V",
     "first_digits": "424242",
     "holder_name": "Marcos Mendes Teste API Create",
     "id": "card_cisp3at4s00fowm6egw1kgit1",
     "last_digits": "4242",
     "object": "card",
     "valid": true
}
```


Create a transaction:

```go
package main

import (
	"encoding/json"
	"os"
	"net/url"
	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/transaction"
)

var transactionRecord transaction.Transaction

func main() {
	data := []byte(`{
		"amount": 100,
		"card_id": "card_cisp3at4s00fowm6egw1kgit1",
		"customer": {
			"name":"Name",
			"email":"example@example.com",
			"document_number":"80802694594",
			"gender":"M",
			"born_at":"09-22-2015",
			"address": {
				"street":"Rua de exemplo",
				"street_number":"808",
				"neighborhood":"Bairro de exemplo",
				"complementary":"Apartamento 8",
				"city":"Cidade",
				"state":"Lordaeron",
				"zipcode":"47850000",
				"country":"Lordaeron"
			},
			"phone": {
				"ddi":"88",
				"ddd":"88",
				"number":"808080808"
			}
		},
		"postback_url": "http://requestb.in/pkt7pgpk",
		"metadata": {
			"idProduto": "10"
		}
	}`)
	create, err, errorsApi := transactionRecord.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		response, _ := json.MarshalIndent(errorsApi, "", "  ")
		os.Stdout.Write(response)
	} else {
		responseCreate, _ := json.MarshalIndent(create, "", " ")
		os.Stdout.Write(responseCreate)
	}
}
```

For more example:

in the folder `examples` exist several examples how you can use, for api reference consult
https://github.com/pagarme/api-reference

