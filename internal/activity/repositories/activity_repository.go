package repositories

import (
	"database/sql"
	"time"
	"todo-list/internal/activity/models"
)

type activity struct {
	db *sql.DB
}

func NewActivity(
	db *sql.DB,
) *activity {
	return &activity{
		db: db,
	}
}

func (a *activity) Create(activity *models.Activity) (err error) {
	tx, err := a.db.Begin()
	if err != nil {
		return
	}

	var sql = `
		INSERT INTO activities
		(
			title, email, created_at, updated_at
		)
		VALUES
		(
			?, ?, ?, ?
		)`

	// create timestamps
	now := time.Now().Unix()
	_, err = tx.Exec(sql,
		activity.Title, activity.Email, now, now,
	)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return
	}

	sql = `
		SELECT
			id, title, email , created_at, updated_at
		FROM
			activities
		ORDER BY id DESC LIMIT 1`

	err = tx.QueryRow(sql).Scan(
		&activity.Id, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt,
	)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return
	}

	err = tx.Commit()
	return
}

func (a *activity) FindAll() (activities []models.Activity, err error) {
	var sql = `
		SELECT
			id, email, title, created_at, updated_at, deleted_at
		FROM
			activities
		WHERE
			deleted_at IS NULL
		ORDER BY
			created_at DESC`

	rows, err := a.db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Activity
		err = rows.Scan(
			&a.Id, &a.Email, &a.Title, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt,
		)
		if err != nil {
			return
		}

		activities = append(activities, a)
	}
	return
}

func (a *activity) GetById(id int64) (activity models.Activity, err error) {
	var sql = `
		SELECT
			id, email, title, created_at, updated_at, deleted_at
		FROM
			activities
		WHERE
			id = ?
			AND
			deleted_at IS NULL`

	err = a.db.QueryRow(sql, id).Scan(
		&activity.Id, &activity.Email, &activity.Title,
		&activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt,
	)
	return
}

func (a *activity) Delete(id int64) (affected int64, serr error) {
	var sql = `
		UPDATE
			activities
		SET
			deleted_at = ?
		WHERE
			id = ?
			AND
			deleted_at IS NULL`

	// create timestamps
	now := time.Now().Unix()

	res, err := a.db.Exec(sql, now, id)
	if err != nil {
		return
	}
	return res.RowsAffected()
}

func (a *activity) Update(id int64, activity *models.Activity) (affected int64, err error) {
	var ac models.Activity
	var getSql = `
		SELECT
			id, title, email, created_at, updated_at
		FROM
			activities
		WHERE
			id = ?
			AND
			deleted_at IS NULL`
	err = a.db.QueryRow(getSql, id).Scan(
		&ac.Id, &ac.Title, &ac.Email, &ac.CreatedAt, &ac.UpdatedAt)
	if err != nil {
		return
	}

	var updateSql = `
	UPDATE
		activities
	SET
		title = ?,
		updated_at = ?`

	// create timestamps
	now := time.Now().Unix()
	var binds = []interface{}{activity.Title, now}

	// optional update
	if activity.Email != "" {
		updateSql += `,email = ?`
		ac.Email = activity.Email
		binds = append(binds, activity.Email)
	}

	binds = append(binds, now, id)

	// pass updated data into model's pointer
	ac.Title = activity.Title
	ac.UpdatedAt = now
	*activity = ac

	updateSql += `
		WHERE
			id = ?
			AND 
			deleted_at IS NULL`

	res, err := a.db.Exec(updateSql, binds...)
	if err != nil {
		return
	}

	return res.RowsAffected()
}
