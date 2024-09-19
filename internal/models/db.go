package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Mapping struct {
	fileId   string
	fileName string
	password string
}

func ConnectDB() (*sql.DB, error) {
	connStr := "user=admin password=pass1234 dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertToDb(db *sql.DB, fileName string, fileId string, password string) error {
	query := `insert into file_mapping values($1,$2,$3)`
	_, err := db.Exec(query, fileId, fileName, password)
	if err != nil {
		return err
	}
	return nil
}

func UserValidate(db *sql.DB, fileId string, password string) (string, error) {
	query := `select fileid , filename, password from file_mapping where fileid=$1 and password=$2)`
	var mapping Mapping
	err := db.QueryRow(query, fileId, password).Scan(&mapping.fileId, &mapping.fileName, &mapping.password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("error querying user by ID: %w", err)
	}
	return mapping.fileName, nil
}
