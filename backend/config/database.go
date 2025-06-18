package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
	// Get database configuration from viper
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")
	dbSSLMode := viper.GetString("database.sslmode")
	maxOpenConns := viper.GetInt("database.max_open_conns")
	maxIdleConns := viper.GetInt("database.max_idle_conns")
	connMaxLifetime := viper.GetDuration("database.conn_max_lifetime")

	// Create the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	zap.L().Info("Connected to database",
		zap.String("host", dbHost),
		zap.Int("port", dbPort),
		zap.String("dbname", dbName),
		zap.Int("max_open_conns", maxOpenConns),
		zap.Int("max_idle_conns", maxIdleConns),
		zap.Duration("conn_max_lifetime", connMaxLifetime),
	)

	return db, nil
}