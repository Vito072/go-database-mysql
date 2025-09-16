package go_database_mysql

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('oriza', 'Oriza')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id   :", id)
		fmt.Println("name :", name)
	}
	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birthdate, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name, email string
		var balance int
		var rating float64
		var birthdate, created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthdate, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("===========================")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		fmt.Println("Email :", email)
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		fmt.Println("Birthdate :", birthdate)
		fmt.Println("Married :", married)
		fmt.Println("Created At :", created_at)

	}
	defer rows.Close()
}
