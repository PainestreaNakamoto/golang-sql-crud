package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerModel struct {
	Customer_id string
	Name        string
}

func main() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")

	if err != nil {
		panic(err)
	}
	customer, err := GetCustomer(1001, db)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)

}

func GetCustomers(db *sqlx.DB) ([]CustomerInsertModel, error) {

	query := "select * from customers"
	customers := []CustomerInsertModel{}
	err := db.Select(&customers, query)

	if err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomer(id int, db *sqlx.DB) (*CustomerModel, error) {

	query := "select customer_id, name from customers where customer_id=?"
	customer := CustomerModel{}
	err := db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

type CustomerInsertModel struct {
	Customer_id   int    `db:"customer_id"`
	Name          string `db:"name"`
	Date_of_birth string `db:"date_of_birth"`
	City          string `db:"city"`
	Zipcode       string `db:"zipcode"`
	Status        int    `db:"status"`
}

func AddCustomer(customer CustomerInsertModel, db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "insert into customers (customer_id,name,date_of_birth,city,zipcode,status) values (?,?,?,?,?,?)"
	result, err := tx.Exec(query, customer.Customer_id, customer.Name, customer.Date_of_birth, customer.City, customer.Zipcode, customer.Status)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.RollBack()
		return err
	}

	if affected <= 0 {
		return errors.New("Can't insert")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateCustomer(id int, name string, db *sqlx.DB) error {
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

func DeleteCustomer(id int, db *sqlx.DB) error {
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
