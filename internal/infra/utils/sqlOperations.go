package utils

import (
	"database/sql"
	"strconv"
)

func SelectMany[T interface{}](db *sql.DB, query string, builder func(*sql.Rows) (T, error), args ...any) ([]T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var obj []T
	for rows.Next() {
		o, err := builder(rows)
		if err != nil {
			return nil, err
		}
		obj = append(obj, o)
	}

	return obj, nil
}

func SelectOneOrError[T interface{}](db *sql.DB, query string, builder func(*sql.Row) (T, error), defaultValue T, args ...any) (T, error) {
	var obj T
	row := db.QueryRow(query, args...)
	obj, err := builder(row)
	if err != nil {
		return defaultValue, err
	}

	return obj, nil
}

func SelectOneOrEmpty[T interface{}](db *sql.DB, query string, builder func(*sql.Row) (T, error), defaultValue T, args ...any) (T, error) {
	var obj T
	row := db.QueryRow(query, args...)
	obj, err := builder(row)
	if err == sql.ErrNoRows {
		return defaultValue, nil
	} else if err != nil {
		return defaultValue, err
	}

	return obj, nil
}

func Insert(db *sql.DB, query string, args ...any) (string, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return "", err
	}
	insertedId, _ := res.LastInsertId()
	id := strconv.Itoa(int(insertedId))

	return id, nil
}

func Delete(db *sql.DB, query string, args ...any) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}
