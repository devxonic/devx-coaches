package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	api "github.com/ibad69/go-devxcoaches/src/pkg"
	"github.com/ibad69/go-devxcoaches/src/pkg/index"
)

type Db struct {
	conn *pgxpool.Pool
}

func New(db *pgxpool.Pool) index.Repo {
	return &Db{conn: db}
}

func (db *Db) AddStudent(ctx context.Context, student api.Student) (api.Student, error) {
	query := `SELECT COUNT(*) FROM students`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	log.Println(count)
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO students(id, name, gender, s_of, email, class_id, batch_id, subject_id, fees_amount, phone, caddress, adm_fee, nic_no, status)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        `
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		student.Name,
		student.Gender,
		student.S_of,
		student.Email,
		student.Class_id,
		student.Batch_id,
		student.Subject_id,
		student.FeesAmount,
		student.Phone,
		student.CAddress,
		student.Adm_fee,
		student.Nic_no,
		student.Status,
	)
	if err != nil {
		return api.Student{}, err
	}
	return student, nil
}

func (db *Db) AddClass(ctx context.Context, class api.Class) (api.Class, error) {
	query := `SELECT COUNT(*) FROM class`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	if err != nil {
		return api.Class{}, fmt.Errorf("unable to query users: %w", err)
	}
	// pgx.CollectOneRow(rows, pgx.RowToStructByName(count))
	defer rows.Close()
	// rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO class(id, description, status)
    VALUES($1, $2, $3)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		class.Description,
		class.Status,
	)
	if err != nil {
		return api.Class{}, err
	}
	return class, nil
}

func (db *Db) AddSubject(ctx context.Context, subject api.Subject) (api.Subject, error) {
	query := `SELECT COUNT(*) FROM subject`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	if err != nil {
		return api.Subject{}, fmt.Errorf("unable to query users: %w", err)
	}
	// pgx.CollectOneRow(rows, pgx.RowToStructByName(count))
	defer rows.Close()
	// rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO subject(id, description, status)
    VALUES($1, $2, $3)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		subject.Description,
		subject.Status,
	)
	if err != nil {
		return api.Subject{}, err
	}
	return subject, nil
}

func (db *Db) AddBatch(ctx context.Context, Batch api.Batch) (api.Batch, error) {
	query := `SELECT COUNT(*) FROM batches`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	if err != nil {
		return api.Batch{}, fmt.Errorf("unable to query users: %w", err)
	}
	// pgx.CollectOneRow(rows, pgx.RowToStructByName(count))
	defer rows.Close()
	// rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO batches(id, description, status)
    VALUES($1, $2, $3)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		Batch.Description,
		Batch.Status,
	)
	if err != nil {
		return api.Batch{}, err
	}
	return Batch, nil
}

func (db *Db) AddYear(ctx context.Context, Year api.Year) (api.Year, error) {
	query := `SELECT COUNT(*) FROM Year`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	if err != nil {
		return api.Year{}, fmt.Errorf("unable to query users: %w", err)
	}
	// pgx.CollectOneRow(rows, pgx.RowToStructByName(count))
	defer rows.Close()
	// rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO Year(id, description, status)
    VALUES($1, $2, $3)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		Year.Description,
		Year.Status,
	)
	if err != nil {
		return api.Year{}, err
	}
	return Year, nil
}

func (db *Db) AddPeriods(ctx context.Context, Period api.Period) (api.Period, error) {
	query := `SELECT COUNT(*) FROM Year`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	if err != nil {
		return api.Period{}, fmt.Errorf("unable to query users: %w", err)
	}
	// pgx.CollectOneRow(rows, pgx.RowToStructByName(count))
	defer rows.Close()
	// rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO Year(id, description, status)
    VALUES($1, $2, $3)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		// Period.Description,
		Period.Status,
	)
	if err != nil {
		return api.Period{}, err
	}
	return Period, nil
}

func (db *Db) GetStudents(ctx context.Context) ([]api.Student, error) {
	q := ` SELECT * FROM students`
	rows, err := db.conn.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.Student])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}

func (db *Db) GetClass(ctx context.Context) ([]api.Class, error) {
	q := ` SELECT * FROM class`
	rows, err := db.conn.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.Class])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}

func (db *Db) GetSubjects(ctx context.Context) ([]api.Subject, error) {
	q := ` SELECT id, description, status FROM subject`
	rows, err := db.conn.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.Subject])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}

func (db *Db) GetBatches(ctx context.Context) ([]api.Batch, error) {
	q := ` SELECT id, description, status FROM batches`
	rows, err := db.conn.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.Batch])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}
