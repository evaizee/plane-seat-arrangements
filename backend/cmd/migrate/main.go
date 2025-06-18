package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	// Initialize configuration
	initConfig()

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [up|down] [steps]", os.Args[0])
	}

	command := os.Args[1]
	steps := 0
	if len(os.Args) > 2 {
		fmt.Sscanf(os.Args[2], "%d", &steps)
	}

	m, err := getMigrate()
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	switch command {
	case "up":
		if steps > 0 {
			err = m.Steps(steps)
		} else {
			err = m.Up()
		}
	case "down":
		if steps > 0 {
			err = m.Steps(-steps)
		} else {
			err = m.Down()
		}
	default:
		log.Fatalf("Unknown command: %s", command)
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migration failed: %v", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No migrations to apply")
	} else {
		fmt.Println("Migrations applied successfully")
	}
}

func initConfig() {
	// Set default config path
	configPath := "./config"
	
	// Check if CONFIG_PATH environment variable is set
	if envConfigPath := os.Getenv("CONFIG_PATH"); envConfigPath != "" {
		configPath = envConfigPath
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	// Read environment variables
	viper.AutomaticEnv()

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Failed to read config file: %v", err)
	}
}

func getMigrate() (*migrate.Migrate, error) {
	// Get database configuration from viper
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")
	dbSSLMode := viper.GetString("database.sslmode")

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

	// Create the postgres driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	// Get migrations path
	migrationsPath := "./migrations"
	if envMigrationsPath := os.Getenv("MIGRATIONS_PATH"); envMigrationsPath != "" {
		migrationsPath = envMigrationsPath
	}

	// Ensure path is absolute
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Create the migrate instance
	sourceURL := fmt.Sprintf("file://%s", absPath)
	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return m, nil
}
