package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
	"golang-employee-crud-example/model"
)

type EmployeeRepository interface {
	Create(ctx context.Context, e *model.Employee) (uuid.UUID, error)
	GetAll(ctx context.Context, limit, offset string) ([]*model.Employee, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Employee, error)
	Update(ctx context.Context, e *model.Employee) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type employeeRepository struct {
	conn *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &employeeRepository{conn: db}
}

func (r *employeeRepository) Create(ctx context.Context, e *model.Employee) (uuid.UUID, error) {
	query := `INSERT INTO employees (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	_, err := r.conn.Exec(ctx, query, e.ID, e.FirstName, e.LastName, e.Email)
	if err != nil {
		return uuid.Nil, err
	}
	return e.ID, nil
}

func (r *employeeRepository) GetAll(ctx context.Context, limit, offset string) ([]*model.Employee, error) {
	query := `SELECT id, first_name, last_name, email FROM employees ORDER BY first_name, last_name LIMIT $1 OFFSET $2`
	rows, err := r.conn.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		var e model.Employee
		err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *employeeRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Employee, error) {
	query := `SELECT id, first_name, last_name, email FROM employees WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	var e model.Employee
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &e, nil
}

func (r *employeeRepository) Update(ctx context.Context, e *model.Employee) error {
	query := `UPDATE employees SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
	res, err := r.conn.Exec(ctx, query, e.FirstName, e.LastName, e.Email, e.ID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (r *employeeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM employees WHERE id = $1`
	res, err := r.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
