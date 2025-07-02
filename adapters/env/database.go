package env

import (
	"os"
)

var (
	SQLiteDB string
)

func init() {
	SQLiteDB = os.Getenv("SQLITE_DB")
}
