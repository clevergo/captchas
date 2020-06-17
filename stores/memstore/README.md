# Memory Store for Captchas

```shell
$ go get clevergo.tech/captchas/stores/memstore
```

```go
store := memstore.New(
	memstore.Expiration(10*time.Minute), // captcha expiration, optional.
	memstore.GCInterval(time.Minute), // garbage collection interval to delete expired captcha, optional.
)
```

> Inspired by [scs.memstore](https://github.com/alexedwards/scs/tree/master/memstore).
