package database

import (
	"encoding/json"
	"log"

	"github.com/arleybri18/API_Rest_GO/consumeAPI"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Declare structs for marshal

type Servers_raw struct {
	Address   string
	Ssl_grade string
	Country   string
	Owner     string
}

type arr_server_raw struct {
	serv_raw []Servers_raw
}

func InsertData(name_domain string, server_changed bool, prev_ssl string, pageLogo string,
	pageTitle string, Is_down bool, server consumeAPI.Servers_j, countries string, organizations string) {

	// open conection to databse
	const addr = "postgresql://yonydb@localhost:26257/infodomains?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// convert data to insert in field json
	serv := Servers_raw{server.Address, server.Ssl_grade, countries, organizations}

	serv_byte, err := json.Marshal(serv)

	// insert into database

	db.Create(&Domains{Domain: name_domain, Servers_changed: server_changed, Previous_ssl_grade: prev_ssl,
		Logo: pageLogo, Title: pageTitle, Is_down: Is_down, Servers: postgres.Jsonb{serv_byte}})

	db.Create(&Servers{Address: server.Address, SSL_grade: server.Ssl_grade, Country: countries, Owner: organizations})

	db.Create(&Traces{Item: name_domain + " info"})

}
