package database

import (
	"context"
	"fmt"
	"log"
	"rogers-software/calipso/models"

	"database/sql"

	_ "github.com/lib/pq"
)

var User string
var Password string
var Host string
var DatabaseName string

var DB *sql.DB
var DNS string

func IniciarDB(ctx context.Context) error {

	User = ctx.Value(models.Key("user")).(string)
	Password = ctx.Value(models.Key("password")).(string)
	Host = ctx.Value(models.Key("host")).(string)
	DatabaseName = ctx.Value(models.Key("database")).(string)

	var err error

	DNS = fmt.Sprintf("postgres://%s:root@%s/%s?sslmode=disable", User, Host, DatabaseName)

	DB, err = sql.Open("postgres", DNS)

	if err != nil {
		fmt.Println("error ", err)
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	//defer DB.Close()

	fmt.Println("Conexion Exitosa con la BD")

	/*
		createProductTable(db)

		product := Product{"Book", 15.55, true}

		pk := insertProduct(db, product)

		fmt.Printf("ID = %d\n", pk)

		var name string
		var available bool
		var price float64

		query := "SELECT name, available, price FROM product WHERE id = $1"

		err = db.QueryRow(query, pk).Scan(&name, &available, &price)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Available: %t\n", available)
		fmt.Printf("Price: %f\n", price)

		// otra forma
		data := []Product{}
		rows, err := db.Query("SELECT name, available, price from product")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&name, &available, &price)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, Product{name, price, available})

		}

		fmt.Println(data)
	*/
	return err

	//age := 21
	//rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)

}

/*
func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXIST product (
	 id SERIAL PRIMARY KEY,
	 name VARCHAR(100) NOT NULL,
	 price NUMERIC(6,2) NOT NULL,
	 vailable BOOLEAN,
	 created timestamp DEFAULT NOW()
	 )`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
			  VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
*/

func BaseConectada(db *sql.DB) bool {
	err := db.Ping()
	return err == nil
}
