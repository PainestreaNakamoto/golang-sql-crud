package main

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerModel struct {
	customer_id int
	name        string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")

	if err != nil {
		panic(err)
	}

	err = DeleteCustomer(2005, db)
	if err != nil {
		panic(err)
	}

}

func GetCustomers(db *sql.DB) ([]CustomerModel, error) {

	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := "select customer_id, name from customers"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	customers := []CustomerModel{}
	for rows.Next() {
		customer := CustomerModel{}
		err = rows.Scan(&customer.customer_id, &customer.name)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func GetCustomer(id int, db *sql.DB) (*CustomerModel, error) {

	query := "select customer_id,name from customers where customer_id=?"
	row := db.QueryRow(query, id)
	customer := CustomerModel{}
	err := row.Scan(&customer.customer_id, &customer.name)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

type CustomerInsertModel struct {
	customer_id   int
	name          string
	date_of_birth string
	city          string
	zipcode       string
	status        int
}

func AddCustomer(customer CustomerInsertModel, db *sql.DB) error {
	query := "insert into customers (customer_id,name,date_of_birth,city,zipcode,status) values (?,?,?,?,?,?)"
	result, err := db.Exec(query, customer.customer_id, customer.name, customer.date_of_birth, customer.city, customer.zipcode, customer.status)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Can't insert")
	}
	return nil
}

func UpdateCustomer(id int, name string, db *sql.DB) error {
	query := "update customers set name=? where customer_id=?"
	result, err := db.Exec(query, name, id)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Can't update")
	}
	return nil
}

func DeleteCustomer(id int, db *sql.DB) error {
	query := "delete from customers where customer_id=?"
	result, err := db.Exec(query, id)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Can't delete")
	}
	return nil
}
