package domain

import (
	"strconv"
	"github.com/jmoiron/sqlx"
	"github.com/kaliayev-proton/banking-go-hex/errors"
	"github.com/kaliayev-proton/banking-go-hex/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errors.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}
	
	id, err := result.LastInsertId()
	
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	tx, err := d.client.Begin()

	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?,?,?,?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
		} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	account, appErr := d.FindBy(t.AccountId)

	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	t.Amount = account.Amount
	return &t, nil
}