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
	if err != nil {
		return api.Student{}, fmt.Errorf("unable to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return api.Student{}, fmt.Errorf("unable to scan row: %w", err)
		}
	} else {
		return api.Student{}, fmt.Errorf("no rows returned")
	}

	if err := rows.Err(); err != nil {
		return api.Student{}, fmt.Errorf("error iterating over rows: %w", err)
	}
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
	if err != nil {
		return api.Class{}, fmt.Errorf("unable to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return api.Class{}, fmt.Errorf("unable to scan row: %w", err)
		}
	} else {
		return api.Class{}, fmt.Errorf("no rows returned")
	}

	if err := rows.Err(); err != nil {
		return api.Class{}, fmt.Errorf("error iterating over rows: %w", err)
	}

	log.Println("Count:", count) // rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
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
	if err != nil {
		return api.Subject{}, fmt.Errorf("unable to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return api.Subject{}, fmt.Errorf("unable to scan row: %w", err)
		}
	} else {
		return api.Subject{}, fmt.Errorf("no rows returned")
	}

	if err := rows.Err(); err != nil {
		return api.Subject{}, fmt.Errorf("error iterating over rows: %w", err)
	}

	log.Println("Count:", count) // rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
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
	if err != nil {
		return api.Batch{}, fmt.Errorf("unable to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return api.Batch{}, fmt.Errorf("unable to scan row: %w", err)
		}
	} else {
		return api.Batch{}, fmt.Errorf("no rows returned")
	}

	if err := rows.Err(); err != nil {
		return api.Batch{}, fmt.Errorf("error iterating over rows: %w", err)
	}

	log.Println("Count:", count) // rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
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
	if err != nil {
		return api.Year{}, fmt.Errorf("unable to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return api.Year{}, fmt.Errorf("unable to scan row: %w", err)
		}
	} else {
		return api.Year{}, fmt.Errorf("no rows returned")
	}

	if err := rows.Err(); err != nil {
		return api.Year{}, fmt.Errorf("error iterating over rows: %w", err)
	}

	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO Year(id, description, status, remarks)
    VALUES($1, $2, $3, $4)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		Year.Description,
		Year.Status,
		Year.Remarks,
	)
	if err != nil {
		fmt.Print("error in adding year")
		fmt.Print(count)
		return api.Year{}, err
	}
	// add all 12 months for the year!
	months := [12]string{
		"JAN-",
		"FEB-",
		"MAR-",
		"APRIL-",
		"MAY-",
		"JUNE-",
		"JULY-",
		"AUG-",
		"SEPT-",
		"OCT-",
		"NOV-",
		"DEC-",
	}
	for _, v := range months {
		// fmt.Printf("2**%d = %d\n", i, v)
		q2 := `INSERT INTO yearmonths(id, description, year_id, status)
    VALUES($1, $2, $3, $4)`
		_, err = db.conn.Exec(
			ctx,
			q2,
			formattedID,
			v+Year.Description,
			formattedID,
			1,
		)
		if err != nil {
			return api.Year{}, err
		}
	}
	return Year, nil
}

func (db *Db) AddPeriods(ctx context.Context, Period api.Period) (api.Period, error) {
	query := `SELECT COUNT(*) FROM Periods`
	var count int
	rows, err := db.conn.Query(ctx, query)
	rows.Scan(&count)
	if err != nil {
		return api.Period{}, fmt.Errorf("unable to query periods: %w", err)
	}
	// pgx.CollectOneRow(rows, pgx.RowToStructByName(count))
	defer rows.Close()
	// rows, err := db.conn.Query(ctx, "SELECT COUNT(*) FROM class").
	// 	Scan(&count)
	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO periods(id, period_description, status, year_id)
    VALUES($1, $2, $3, $4)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		Period.PeriodDescription,
		Period.Status,
		Period.YearId,
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

func (db *Db) GetYearWithMonths(ctx context.Context, yearId string) ([]api.YearMonthJoin, error) {
	q := `SELECT year.id, 
               year.description, 
               year.status, 
               (
                 SELECT json_agg(
                   jsonb_build_object(
                     'description', yearmonth.description,
                     'id', yearmonth.id,
                     'status', yearmonth.status,
                     'yearId', yearmonth.year_id
                   )
                 )
                 FROM yearmonths AS yearmonth
                 WHERE yearmonth.year_id = year.id
               ) AS months
        FROM year 
        WHERE year.id = $1
        `
	rows, err := db.conn.Query(ctx, q, yearId)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.YearMonthJoin])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}

func (db *Db) GetYears(ctx context.Context) ([]api.Year, error) {
	q := `SELECT *
        FROM year 
        `
	rows, err := db.conn.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.Year])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}

func (db *Db) GetStudentReceipts(ctx context.Context) ([]api.Student_receipt_body, error) {
	q := `SELECT receipt.id, 
               receipt.description, 
               receipt.amount, 
               receipt.student_id,
               receipt.incadmissionfee,
               receipt.discount,
               receipt.discountinpercent,
               receipt.subtotal,
               (
                 SELECT json_agg(
                   jsonb_build_object(
                     'month', receipt_month.description,
                     'description', receipt_month.id,
                     'amount', receipt_month.amount
                   )
                 )
                 FROM student_receipt_months AS receipt_month
                 WHERE receipt_month.receipt_id = receipt.id
               ) AS months
        FROM student_receipts AS receipt 
        `
	rows, err := db.conn.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()
	// var data []api.Student
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[api.Student_receipt_body])
	if err != nil {
		return nil, fmt.Errorf("unable to collect rows: %w", err)
	}
	return products, nil
}

func (db *Db) AddStudentReceipt(
	ctx context.Context,
	receipt api.Student_receipt_body,
) (api.Student_receipt_body, error) {
	query := `SELECT COUNT(*) FROM student_receipts`
	var count int
	rows, err := db.conn.Query(ctx, query)
	if err != nil {
		return api.Student_receipt_body{}, fmt.Errorf("unable to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return api.Student_receipt_body{}, fmt.Errorf("unable to scan row: %w", err)
		}
	} else {
		return api.Student_receipt_body{}, fmt.Errorf("no rows returned")
	}

	if err := rows.Err(); err != nil {
		return api.Student_receipt_body{}, fmt.Errorf("error iterating over rows: %w", err)
	}

	nextId := count + 1
	formattedID := fmt.Sprintf("%04d", nextId)
	q := `INSERT INTO student_receipts(id, description, student_id, subtotal, discount, incadmissionfee, 
    amount, discountinpercent)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.conn.Exec(
		ctx,
		q,
		formattedID,
		receipt.Description,
		receipt.Student_Id,
		receipt.SubTotal,
		receipt.Discount,
		receipt.IncAdmissionFee,
		receipt.Amount,
		receipt.DiscountInPercent,
	)
	if err != nil {
		fmt.Println("student receipt adding issue?")
		return api.Student_receipt_body{}, err
	}
	for _, v := range receipt.Months {
		q2 := `INSERT INTO student_receipt_months(month, amount, description, receipt_id)
    VALUES($1, $2, $3, $4)`
		_, err = db.conn.Exec(
			ctx,
			q2,
			v.Month,
			v.Amount,
			v.Description,
			formattedID,
		)
		if err != nil {
			fmt.Println("student month receipt adding issue?")
			return api.Student_receipt_body{}, err
		}
	}
	return receipt, nil
}
