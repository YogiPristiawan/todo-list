package migrations

import (
	"database/sql"
)

func Migrate(db *sql.DB) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// CREATE TABLE activites
	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS activities (
			id 			BIGINT NOT NULL AUTO_INCREMENT,
			title 		VARCHAR(255) NOT NULL,
			email 		VARCHAR(255) NOT NULL,
			created_at 	BIGINT NOT NULL,
			updated_at 	BIGINT NOT NULL,
			deleted_at 	BIGINT NULL,

			PRIMARY KEY (id)
		)
	`)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return
	}

	// CREATE TABLE todos
	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id 					BIGINT NOT NULL AUTO_INCREMENT,
			activity_group_id 	BIGINT NOT NULL,
			title				VARCHAR(255) NOT NULL,
			is_active			BOOLEAN DEFAULT true,
			priority			ENUM('very-low', 'low', 'high', 'very-high') NOT NULL DEFAULT 'very-high',
			created_at			BIGINT NOT NULL,
			updated_at			BIGINT NOT NULL,
			deleted_at			BIGINT NULL,

			PRIMARY KEY (id),
			INDEX (activity_group_id)
		)
	`)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return
	}

	return tx.Commit()
}
