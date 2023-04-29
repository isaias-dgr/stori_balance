package adapter

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"

	"github.com/isaias-dgr/story-balance/src/internal/core/domain"
)

type Repository struct {
	log    *log.Logger
	client *sql.DB
}

func NewRepository(
	l *log.Logger, user, password, url, db_name string) *Repository {
	conn_string := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, db_name)
	l.Info(conn_string)
	db, err := sql.Open("mysql", conn_string)
	if err != nil {
		l.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		l.Fatal("Error pinging database:", err)
		return nil
	}
	l.Info("Ping successful")
	return &Repository{
		log:    l,
		client: db,
	}
}

func (r Repository) Save(
	account, product string, transaction *domain.Transaction) error {
	query := "INSERT INTO transactions (user_id, product, code, description, date, amount) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.client.Exec(query, account, product, transaction.Code,
		transaction.Description, transaction.Date, transaction.Amount)
	if err != nil {
		r.log.Errorf("Data not inserted successfully %s", err.Error())
		return err
	}
	r.log.Info("Data saved")
	return nil
}

func (r Repository) GetUser(id string) (*domain.User, error) {
	email := os.Getenv("TEMP_EMAIl") //setup and validate this on ses service
	tel := os.Getenv("TEMP_TEL")     //setup and validate this on sns service
	return &domain.User{
		Id:                 id,
		Email:              email,
		Tel:                tel,
		Notification_email: email != "",
		Notification_sms:   tel != "",
	}, nil
}
