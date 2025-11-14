package postgres

import (
	// "database/sql"

	"context"
	"log"
	"sync"

	"github.com/ahsansaif47/cdc-app/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Connection *pgx.Conn
	Pool       *pgxpool.Conn
}

var dbInstance *Database
var dbOnce sync.Once

func GetDatabaseConnection() *Database {
	dbOnce.Do(func() {
		dbInstance = &Database{
			Connection: connect(),
		}
	})

	return dbInstance
}

func connect() *pgx.Conn {
	c := config.GetConfig()

	conn, err := pgx.Connect(context.Background(), c.DBUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return conn
}
