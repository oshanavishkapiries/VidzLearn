package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/Cenzios/pf-backend/pkg/db/dbiface"
	"github.com/Cenzios/pf-backend/pkg/logger"
	_ "github.com/lib/pq"
)

type postgresImpl struct {
	conn *sql.DB
}

var conn *sql.DB

// New returns a new postgresImpl as dbiface.Database
func New() dbiface.Database {
	dsn := os.Getenv("POSTGRES_DSN") // e.g., "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
	if dsn == "" {
		logger.Error.Fatalln("❌ POSTGRES_DSN not set")
	}

	var err error
	conn, err = sql.Open("postgres", dsn)
	if err != nil {
		logger.Error.Fatalln("❌ Failed to connect to Postgres:", err)
	}

	if err = conn.Ping(); err != nil {
		logger.Error.Fatalln("❌ Postgres ping failed:", err)
	}

	logger.Info.Println("✅ PostgreSQL connected")
	return &postgresImpl{conn: conn}
}

func (pg *postgresImpl) FindOne(ctx context.Context, table string, filter interface{}) (interface{}, error) {
	// Not implemented: map[filter] to query string safely
	return nil, fmt.Errorf("FindOne not implemented for Postgres")
}

func (pg *postgresImpl) FindMany(ctx context.Context, table string, filter interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("FindMany not implemented for Postgres")
}

func (pg *postgresImpl) InsertOne(ctx context.Context, table string, data interface{}) error {
	return fmt.Errorf("InsertOne not implemented for Postgres")
}

func (pg *postgresImpl) UpdateOne(ctx context.Context, table string, filter interface{}, update interface{}) error {
	return fmt.Errorf("UpdateOne not implemented for Postgres")
}

func (pg *postgresImpl) DeleteOne(ctx context.Context, table string, filter interface{}) error {
	return fmt.Errorf("DeleteOne not implemented for Postgres")
}

func (pg *postgresImpl) Aggregate(ctx context.Context, table string, pipeline interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("Aggregate not supported in SQL")
}
