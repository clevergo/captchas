# Captchas
[![Build Status](https://img.shields.io/travis/clevergo/captchas?style=for-the-badge)](https://travis-ci.org/clevergo/captchas)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/captchas?style=for-the-badge)](https://coveralls.io/github/clevergo/captchas?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/captchas?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/captchas?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/captchas)

Base64 Captchas Manager, supports multiple [drivers](#drivers) and [stores](#stores).

## Drivers

- [Digit](drivers#digit)
- [Math](drivers#math)
- [Audio](drivers#audio)
- [String](drivers#string)
- [Chinese](drivers#chinese)

## Stores

- [Memory Store](stores/memstore)
- [Redis Store](stores/redisstore)
- [Memcached Store](stores/memcachedstore)
- [MySQL Store](stores/mysqlstore)
- [SQLite3 Store](stores/sqlite3store)
- [PostgreSQL Store](stores/postgresstore)
- Add your store here by PR or [request a new store](https://github.com/clevergo/captchas/issues/new).

## Usage

Checkout [example](https://github.com/clevergo/examples/tree/master/captchas) for details.

### Quick Start

```shell
$ go get clevergo.tech/captchas \
	clevergo.tech/captchas/drivers \
	clevergo.tech/captchas/stores/memstore
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"clevergo.tech/captchas"
	"clevergo.tech/captchas/drivers"
	"clevergo.tech/captchas/stores/memstore"
)

var (
	store   = memstore.New()              // memory store.
	driver  = drivers.NewDigit()          // digit driver.
	manager = captchas.New(store, driver) // manager
	tmpl    = template.Must(template.New("captcha").Parse(`
<html>
<body>
<form action="/validate" method="POST">
	<input name="captcha">
	{{ .captcha.HTMLField "captcha_id" }}
	<input type="submit" value="Submit">
</form>
</body>
</html>
	`))
)

func main() {
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/validate", validate)
	log.Println(http.ListenAndServe(":8080", http.DefaultServeMux))
}

// generates a new captcha
func generate(w http.ResponseWriter, r *http.Request) {
	captcha, err := manager.Generate(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
                return
	}

	// returns JSON data.
	if r.URL.Query().Get("format") == "json" {
		v := map[string]string{
			"captcha_id":   captcha.ID(),             // captcha ID.
			"captcha_data": captcha.EncodeToString(), // base64 encode string.
		}
		data, _ := json.Marshal(v)
		w.Write(data)
		return
	}

	// render captcha via template.
	tmpl.Execute(w, map[string]interface{}{
		"captcha": captcha,
	})

}

// validates a captcha.
func validate(w http.ResponseWriter, r *http.Request) {
	captchaID := r.PostFormValue("captcha_id")
	captcha := r.PostFormValue("captcha")

	// verify
	if err := manager.Verify(r.Context(), captchaID, captcha, true); err != nil {
		io.WriteString(w, fmt.Sprintf("captcha is invalid: %s", err.Error()))
		return
	}

	io.WriteString(w, "valid")
}
```
