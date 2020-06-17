# Memcached Store for Captchas

```shell
$ go get clevergo.tech/captchas/stores/memcachedstore
```

```go
import (
	"clevergo.tech/captchas/stores/memcachedstore"
	"github.com/bradfitz/gomemcache/memcache"
)
```

```go
// client.
client := memcache.New("localhost:11211")
store := memcachedstore.New(
	client,
	memcachedstore.Expiration(int32(600)), // captcha expiration, optional.
	memcachedstore.Prefix("captchas"),     // key prefix, optional.
)
```
