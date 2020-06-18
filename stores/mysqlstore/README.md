# MySQL Store for Captchas

Firstly, you need to create a captcha table, [migration](migrations/1584366110_create_captchas_table.sql).

```shell
$ go get clevergo.tech/captchas/stores/mysqlstore
```

```go
import (
	"clevergo.tech/captchas/stores/dbstore"
	"clevergo.tech/captchas/stores/mysqlstore"
	_ "github.com/go-sql-driver/mysql"
)
```

```go
db, err := sql.Open("mysql", "")
if err != nil {
	// ...
}
store := mysqlstore.New(
	db,
	dbstore.Expiration(10*time.Minute), // captcha expiration, optional.
	dbstore.GCInterval(time.Minute), // garbage collection interval to delete expired captcha, optional.
	dbstore.TableName("captchas"), // table name, optional.
	dbstore.Category("default"), // category, optional.
)
```
