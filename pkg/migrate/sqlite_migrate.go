package migrate

import (
	"context"
	"database/sql"
	"time"
)

type SQLiteMigrator struct {
	Db         *sql.DB
	Migrations []Migration
}

func NewSQLiteMigrator(db *sql.DB, commands []Migration) *SQLiteMigrator {
	db.Exec(Migrate_table)
	return &SQLiteMigrator{
		Db:         db,
		Migrations: commands,
	}
}

func (m *SQLiteMigrator) Up() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	tx, err := m.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, stmt := range m.Migrations {
		var count int
		err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM migrations WHERE version=?", stmt.Version).Scan(&count)		
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		if _, err := tx.Exec(stmt.SQL); err != nil {
			return err
		}
		if _, err := tx.Exec("INSERT INTO migrations(version) VALUES (?)", stmt.Version); err != nil {
			return err
		}

	}
	return tx.Commit()
}
