# Captchas
[![Build Status](https://travis-ci.org/clevergo/captchas.svg?branch=master)](https://travis-ci.org/clevergo/captchas)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/captchas/badge.svg?branch=master)](https://coveralls.io/github/clevergo/captchas?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/captchas)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/captchas)](https://goreportcard.com/report/github.com/clevergo/captchas)
[![Release](https://img.shields.io/github/release/clevergo/captchas.svg?style=flat-square)](https://github.com/clevergo/captchas/releases)

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
- [Memcached Store](stores/memcached)
- Add your store here by PR or [request a new store](https://github.com/clevergo/captchas/issues/new).

## Usage

Preview: https://captcha.clevergo.dev/.

```shell
$ cd example
$ go run main.go
```

### Quick Start

```shell
$ go get github.com/clevergo/captchas \
	github.com/clevergo/captchas/drivers \
	github.com/clevergo/captchas/stores/memstore
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

	"github.com/clevergo/captchas"
	"github.com/clevergo/captchas/drivers"
	"github.com/clevergo/captchas/stores/memstore"
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
	captcha, err := manager.Generate()
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
	if err := manager.Verify(captchaID, captcha, true); err != nil {
		io.WriteString(w, fmt.Sprintf("captcha is invalid: %s", err.Error()))
		return
	}

	io.WriteString(w, "valid")
}
```
