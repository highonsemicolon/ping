package db

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/highonsemicolon/ping/user-service/pkg/utils"
)

func setupTLSConfig(CACertPath string) {
	caCert, err := os.ReadFile(CACertPath)
	if err != nil {
		log.Fatalf("Failed to read CA certificate file: %v", err)
	}

	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA certificate")
	}
	tlsConfig := &tls.Config{
		RootCAs: rootCertPool,
	}
	mysql.RegisterTLSConfig("custom", tlsConfig)
}

func InitMySQL(config utils.MySQLConfig) (*sql.DB, error) {
	setupTLSConfig(config.CACertPath)

	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
