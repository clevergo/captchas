
# Redis Store for Captchas

```shell
$ go get clevergo.tech/captchas/stores/redisstore
```

```go
import (
    "clevergo.tech/captchas/stores/redisstore"
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
	redisstore.Expiration(10*time.Minute), // captcha expiration, optional.
	redisstore.Prefix("captchas"), // redis key prefix, optional.
)
```