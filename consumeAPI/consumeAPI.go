package consumeAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Declare structs for JSON
type Servers_j struct {
	Address   string `json:"serverName"`
	Ssl_grade string `json:"grade"`
}

type Info_domain struct {
	Host     string      `json:"host"`
	Port     int64       `json:"port"`
	Protocol string      `json:"protocol"`
	Server   []Servers_j `json:"endpoints"`
}
// the function should begining with the uppercase name for export it
func ConsumeAPI(url string) Info_domain {
	var infoDomain Info_domain

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request fail with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &infoDomain)
		return infoDomain
	}
	return infoDomain
}
