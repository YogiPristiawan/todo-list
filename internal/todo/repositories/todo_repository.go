package repositories

import (
	"database/sql"
	"time"
	"todo-list/internal/todo/models"
)

type todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *todo {
	return &todo{
		db: db,
	}
}

func (t *todo) Create(todo *models.Todo) (err error) {
	tx, err := t.db.Begin()
	if err != nil {
		return
	}

	var sql = `
		INSERT INTO todos
		(
			activity_group_id, title, is_active, priority, created_at, updated_at
		)
		VALUES
		(
			?, ?, ?, ?, ?, ?
		)`

	// set to default value
	if todo.Priority == "" {
		todo.Priority = "very-high"
	}

	if !todo.IsActive.Valid {
		todo.IsActive.Bool = true
		todo.IsActive.Valid = true
	}

	// create timestamps
	now := time.Now().Unix()

	_, err = tx.Exec(sql,
		todo.ActivityGroupId, todo.Title, todo.IsActive,
		todo.Priority, now, now)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return
	}

	sql = `
		SELECT
			id, activity_group_id, title, is_active, priority,
			created_at, updated_at
		FROM
			todos
		ORDER BY
			id DESC LIMIT 1`

	err = tx.QueryRow(sql).Scan(
		&todo.Id, &todo.ActivityGroupId, &todo.Title, &todo.IsActive,
		&todo.Priority, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return
	}

	err = tx.Commit()
	return
}

func (t *todo) FindAll(activityGroupId int64) (todos []models.Todo, err error) {
	var sql = `
		SELECT
			id, activity_group_id, title, is_active, priority,
			created_at, updated_at, deleted_at
		FROM
			todos
		WHERE
			deleted_at IS NULL`
	var binds = []interface{}{}

	// filter
	if activityGroupId != 0 {
		sql += ` AND activity_group_id = ?`
		binds = append(binds, activityGroupId)
	}

	sql += ` ORDER BY created_at DESC`

	rows, err := t.db.Query(sql, binds...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Todo
		err = rows.Scan(
			&t.Id, &t.ActivityGroupId, &t.Title, &t.IsActive,
			&t.Priority, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return
		}

		todos = append(todos, t)
	}
	return
}

func (t *todo) Detail(id int64) (todo models.Todo, err error) {
	var sql = `
	SELECT
		id, activity_group_id, title, is_active, priority,
		created_at, updated_at, deleted_at
	FROM
		todos
	WHERE
		id = ?
		AND
		deleted_at IS NULL
	ORDER BY
		created_at DESC`

	err = t.db.QueryRow(sql, id).Scan(
		&todo.Id, &todo.ActivityGroupId, &todo.Title,
		&todo.IsActive, &todo.Priority, &todo.CreatedAt,
		&todo.UpdatedAt, &todo.DeletedAt)
	return
}

func (t *todo) Update(id int64, todo *models.Todo) (affected int64, err error) {
	// get todo
	var sql = `
		SELECT
			id, activity_group_id, title, is_active, priority,
			created_at, updated_at, deleted_at
		FROM
			todos
		WHERE
			id = ?
			AND
			deleted_at IS NULL`

	var td models.Todo
	err = t.db.QueryRow(sql, id).Scan(
		&td.Id, &td.ActivityGroupId, &td.Title, &td.IsActive, &td.Priority,
		&td.CreatedAt, &td.UpdatedAt, &td.DeletedAt)
	if err != nil {
		return
	}

	sql = `
		UPDATE todos
		SET
			updated_at = ?`

	// create timestamps
	now := time.Now().Unix()
	var binds = []interface{}{now}

	// optional update
	if todo.Title != "" {
		sql += `,title = ?`
		td.Title = todo.Title
		binds = append(binds, todo.Title)
	}

	if todo.ActivityGroupId != 0 {
		sql += `,activity_group_id = ?`
		td.ActivityGroupId = todo.ActivityGroupId
		binds = append(binds, todo.ActivityGroupId)
	}

	if todo.Priority != "" {
		sql += `,priority = ?`
		td.Priority = todo.Priority
		binds = append(binds, todo.Priority)
	}

	if todo.IsActive.Valid {
		sql += `,is_active = ?`
		td.IsActive = todo.IsActive
		binds = append(binds, todo.IsActive)
	}

	binds = append(binds, id)

	td.UpdatedAt = now
	*todo = td

	sql += ` WHERE id = ? AND deleted_at IS NULL`

	res, err := t.db.Exec(sql, binds...)
	if err != nil {
		return
	}

	return res.RowsAffected()
}

func (t *todo) Delete(id int64) (affected int64, err error) {
	var sql = `
		UPDATE todos
		SET
			deleted_at = ?
		WHERE
			id = ?
			AND
			deleted_at IS NULL`

	// create timestamps
	now := time.Now().Unix()

	res, err := t.db.Exec(sql, now, id)
	if err != nil {
		return
	}

	return res.RowsAffected()
}
