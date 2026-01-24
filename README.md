WEDOS DNS for [`libdns`](https://github.com/libdns/libdns)
=======================

[![Go Reference](https://pkg.go.dev/badge/test.svg)](https://pkg.go.dev/github.com/libdns/TODO:PROVIDER_NAME)

This package implements the [libdns interfaces](https://github.com/libdns/libdns) for WEDOS, 
allowing you to manage DNS records.

# Authentication
WEDOS API (WAPI) doesn't use API keys, but rather login username and password (precisely their SHA-1 hash). To get started,
please see the official instructions [here](https://kb.wedos.global/wapi-manual/#activate). Please note that your WAPI password
is different from your Wedos account password.

When your WAPI account is ready, put your username (email) in the `Username` field and your WAPI password in the `Password` field.
You can also use enviromental variables instaed of hardcoding your credentials.

# Usage
```go
package main

import (
	"context"
	"os"
	"log"

	"github.com/libdns/wedos"
)

func main() {
	provider := wedos.Provider{
		Username: os.Getenv("WEDOS_USERNAME"),
		Password: os.Getenv("WEDOS_PASSWORD"),
	}

	records, err := provider.GetRecords(context, "example.org")
	if err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	fmt.Printf("%#v", records)
}
```

# Known limitations
- WEDOS allows `TTL` values from `300` to `172800`. Setting your records with `TTL=0` will fail.
