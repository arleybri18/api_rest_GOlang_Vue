package main

import (
	"fmt"

	"github.com/arleybri18/api_rest_GOlang_Vue/consumeAPI"
	"github.com/arleybri18/api_rest_GOlang_Vue/consumeWhoIs"
	"github.com/arleybri18/api_rest_GOlang_Vue/database"
	"github.com/arleybri18/api_rest_GOlang_Vue/scrappingWebPage"

	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
)

// global variables
var user string = "yonydb"
var host_server string = "localhost"
var port_server string = "26257"
var name_bd string = "infodomains"

// functions to handle endpoints
func GetDomainEndpoint(w http.ResponseWriter, req *http.Request) {

	// Connect to the "infodomains" database as the "yonydb" user.
	addr := "postgresql://" + user + "@" + host_server + ":" + port_server + "/" + name_bd + "?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var domains_db []database.Domains

	// get the information in url with the key dom
	domain := chi.URLParam(req, "dom")

	// ----------------------

	url := "https://api.ssllabs.com/api/v3/analyze?host=" + domain
	// send url
	domains := consumeAPI.ConsumeAPI(url)
	// fmt.Printf("host: %s \n port: %d \nprocotol: %s \nServers: %+v \n", domains.Host, domains.Port, domains.Protocol, domains.Server)

	// call function to parse country and organization
	name_domain := domains.Host
	countries, organizations := consumeWhoIs.RunWhoisCommand(name_domain)
	// fmt.Printf("Country is %s organization is %s\n", countries, organizations)

	// call to function to get a logo and title
	webUrl := "https://www." + name_domain

	pageTitle, pageLogo := scrappingWebPage.ScrappingWebPage(webUrl)
	// // Print out the result
	// fmt.Printf("Page Title: %s\n", pageTitle)
	// // Print out the result
	// fmt.Printf("Page Logo: %s\n", pageLogo)

	// call to function to insert data, send server stored in domains.server
	database.InsertData(name_domain, true, domains.Server[0].Ssl_grade, pageLogo, pageTitle, false, domains.Server, countries, organizations)

	// ------------------------

	// select into database
	db.Limit(2).Where("domain = ?", domain).Last(&domains_db)
	json.NewEncoder(w).Encode(domains_db)

}

func ShowReportEndpoint(w http.ResponseWriter, req *http.Request) {
	// Connect to the "infodomains" database as the "yonydb" user.
	addr := "postgresql://" + user + "@" + host_server + ":" + port_server + "/" + name_bd + "?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var traces []database.Traces

	db.Order("created_at desc").Find(&traces)
	json.NewEncoder(w).Encode(traces)
}

func main() {

	port := ":5000"
	// create database and tables
	database.InitialMigration()

	// create a new router
	router := chi.NewRouter()

	// endpoints
	// RESTy routes for "domain" resource

	router.Route("/domain", func(router chi.Router) {
		// dom save the domain in the url
		router.Get("/{dom}", GetDomainEndpoint)
	})

	router.Route("/report", func(router chi.Router) {
		router.Get("/", ShowReportEndpoint)
	})

	// start a server
	fmt.Println("Serving on port:" + port)
	log.Fatal(http.ListenAndServe(port, router))

}
