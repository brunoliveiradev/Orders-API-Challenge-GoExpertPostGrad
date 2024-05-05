package database

import (
	"GoExpertPostGrad-Orders-Challenge/configs"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func SetupDatabase(config *configs.Envs) (*sql.DB, error) {
	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.DBUser, config.DBPassword, config.DBHost, config.DBPort))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(config.DBMaxOpenConns)
	db.SetMaxIdleConns(config.DBMaxIdleConns)

	// Check database connectivity
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if !tableExists(db, config.DBName) {
		if err := runMigrations(db, config.DBName); err != nil {
			return nil, err
		}
	}

	_, err = db.Exec("USE " + config.DBName)
	if err != nil {
		return nil, fmt.Errorf("failed to switch database: %w", err)
	}

	return db, nil
}

func tableExists(db *sql.DB, tableName string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = ?)", tableName).Scan(&exists)
	if err != nil {
		log.Fatalf("failed to check if table exists: %v", err)
	}
	return exists
}

func runMigrations(db *sql.DB, dbName string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("migration driver setup failed: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", dbName, driver)
	if err != nil {
		return fmt.Errorf("migration setup failed: %w", err)
	}

	err = m.Migrate(1)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}
