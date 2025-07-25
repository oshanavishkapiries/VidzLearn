package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/Cenzios/pf-backend/pkg/db/dbiface"
	"github.com/Cenzios/pf-backend/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlImpl struct {
	conn *sql.DB
}

// New returns a new mysqlImpl as dbiface.Database
func New() dbiface.Database {
	dsn := os.Getenv("MYSQL_DSN") // e.g., "user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true"
	if dsn == "" {
		logger.Error.Fatalln("❌ MYSQL_DSN not set")
	}

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error.Fatalln("❌ Failed to connect to MySQL:", err)
	}

	if err := conn.Ping(); err != nil {
		logger.Error.Fatalln("❌ MySQL ping failed:", err)
	}

	logger.Info.Println("✅ MySQL connected")
	return &mysqlImpl{conn: conn}
}

func (m *mysqlImpl) FindOne(ctx context.Context, table string, filter interface{}) (interface{}, error) {
	// Example stub (not secure, real implementation should use prepared statements & query builder)
	return nil, fmt.Errorf("FindOne not implemented for MySQL")
}

func (m *mysqlImpl) FindMany(ctx context.Context, table string, filter interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("FindMany not implemented for MySQL")
}

func (m *mysqlImpl) InsertOne(ctx context.Context, table string, data interface{}) error {
	return fmt.Errorf("InsertOne not implemented for MySQL")
}

func (m *mysqlImpl) UpdateOne(ctx context.Context, table string, filter interface{}, update interface{}) error {
	return fmt.Errorf("UpdateOne not implemented for MySQL")
}

func (m *mysqlImpl) DeleteOne(ctx context.Context, table string, filter interface{}) error {
	return fmt.Errorf("DeleteOne not implemented for MySQL")
}

func (m *mysqlImpl) Aggregate(ctx context.Context, table string, pipeline interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("Aggregate not supported in SQL")
}
