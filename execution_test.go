package go_database_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var birthdate sql.NullTime
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthdate, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("===========================")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthdate.Valid {
			fmt.Println("Birthdate :", birthdate)
		}
		fmt.Println("Married :", married)
		fmt.Println("Created At :", created_at)

	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"
	//rawan peretasan menggunakan SQL injection
	script := "SELECT username, password FROM user WHERE username ='" + username +
		"'AND password = '" + password + "'LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		rows.Scan(username)
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("Wrong Username Or Password")
	}
	defer rows.Close()
}

// berikut adalah solusi untuk permasalahan SQL Injection
func TestQueryWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username, password FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		rows.Scan(username)
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("Wrong Username Or Password")
	}
	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "vito"
	password := "vito"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "vito@gmail.com"
	comment := "Auto"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "vito" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("comment Id :", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "vito" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("comment Id :", lastInsertId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
	tx.Commit()
}
