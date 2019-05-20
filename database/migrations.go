package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// declare a variable db
var db *gorm.DB

//declare structure for DB

type Domains struct {
	gorm.Model
	Domain             string
	Servers_changed    bool
	Previous_ssl_grade string
	Logo               string
	Title              string
	Is_down            bool
	Servers            postgres.Jsonb
}
type Servers struct {
	gorm.Model
	Address   string
	SSL_grade string
	Country   string
	Owner     string
}

type Traces struct {
	gorm.Model
	Item string
}

func InitialMigration() {
	// create database
	dbroot, err := gorm.Open("postgres", "postgresql://root@localhost:26257/defaultdbpostgres/?sslmode=disable")
	dbroot.Exec("CREATE USER IF NOT EXISTS yonydb;")
	dbroot.Exec("CREATE DATABASE IF NOT EXISTS infodomains;")
	dbroot.Exec("GRANT ALL ON DATABASE infodomains TO yonydb;")
	defer dbroot.Close()

	// end create database
	// open conection to databse
	const addr = "postgresql://yonydb@localhost:26257/infodomains?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	// Automigrate create tables
	db.AutoMigrate(&Domains{})
	db.AutoMigrate(&Servers{})
	db.AutoMigrate(&Traces{})

}
