# Captchas
[![Build Status](https://travis-ci.org/clevergo/captchas.svg?branch=master)](https://travis-ci.org/clevergo/captchas)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/captchas/badge.svg?branch=master)](https://coveralls.io/github/clevergo/captchas?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/captchas)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/captchas)](https://goreportcard.com/report/github.com/clevergo/captchas)
[![Release](https://img.shields.io/github/release/clevergo/captchas.svg?style=flat-square)](https://github.com/clevergo/captchas/releases)

## Usage

```shell
$ cd example
$ go run main.go
```

## Drivers

- digit
- audio
- math
- string
- chinese

## Stores

- memstore: memory store.
- redisstore: redis store.
- add your store here by PR or [request a new store](https://github.com/clevergo/captchas/issues/new).

