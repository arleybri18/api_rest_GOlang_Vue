package consumeWhoIs

import (
	"bytes"
	"os/exec"
	"strings"
)

// the function should begining with the uppercase name for export it
func RunWhoisCommand(name_domain string) (string, string) {
	// Store output on buffer
	var out bytes.Buffer
	var country string
	var organization string

	// Execute command
	cmd := exec.Command("whois", name_domain)
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Run()

	// cast to string
	out_s := out.String()

	// split colon
	s := strings.Split(out_s, ":")
	// flags to out of the loop
	flagc := 0
	flago := 0

	// iterate over slice
	for index, element := range s {

		// validate pattern for search country
		if strings.Contains(element, "Registrant Country") && flagc == 0 {

			esp := strings.Split(s[index+1], " ")
			country = esp[1][0:2]
			// fmt.Printf("Country es -%s- \n", country )
			flagc = 1
		}

		// validate pattern for search Organization
		if strings.Contains(element, "Registrant Organization") && flago == 0 {
			esp := strings.Split(s[index+1], " ")
			organization = esp[1][0:]
			// fmt.Printf("Organization is es -%s- \n", organization )
			flago = 1
		}
	}
	return country, organization
}
