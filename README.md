# Captchas
[![Build Status](https://travis-ci.org/clevergo/captchas.svg?branch=master)](https://travis-ci.org/clevergo/captchas)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/captchas/badge.svg?branch=master)](https://coveralls.io/github/clevergo/captchas?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/captchas)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/captchas)](https://goreportcard.com/report/github.com/clevergo/captchas)
[![Release](https://img.shields.io/github/release/clevergo/captchas.svg?style=flat-square)](https://github.com/clevergo/captchas/releases)

Base64 Captchas Manager, supports multiple [drivers](#drivers) and [stores](#stores).

## Usage

```shell
$ cd example
$ go run main.go
```

## Drivers

```go
import "github.com/clevergo/captchas/drivers"
```

### Digit

```go
// all options are optianal.
opts := []drivers.DigitOption{
	drivers.DigitHeight(50),
	drivers.DigitWidth(120),
	drivers.DigitLength(6),
	drivers.DigitMaxSkew(0.8),
	drivers.DigitDotCount(80),
}
driver := drivers.NewDigit(opts...)
```

### Audio

```go
// all options are optianal.
opts := []drivers.AudioOption{
	drivers.AudioLangauge("en"),
	drivers.AudioLength(6),
}
driver := drivers.NewAudio(opts...)
```

### Math

```go
// all options are optianal.
opts := []drivers.MathOption{
	drivers.MathHeight(50),
	drivers.MathWidth(120),
	drivers.MathNoiseCount(0),
	drivers.MathFonts([]string{}),
	drivers.MathBGColor(&color.RGBA{}),
}
driver := drivers.NewMath(opts...)
```

### String

```go
// all options are optianal.
opts := []drivers.StringOption{
	drivers.StringHeight(50),
	drivers.StringWidth(120),
	drivers.StringLength(4),
	drivers.StringNoiseCount(0),
	drivers.StringFonts([]string{}),
	drivers.StringSource("abcdefghijklmnopqrstuvwxyz"),
	drivers.StringBGColor(&color.RGBA{}),
}
driver := drivers.NewString(opts...)
```

### Chinese

```go
// all options are optianal.
opts := []drivers.ChineseOption{
	drivers.ChineseHeight(50),
	drivers.ChineseWidth(120),
	drivers.ChineseLength(4),
	drivers.ChineseNoiseCount(0),
	drivers.ChineseFonts([]string{"wqy-microhei.ttc"}),
	drivers.ChineseSource("零一二三四五六七八九十"),
	drivers.ChineseBGColor(&color.RGBA{}),
}
driver := drivers.NewChinese(opts...)
```

## Stores

### Memory Store

```go
import "github.com/clevergo/captchas/memstore"
```

```go
expiration := 10 * time.Minute // captcha expiration.
gcInterval := time.Minute // garbage collection interval to delete expired captcha.
store := memstore.New(expiration, gcInterval)
```

### Redis Store

```go
import (
    "github.com/clevergo/captchas/redisstore"
    "github.com/go-redis/redis/v7"
)
```

```go
// redis client.
client := redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
prefix := "captcha"            // redis key prefix, optional.
expiration := 10 * time.Minute // captcha expiration.
store := redisstore.New(client, expiration, redisstore.Prefix(prefix))
```

### More

Add your store here by PR or [request a new store](https://github.com/clevergo/captchas/issues/new).

