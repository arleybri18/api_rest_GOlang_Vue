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
	// Servers string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}
type Servers struct {
	gorm.Model
	Domain string
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
	dbroot, err := gorm.Open("postgres", "postgresql://root@"+host_server+":"+port_server+"/defaultdbpostgres/?sslmode=disable")
	dbroot.Exec("CREATE USER IF NOT EXISTS yonydb;")
	dbroot.Exec("CREATE DATABASE IF NOT EXISTS infodomains;")
	dbroot.Exec("GRANT ALL ON DATABASE infodomains TO yonydb;")
	defer dbroot.Close()

	// end create database
	// open conection to databse
	addr := "postgresql://" + user + "@" + host_server + ":" + port_server + "/" + name_bd + "?sslmode=disable"
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
