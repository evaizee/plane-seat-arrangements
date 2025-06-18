package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations to set up or update the database schema.`,
}

// migrateUpCmd represents the migrate up command
var migrateUpCmd = &cobra.Command{
	Use:   "up [steps]",
	Short: "Run migrations up",
	Long:  `Run migrations up to apply schema changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		steps := 0
		if len(args) > 0 {
			fmt.Sscanf(args[0], "%d", &steps)
		}

		m, err := getMigrate()
		if err != nil {
			log.Fatalf("Failed to create migrate instance: %v", err)
		}

		if steps > 0 {
			err = m.Steps(steps)
		} else {
			err = m.Up()
		}

		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migration failed: %v", err)
		}

		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
		} else {
			fmt.Println("Migrations applied successfully")
		}
	},
}

// migrateDownCmd represents the migrate down command
var migrateDownCmd = &cobra.Command{
	Use:   "down [steps]",
	Short: "Run migrations down",
	Long:  `Run migrations down to revert schema changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		steps := 0
		if len(args) > 0 {
			fmt.Sscanf(args[0], "%d", &steps)
		}

		m, err := getMigrate()
		if err != nil {
			log.Fatalf("Failed to create migrate instance: %v", err)
		}

		if steps > 0 {
			err = m.Steps(-steps)
		} else {
			err = m.Down()
		}

		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Migration failed: %v", err)
		}

		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
		} else {
			fmt.Println("Migrations applied successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
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
	migrationsPath := "./db/migrations"
	if envMigrationsPath := os.Getenv("MIGRATIONS_PATH"); envMigrationsPath != "" {
		migrationsPath = envMigrationsPath
	}

	// Create the migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://" + migrationsPath,
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return m, nil
}