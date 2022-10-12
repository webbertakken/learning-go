package mdb

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type EmailEntry struct {
	Id           int64
	EmailAddress string
	ConfirmedAt  *time.Time
	IsOptedOut   bool
}

func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE emails (
        id INTEGER PRIMARY KEY,
        email_address TEXT UNIQUE ,
        confirmed_at_date INTEGER,
        is_opted_out INTEGER,
    );
  `)

	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			// code 1 == "Table already exists
			if sqlError.Code == 1 {
				log.Fatalln(sqlError)
			}
		} else {
			log.Fatalln(err)
		}
	}
}

func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var id int64
	var emailAddress string
	var confirmedAt int64
	var isOptedOut bool

	err := row.Scan(&id, &emailAddress, &confirmedAt, &isOptedOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	t := time.Unix(confirmedAt, 0)

	return &EmailEntry{id, emailAddress, &t, isOptedOut}, nil
}

func CreateEmail(db *sql.DB, email string) error {
	_, err := db.Exec(`
        INSERT INTO
          emails(email_address, confirmed_at_date, is_opted_out)
          VALUES(?, 0, false)
    `, email)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetEmail(db *sql.DB, email string) (*EmailEntry, error) {
	rows, err := db.Query(`
    SELECT id, email_address, confirmed_at_date, is_opted_out
    FROM emails
    WHERE email = ?
  `, email)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		return emailEntryFromRow(rows)
	}

	return nil, nil
}

func UpdateEmail(db *sql.DB, entry EmailEntry) error {
	confirmedAtDate := entry.ConfirmedAt.Unix()

	_, err := db.Exec(`
    INSERT INTO
      emails(email_address, confirmed_at_date, is_opted_out)
      VALUES (?, ?, ?)
      ON CONFLICT(email) DO UPDATE SET
        confirmed_at_date=?
        is_opted_out=?
  `, entry.EmailAddress, confirmedAtDate, entry.IsOptedOut, confirmedAtDate, entry.IsOptedOut)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func OptOutEmail(db *sql.DB, emailAddress string) error {
	_, err := db.Exec(`
    UPDATE emails
    SET is_opted_out=true
    WHERE email=?
  `, emailAddress)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type GetEmailBatchQueryParams struct {
	Page  int
	Count int
}

func GetEmailBatch(db *sql.DB, params GetEmailBatchQueryParams) ([]EmailEntry, error) {
	rows, err := db.Query(`
    SELECT id, email_address, confirmed_at_date, is_opted_out
    FROM emails
    WHERE is_opted_out=false
    ORDER BY id asc
    LIMIT ? OFFSET ?
  `, params.Count, (params.Page-1)*params.Count)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	emails := make([]EmailEntry, 0, params.Count)
	for rows.Next() {
		email, err := emailEntryFromRow(rows)
		if err != nil {
			return nil, err
		}

		emails = append(emails, *email)
	}

	return emails, nil
}
