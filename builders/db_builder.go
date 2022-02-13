package builders

import (
	"context"
	"database/sql"
	"linkconverter-api/models"
	"log"
	"time"
)

type DbBuilderInterface interface {
	DbConnection() error
	InsertLogEvent(logEvent models.LogEventModel) error
}

type DbBuilder struct {
	db *sql.DB
}

func (dbBuilder *DbBuilder) DbConnection() error {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/mysql")
	if err != nil {
		return err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+"my_scheme")
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	db.Close()
	db, err = sql.Open("mysql", "root:root@tcp(db:3306)/my_scheme")
	if err != nil {
		return err
	}
	//defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	err = dbBuilder.CreateLogsTable(db)
	if err != nil {
		return err
	}
	dbBuilder.db = db
	return nil
}

func (dbBuilder *DbBuilder) CreateLogsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS logs(id int primary key auto_increment, requestUrl text, responseUrl text, 
		created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (dbBuilder *DbBuilder) InsertLogEvent(logEvent models.LogEventModel) error {

	query := "INSERT INTO logs(requestUrl, responseUrl) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := dbBuilder.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, logEvent.RequestUrl, logEvent.ResponseUrl)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("request log created")
	return nil
}
func NewDbBuilder() DbBuilderInterface {
	return &DbBuilder{}
}
