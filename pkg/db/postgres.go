package db

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

//TABLES NAMES
const (
	dbSchema         string = "public"
	dbSchemaJoiner   string = "."
	AdminTable       string = dbSchema + dbSchemaJoiner + "admin"
	UserTable        string = dbSchema + dbSchemaJoiner + "user"
	CategoryTable    string = dbSchema + dbSchemaJoiner + "category"
	BookTable        string = dbSchema + dbSchemaJoiner + "book"
	TransactionTable string = dbSchema + dbSchemaJoiner + "transaction"
	OrderTable       string = dbSchema + dbSchemaJoiner + "order"
	UserBookTable    string = dbSchema + dbSchemaJoiner + "user_book"
	CardOrderTable   string = dbSchema + dbSchemaJoiner + "order_card"
)

//It Connect to database and return the connection
func GetPostgresConn(dbHost string, dbPort uint16, dbUser string, dbPassword string, dbName string, tlsConfig *tls.Config, maxConnections int) *pgx.ConnPool {
	if maxConnections == 0 {
		maxConnections = 100
	}
	connConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:      dbHost,
			Port:      dbPort,
			User:      dbUser,
			Password:  dbPassword,
			Database:  dbName,
			TLSConfig: tlsConfig,
		},
		MaxConnections: maxConnections,
	}
	conn, err := pgx.NewConnPool(connConfig)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return conn
}
