package db

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"

	"errors"

	"github.com/jackc/pgx"
	"github.com/hassannmoussaa/bookery/pkg/config"
	"github.com/hassannmoussaa/pill.go/clean"
)

func ConnectToDBs(postgresConfig *config.PostgresDBConfig) *pgx.ConnPool {
	if postgresConfig != nil {
		return GetPostgresConn(postgresConfig.Host, postgresConfig.Port, postgresConfig.User, postgresConfig.Password, postgresConfig.Database, getTlsConfig(postgresConfig.GetCertPath()), postgresConfig.MaxConnections)
	}
	clean.Error(errors.New("Databases configurations was missing!"))
	os.Exit(1)
	return nil
}

func getTlsConfig(certPath string) *tls.Config {
	if certPath != "" {
		roots := x509.NewCertPool()
		if cert, err := ioutil.ReadFile(certPath); err == nil {
			roots.AppendCertsFromPEM(cert)
			return &tls.Config{RootCAs: roots, InsecureSkipVerify: true}
		} else {
			clean.Error(err)
		}
	}
	return nil
}
