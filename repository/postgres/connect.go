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
	Pool       *pgxpool.Pool
}

var dbInstance *Database
var dbOnce sync.Once

func GetDatabaseConnection() *Database {
	dbOnce.Do(func() {
		dbInstance = connect()
	})

	return dbInstance
}

func connect() *Database {
	c := config.GetConfig()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, c.DBUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbPool, err := pgxpool.New(ctx, c.DBUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return &Database{
		Connection: conn,
		Pool:       dbPool,
	}
}
