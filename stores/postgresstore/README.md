# PostgreSQL Store for Captchas

Firstly, you need to create a captcha table, [migration](migrations/1584366110_create_captchas_table.sql).

```shell
$ go get github.com/clevergo/captchas/stores/postgresstore
```

```go
import (
	"github.com/clevergo/captchas/stores/dbstore"
	"github.com/clevergo/captchas/stores/postgresstore"
	_ "github.com/lib/pq"
)
```

```go
db, err := sql.Open("postgres", "")
if err != nil {
	// ...
}
store := postgresstore.New(
	db,
	dbstore.Expiration(10*time.Minute), // captcha expiration, optional.
	dbstore.GCInterval(time.Minute), // garbage collection interval to delete expired captcha, optional.
	dbstore.TableName("captchas"), // table name, optional.
	dbstore.Category("default"), // category, optional.
)
```
