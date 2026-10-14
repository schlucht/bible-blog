package migrate

const (
	units = `
	CREATE TABLE IF NOT EXISTS units (
		unitId VARCHAR(64) PRIMARY KEY,
		unitName VARCHAR(150) NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	paths = `
	CREATE TABLE IF NOT EXISTS paths (
		pathId VARCHAR(64) PRIMARY KEY,
		unit_id VARCHAR(64) NOT NULL,
		type TEXT NOT NULL,
		number INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(unit_id) REFERENCES units(unitId),
		UNIQUE(unit_id, type, number)
	);
`
	options = `
	CREATE TABLE IF NOT EXISTS options(
		optionId VARCHAR(64) PRIMARY KEY,
		optionName VARCHAR(4) NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	valves = `
	CREATE TABLE IF NOT EXISTS valves(
		valveId VARCHAR(64) PRIMARY KEY,
		position VARCHAR(64),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	path_valves = `
	CREATE TABLE IF NOT EXISTS path_valves (
		path_valve_id VARCHAR(64) PRIMARY KEY,
		path_id VARCHAR(64) NOT NULL,
		option_id VARCHAR(64) NOT NULL,
		valve_id VARCHAR(64) NOT NULL,
		installed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		removed_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(path_id) REFERENCES paths(pathId),
		FOREIGN KEY(option_id) REFERENCES options(optionId),
		FOREIGN KEY(valve_id) REFERENCES valves(valveId)
	);
	`

	idx_paths_unit        = "CREATE INDEX idx_paths_unit ON paths(unit_id);"
	idx_path_valves_valve = "CREATE INDEX idx_path_valves_valve ON path_valves(valve_id);"
)

var Migrations = []Migration{
	{
		Version: "0001_create_units",
		SQL:     units,
	},
	{
		Version: "0002_create_paths",
		SQL:     paths,
	},
	{
		Version: "0003_create_options",
		SQL:     options,
	},
	{
		Version: "0004_create_valves",
		SQL:     valves,
	},
	{
		Version: "0005_create_path_valves",
		SQL:     path_valves,
	},
	{
		Version: "0006_create_idx_paths_unit",
		SQL:     idx_paths_unit,
	},
	{
		Version: "0007_create_idx_path_valves_valve",
		SQL:     idx_path_valves_valve,
	},
}
