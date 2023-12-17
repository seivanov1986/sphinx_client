package sphinx_client

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

type sphinx struct {
	db *sqlx.DB
}

func NewSphinx() *sphinx {
	host := os.Getenv("SPHINX_HOST")
	port := os.Getenv("SPHINX_PORT")

	source := fmt.Sprintf(
		"tcp(%v:%v)/",
		host, port,
	)

	conn, err := sqlx.Connect("mysql", source)
	if err != nil {
		//logger.Panic(err)
	} else {
		//conn.SetConnMaxIdleTime(5 * time.Second)
		conn.SetConnMaxLifetime(60 * time.Second)
		conn.SetMaxIdleConns(10)
		conn.SetMaxOpenConns(10)
	}
	return &sphinx{
		db: conn,
	}
}
