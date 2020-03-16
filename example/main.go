package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/clevergo/captchas"
	"github.com/clevergo/captchas/drivers"
	"github.com/clevergo/captchas/stores/memcachedstore"
	"github.com/clevergo/captchas/stores/memstore"
	"github.com/clevergo/captchas/stores/redisstore"
	"github.com/go-redis/redis/v7"
)

var (
	addr      = flag.String("addr", "localhost:8080", "address")
	store     captchas.Store
	managers  map[string]*captchas.Manager
	indexTmpl = template.Must(template.ParseFiles("layout.tmpl", "index.tmpl"))
	apiTmpl   = template.Must(template.ParseFiles("layout.tmpl", "api.tmpl"))
)

func redisStore() *redisstore.Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return redisstore.New(
		client,
		redisstore.Expiration(10*time.Minute), // key expiration, optional.
		redisstore.Prefix("captchas"),         // key prefix, optional.
	)
}

func memcachedStore() *memcachedstore.Store {
	memcachedClient := memcache.New("localhost:11211")
	return memcachedstore.New(
		memcachedClient,
		memcachedstore.Expiration(int32(600)), // captcha expiration, optional.
		memcachedstore.Prefix("captchas"),     // key prefix, optional.
	)
}

func memStore() *memstore.Store {
	return memstore.New(
		memstore.Expiration(10*time.Minute), // expiration, optional.
		memstore.GCInterval(time.Minute),    // garbage collection interval, optional.
	)
}

func main() {
	flag.Parse()

	store = redisStore()     // redis store
	store = memcachedStore() // memcached store
	store = memStore()       // memory store

	managerOpts := []captchas.Option{
		// disable case sensitive, enabled by default, it will effects on string driver.
		captchas.CaseSensitive(false),
	}
	managers = map[string]*captchas.Manager{
		"digit":   captchas.New(store, drivers.NewDigit(), managerOpts...),
		"audio":   captchas.New(store, drivers.NewAudio(), managerOpts...),
		"math":    captchas.New(store, drivers.NewMath(), managerOpts...),
		"string":  captchas.New(store, drivers.NewString(), managerOpts...),
		"chinese": captchas.New(store, drivers.NewChinese(), managerOpts...),
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/api", api)
	http.HandleFunc("/validate", validate)
	http.HandleFunc("/generate", generate)
	log.Println(http.ListenAndServe(*addr, http.DefaultServeMux))
}

func index(w http.ResponseWriter, r *http.Request) {
	manager, err := getManager(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	captcha, err := manager.Generate()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	alert := ""
	valid := false
	if r.Method == http.MethodPost {
		captchaID := r.PostFormValue("captcha_id")
		captchaVal := r.PostFormValue("captcha")
		if err := manager.Verify(captchaID, captchaVal, true); err != nil {
			alert = err.Error()
		} else {
			valid = true
			alert = "captcha is valid"
		}
	}

	render(indexTmpl, w, map[string]interface{}{
		"driver":  r.URL.Query().Get("driver"),
		"captcha": captcha,
		"alert":   alert,
		"valid":   valid,
	})
}

func api(w http.ResponseWriter, r *http.Request) {
	render(apiTmpl, w, map[string]interface{}{
		"driver": r.URL.Query().Get("driver"),
	})
}

func generate(w http.ResponseWriter, r *http.Request) {
	manager, err := getManager(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	captcha, err := manager.Generate()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data, err := json.Marshal(map[string]string{
		"id":   captcha.ID(),
		"data": captcha.EncodeToString(),
	})
	if _, err = w.Write(data); err != nil {
		io.WriteString(w, err.Error())
	}
}

func validate(w http.ResponseWriter, r *http.Request) {
	manager, err := getManager(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	captchaID := r.PostFormValue("captcha_id")
	captchaValue := r.PostFormValue("captcha")
	err = manager.Verify(captchaID, captchaValue, true)
	if err != nil {
		io.WriteString(w, `{"msg":"`+err.Error()+`"}`)
		return
	}
	io.WriteString(w, `{"msg":"success"}`)
}

func getManager(r *http.Request) (*captchas.Manager, error) {
	driver := r.URL.Query().Get("driver")
	if driver == "" {
		driver = "digit"
	}

	if m, ok := managers[driver]; ok {
		return m, nil
	}

	return nil, fmt.Errorf("unsupported driver: %s", driver)
}

func render(tmpl *template.Template, w http.ResponseWriter, data interface{}) {
	if err := tmpl.Execute(w, data); err != nil {
		io.WriteString(w, err.Error())
	}
}
