package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
	sqlDB *sql.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection creates a new database connection
func NewConnection(config Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &DB{
		DB:    gormDB,
		sqlDB: sqlDB,
	}, nil
}

// Ping checks if the database connection is alive
func (db *DB) Ping(ctx context.Context) error {
	return db.sqlDB.PingContext(ctx)
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.sqlDB.Close()
}

// BeginTx starts a transaction
func (db *DB) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return db.sqlDB.BeginTx(ctx, nil)
}

// CommitTx commits a transaction
func (db *DB) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

// RollbackTx rolls back a transaction
func (db *DB) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}