package migrate


type Migrator interface {
	Up() error
}

type Migration struct {
	Version string
	SQL string
}

const Migrate_table = `
	CREATE TABLE IF NOT EXISTS migrations(
		version VARCHAR(150) PRIMARY KEY,
		applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
`

