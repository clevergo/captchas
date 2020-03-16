
# Redis Store for Captchas

```shell
$ go get github.com/clevergo/captchas/stores/redisstore
```

```go
import (
    "github.com/clevergo/captchas/stores/redisstore"
    "github.com/go-redis/redis/v7"
)
```

```go
// redis client.
client := redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
store := redisstore.New(
	client,
	redisstore.Expiration(expiration), // captcha expiration, optional.
	redisstore.Prefix("caotchas"), // redis key prefix, optional.
)
```