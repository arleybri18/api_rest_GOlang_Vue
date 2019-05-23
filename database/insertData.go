package database

import (
	"encoding/json"
	"log"

	"github.com/arleybri18/api_rest_GOlang_Vue/consumeAPI"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// // global variables
var user string = "yonydb"
var host_server string = "localhost"
var port_server string = "26257"
var name_bd string = "infodomains"

// note: function receive array of Servers_j
func InsertData(name_domain string, server_changed bool, prev_ssl string, pageLogo string,
	pageTitle string, Is_down bool, server []consumeAPI.Servers_j, countries string, organizations string) {

	// open conection to databse
	addr := "postgresql://" + user + "@" + host_server + ":" + port_server + "/" + name_bd + "?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Put data country and owner in struct servers
	for index, _ := range server {
		server[index].Domain = name_domain
		server[index].Country = countries
		server[index].Owner = organizations
	}

	// insert into database

	for _, server_n := range server {
		db.Create(&Servers{Domain: server_n.Domain, Address: server_n.Address, SSL_grade: server_n.Ssl_grade, Country: countries, Owner: organizations})
	}

	// declacre array to save data servers
	var serv_arr []Servers
	var change bool

	for idx, _ := range server {
		db.Order("created_at desc").Limit(2).Select([]string{"Domain", "Address", "SSL_grade", "Country"}).Where("Domain= ? AND Address= ?", name_domain, server[idx].Address).Find(&serv_arr)

		if len(serv_arr) > 1 {
			if serv_arr[0].Address == serv_arr[1].Address &&
				serv_arr[0].Country == serv_arr[1].Country &&
				serv_arr[0].SSL_grade == serv_arr[1].SSL_grade &&
				serv_arr[0].Owner == serv_arr[1].Owner {
				change = false
			} else {
				change = true
			}
		} else {
			change = false
		}
	}

	// insert trace
	db.Create(&Traces{Item: name_domain + " info"})

	// insert data in domains
	serv_byte_arr, err := json.Marshal(server)
	db.Create(&Domains{Domain: name_domain, Servers_changed: change, Previous_ssl_grade: serv_arr[len(serv_arr)-1].SSL_grade,
		Logo: pageLogo, Title: pageTitle, Is_down: Is_down, Servers: postgres.Jsonb{serv_byte_arr}})

}
