package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Mapping struct {
	fileId   string
	fileName string
	password string
}

func ConnectDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GenerateItemToDb() (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	query := `insert into file_mapping(filename, password) values($1,$2) returning fileid`
	var fileId string
	err2 := db.QueryRow(query, "", "").Scan(&fileId)
	if err2 != nil {
		return "", err2
	}
	return fileId, nil
}

func UpdateDb(fileId string, fileName string, password string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	query := `update file_mapping set filename=$1, password=$2 where fileid=$3`
	_, err2 := db.Exec(query, fileName, password, fileId)
	if err2 != nil {
		return err2
	}
	return nil
}

func UserValidate(fileId string, password string) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "", err
	}
	query := `select fileid::TEXT , filename, password from file_mapping where fileid=$1 and password=$2`
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
