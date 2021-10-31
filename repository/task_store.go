package repository

import (
	"context"
	"github.com/SemmiDev/Go-Scheduled/entity"
	"time"
)

const createTask = `
INSERT INTO task (
    id, 
    title, 
    description,
    is_done,
    is_over,
    duration,
    matrix,
    start_at,
    due,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING 
	id, 
    title, 
    description,
    is_done,
    is_over,
    duration,
    matrix,
    start_at,
    due,
    created_at,
    updated_at
`

func (q *Queries) Save(ctx context.Context, arg *entity.Task) (*entity.Task, error) {
	row := q.db.QueryRow(ctx, createTask,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.IsDone,
		arg.IsOver,
		arg.Duration,
		arg.Matrix,
		arg.StartAt,
		arg.Due,
		arg.CreatedAt,
		arg.UpdatedAt,
	)

	var i entity.Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.IsOver,
		&i.Duration,
		&i.Matrix,
		&i.StartAt,
		&i.Due,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	return &i, err
}

const getTask = `
SELECT
	id, 
    title, 
    description,
    is_done,
    is_over,
    duration,
    matrix,
    start_at,
    due,
    created_at,
    updated_at
FROM task
WHERE id = $1
`

func (q *Queries) FindById(ctx context.Context, ID string) (*entity.Task, error) {
	row := q.db.QueryRow(ctx, getTask, ID)

	var i entity.Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.IsOver,
		&i.Duration,
		&i.Matrix,
		&i.StartAt,
		&i.Due,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const listTasks = `
SELECT
	id, 
    title, 
    description,
    is_done,
    is_over,
    duration,
    matrix,
    start_at,
    due,
    created_at,
    updated_at
FROM task
`

func (q *Queries) FindAll(ctx context.Context) ([]*entity.Task, error) {
	rows, err := q.db.Query(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*entity.Task
	for rows.Next() {
		var i entity.Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsDone,
			&i.IsOver,
			&i.Duration,
			&i.Matrix,
			&i.StartAt,
			&i.Due,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, &i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const updateTask = `
UPDATE task SET
	title = $2, 
    description = $3,
    is_done = $4,
    is_over = $5,
    matrix = $6,
	duration = $7,
    start_at = $8,
    due = $9,
    updated_at = $10
WHERE id = $1
RETURNING 
	id, 
    title, 
    description,
    is_done,
    is_over,
    duration,
    matrix,
    start_at,
    due,
    created_at,
    updated_at
`

func (q *Queries) Update(ctx context.Context, arg *entity.Task) (*entity.Task, error) {
	row := q.db.QueryRow(ctx, updateTask,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.IsDone,
		arg.IsOver,
		arg.Matrix,
		arg.Duration,
		arg.StartAt,
		arg.Due,
		arg.UpdatedAt,
	)

	var i entity.Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.IsOver,
		&i.Duration,
		&i.Matrix,
		&i.StartAt,
		&i.Due,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateIsDone = `
UPDATE task SET
    is_done = $2,
    updated_at = $3
WHERE id = $1
RETURNING 
	id, 
    title, 
    description,
    is_done,
    is_over,
    duration,
    matrix,
    start_at,
    due,
    created_at,
    updated_at
`

func (q *Queries) UpdateIsDone(ctx context.Context, ID string) (*entity.Task, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowInLoc := time.Now().In(loc).Unix()

	row := q.db.QueryRow(ctx, updateIsDone,
		ID,
		true,
		nowInLoc,
	)

	var i entity.Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.IsOver,
		&i.Duration,
		&i.Matrix,
		&i.StartAt,
		&i.Due,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const deleteTask = `
DELETE FROM task WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, ID string) error {
	_, err := q.db.Exec(ctx, deleteTask, ID)
	return err
}
