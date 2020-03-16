module github.com/clevergo/captchas/example

go 1.14

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/clevergo/captchas v0.3.2
	github.com/clevergo/captchas/drivers v0.3.2
	github.com/clevergo/captchas/stores/memcachedstore v0.0.0-20200316114439-226965818d8e
	github.com/clevergo/captchas/stores/memstore v0.0.0-20200316112430-5701c0af2636
	github.com/clevergo/captchas/stores/redisstore v0.0.0-20200316112430-5701c0af2636
	github.com/go-redis/redis/v7 v7.2.0
	golang.org/x/image v0.0.0-20200119044424-58c23975cae1 // indirect
)
