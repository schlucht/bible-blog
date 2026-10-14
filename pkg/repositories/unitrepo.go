package repositories

import (
	"context"
	"time"
	"uuid"
)

type Unit struct {
	UnitId    uuid.UUID `json:"unitId"`
	UnitName  string    `json:"unitName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func (r *Repo) AllUnits() ([]*Unit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	var units []*Unit

	query := `select unitid, userName, created_at, updated_at from Units`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u Unit
		err := rows.Scan(
			&u.UnitId,
			&u.UnitName,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		units = append(units, &u)
	}

	return units, nil
}

// Speichert eine Neue Anlage in der Datenbank
func (r *Repo) NewUnit(unitName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := `INSERT INTO Units (unitid, userName, created_at, updated_at) VALUES (?,?,?,?)`
	id := uuid.New()
	now := time.Now()

	_, err := r.db.ExecContext(ctx, query,
		id, unitName, now, now,
	)
	if err != nil {
		if r.IsUniqueConstraintError(err) {
			return r.DBError.ErrDuplicateKey
		}
		return err
	}

	return nil
}
