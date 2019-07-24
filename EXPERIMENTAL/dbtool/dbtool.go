package dbtool

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
)

// Config for database configuration
type Config interface {
	DatabaseName() string
	DataSource() string
	AdminDataSource() string
	DriverName() string
	MigrationSource() string
}

// Tool contain tool of database operation
type Tool struct {
	CreateDatabaseScriptTemplate string
	DropDatabaseScriptTemplate   string
}

// NewPostgresTool return new instance of Tool for Postgres
func NewPostgresTool() *Tool {
	return &Tool{
		CreateDatabaseScriptTemplate: `CREATE DATABASE "%s"`,
		DropDatabaseScriptTemplate:   `DROP DATABASE IF EXISTS "%s"`,
	}
}

// CreateDB is tool to create new database
func (t *Tool) CreateDB(ctx typictx.ActionContext) (err error) {
	return ctx.Typical.Container().Invoke(t.createDB)
}

// DropDB is tool to drop database
func (t *Tool) DropDB(ctx typictx.ActionContext) (err error) {
	return ctx.Typical.Container().Invoke(t.dropDB)
}

// MigrateDB is tool to migrate database
func (t *Tool) MigrateDB(ctx typictx.ActionContext) (err error) {
	return ctx.Typical.Container().Invoke(t.migrateDB)
}

// RollbackDB is tool to rollback database
func (t *Tool) RollbackDB(ctx typictx.ActionContext) (err error) {
	return ctx.Typical.Container().Invoke(t.rollbackDB)
}

func (t *Tool) createDB(config Config) (err error) {
	query := fmt.Sprintf(t.CreateDatabaseScriptTemplate, config.DatabaseName())
	log.Infof(query)

	conn, err := sql.Open(config.DriverName(), config.AdminDataSource())
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec(query)
	return
}

func (t *Tool) dropDB(config Config) (err error) {
	query := fmt.Sprintf(t.DropDatabaseScriptTemplate, config.DatabaseName())
	log.Infof(query)

	conn, err := sql.Open(config.DriverName(), config.AdminDataSource())
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = conn.Exec(query)
	return
}

func (t *Tool) migrateDB(config Config) error {
	sourceURL := fmt.Sprintf("file://%s", config.MigrationSource())
	log.Infof("Migrate database from source '%s'\n", sourceURL)

	migration, err := migrate.New(sourceURL, config.DataSource())
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.Up()
}

func (t *Tool) rollbackDB(config Config) error {
	sourceURL := fmt.Sprintf("file://%s", config.MigrationSource())
	log.Infof("Migrate database from source '%s'\n", sourceURL)

	migration, err := migrate.New(sourceURL, config.DataSource())
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.Down()
}