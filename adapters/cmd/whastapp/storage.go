package whastapp

import (
	"context"

	"github.com/LeonardoSola/whatsapp-go-boilerplate/adapters/database/sqlite"

	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func getStorage() *store.Device {
	var container *sqlstore.Container
	dbLog := waLog.Stdout("Database", "ERROR", true)
	ctx := context.Background()

	// Can use other SQL database, like PostgreSQL, MySQL, etc.
	container = sqlstore.NewWithDB(sqlite.SQLiteDB, "sqlite3", dbLog)

	if err := container.Upgrade(ctx); err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	return deviceStore
}
