package usecase

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/isaias-dgr/story-balance/src/internal/core/domain"
	"github.com/isaias-dgr/story-balance/src/internal/core/ports"
)

type NotifyBalance struct {
	log        *log.Logger
	repo       ports.Repository
	mailer     ports.Mailer
	stg_ingest ports.Storage
	stg_notify ports.Storage
	notifier   ports.Notifier
	writter    ports.Docs
}

func NewNotifyBalance(
	l *log.Logger,
	r ports.Repository,
	m ports.Mailer,
	si ports.Storage,
	sn ports.Storage,
	n ports.Notifier,
	w ports.Docs,
) *NotifyBalance {
	return &NotifyBalance{
		log:        l,
		repo:       r,
		mailer:     m,
		stg_ingest: si,
		stg_notify: sn,
		notifier:   n,
		writter:    w,
	}
}

func (nb NotifyBalance) Execute(file_address string) (err error) {
	nb.log.Infof("📲 Notify balance %s", file_address)
	file, err := nb.stg_ingest.GetFile(file_address)
	path_info := strings.Split(file_address, "/")
	if err != nil {
		nb.log.Errorf("Retrival file fails: %s", err.Error())
		return err
	}
	defer file.Close()

	// TODO Add transactions
	account_balance := domain.NewAccount(
		path_info[0], path_info[1], path_info[2])
	err = nb.savingTransaction(account_balance, file)
	if err != nil {
		nb.log.Errorf("Retrival file fails: %s", err.Error())
		return err
	}
	nb.sortTransaction(account_balance)

	balance_doc, err := nb.writter.GetDoc(account_balance)
	if err != nil {
		nb.log.Errorf("Writting doc fails: %s", err.Error())
		return err
	}

	user, err := nb.repo.GetUser(account_balance.Id)
	if err != nil {
		nb.log.Errorf("Getting user %s", err.Error())
		return err
	}

	balance_name := fmt.Sprintf("%d/%d/balance_%s.html", account_balance.Year, account_balance.Month, account_balance.Id)
	url, err := nb.stg_notify.PutFile(balance_name, balance_doc)
	if err != nil {
		nb.log.Errorf("Saving balance fails: %s", err.Error())
		return err
	}

	account_balance.URL = url
	if user.Notification_email {
		err = nb.mailer.Send(user.Email, *account_balance)
		if err != nil {
			nb.log.Errorf("Sending Balance mail balance fails: %s", err.Error())
			return err
		}
	}

	if user.Notification_sms {
		// TODO: Crear un diccionario de mensajes con capacidad de traduccion
		err = nb.notifier.Send_Messages("Tu Balance fue creado", user.Tel)
		if err != nil {
			nb.log.Errorf("Sending Balance sms balance fails: %s", err.Error())
			return err
		}
	}
	return nil
}

func (nb NotifyBalance) savingTransaction(
	acc *domain.Account, file io.ReadCloser) error {
	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			nb.log.Warn("Fail read cvs io.EOF")
			break
		}
		if err != nil {
			nb.log.Errorf("Error reading row: %s", err.Error())
			return err
		}

		transaction, err := nb.getTransactionFromRow(row, acc)
		if err != nil {
			nb.log.Errorf("Error reading row: %s", err.Error())
			return err
		}

		nb.setTransactionToAccount(row[2], transaction, acc)
		acc.Id = row[1]
		err = nb.repo.Save(row[1], row[2], transaction)
		if err != nil {
			nb.log.Errorf("Saving row: %s", err.Error())
			return err
		}
	}

	return nil
}

func (nb NotifyBalance) getTransactionFromRow(row []string, acc *domain.Account) (*domain.Transaction, error) {
	day_transaction := strings.Split(row[3], "/")
	day, err := strconv.Atoi(day_transaction[1])
	if err != nil {
		nb.log.Errorf("Error al convertir la cadena: %s", err)
		return nil, err
	}
	raw_date := fmt.Sprintf("%d-%02d-%02d", acc.Year, acc.Month, day)
	date, err := time.Parse("2006-01-02", raw_date)
	if err != nil {
		nb.log.Errorf("Error parsing date: %s", err)
		return nil, err
	}
	amount, err := strconv.ParseFloat(row[5], 32)
	if err != nil {
		nb.log.Infof("Error parsing float: %s: %s", row[5], err.Error())
		return nil, err
	}
	transaction := &domain.Transaction{
		Code:        row[4],
		Description: row[6],
		Amount:      float32(amount),
		Date:        &date,
	}
	return transaction, nil
}

func (nb NotifyBalance) setTransactionToAccount(
	product string, tran *domain.Transaction, acc *domain.Account) {
	_, exist := acc.Products[product]
	if !exist {
		acc.Products[product] = domain.NewProduct(product, 0)
	}

	acc.Products[product].Transactions = append(acc.Products[product].Transactions, tran)
	acc.Products[product].Total += tran.Amount
}

func (nb NotifyBalance) sortTransaction(acc *domain.Account) {

	for _, prod := range acc.Products {
		tran := prod.Transactions
		sortFunc := func(i, j int) bool {
			return tran[i].Date.Before(*(tran[j].Date))
		}
		sort.Slice(prod.Transactions, sortFunc)
	}
}
