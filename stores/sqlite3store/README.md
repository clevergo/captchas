# SQLite3 Store for Captchas

Firstly, you need to create a captcha table, [migration](migrations/1584366110_create_captchas_table.sql).

```shell
$ go get clevergo.tech/captchas/stores/sqlite3store
```

```go
import (
	"clevergo.tech/captchas/stores/dbstore"
	"clevergo.tech/captchas/stores/sqlite3store"
	_ "github.com/mattn/go-sqlite3"
)
```

```go
db, err := sql.Open("sqlite3", "./data.db")
if err != nil {
	// ...
}
store := sqlite3store.New(
	db,
	dbstore.Expiration(10*time.Minute), // captcha expiration, optional.
	dbstore.GCInterval(time.Minute), // garbage collection interval to delete expired captcha, optional.
	dbstore.TableName("captchas"), // table name, optional.
	dbstore.Category("default"), // category, optional.
)
```
