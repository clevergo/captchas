# Memcached Store for Captchas

```shell
$ go get github.com/clevergo/captchas/stores/memcachedstore
```

```go
import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/clevergo/captchas/stores/memcachedstore"
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
