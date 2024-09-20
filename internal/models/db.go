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

func InsertToDb(fileName string, password string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	query := `insert into file_mapping(filename, password) values($1,$2) returning fileid`
	var fileId string
	err2 := db.QueryRow(query, fileName, password).Scan(&fileId)
	if err2 != nil {
		return "", err2
	}
	return fileId, nil
}

func UserValidate(fileId string, password string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	query := `select fileid::TEXT , filename, password from file_mapping where fileid=$1 and password=$2)`
	var mapping Mapping
	err2 := db.QueryRow(query, fileId, password).Scan(&mapping.fileId, &mapping.fileName, &mapping.password)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("error querying user by ID: %w", err2)
	}
	return mapping.fileName, nil
}
