package domain

import (
	"database/sql"
	"log"
	"github.com/kaliayev-proton/banking-go-hex/errors"
	"github.com/kaliayev-proton/banking-go-hex/logger"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errors.AppError) {

	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
		} else {
			findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
			err = d.client.Select(&customers, findAllSql, status)
			// rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customer " + err.Error())
	// 	return nil, errors.NewUnexpectedError("Unexpected database error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customer " + err.Error())
	// 		return nil, errors.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	// fmt.Println(customers)


	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// row := d.client.QueryRow(customerSql, id)

	var c Customer
	err :=d.client.Get(&c, customerSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer" + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}