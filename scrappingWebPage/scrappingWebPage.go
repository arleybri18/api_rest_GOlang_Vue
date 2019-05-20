package scrappingWebPage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ScrappingWebPage(webUrl string) (string, string) {
	var pageLogo []byte
	var pageTitle []byte

	response_logo, err := http.Get(webUrl)
	if err != nil {
		fmt.Printf("The HTTP request fail with error %s\n", err)
	} else {
		defer response_logo.Body.Close()
		//read data in response
		dataBytes, _ := ioutil.ReadAll(response_logo.Body)
		// cast to string
		pageContent := string(dataBytes)
		pageContent2 := string(dataBytes)

		//find substring
		logoStartIndex := strings.Index(pageContent, "<link href=")
		if logoStartIndex == -1 {
			fmt.Println("No logo element found")
			pageLogo = []byte("Not Found")
		} else {
			// The start index of the logo is the index of the first
			// character, the < symbol. We don't want to include
			// <logo> as part of the final value, so let's offset
			// the index by the number of characers in <logo>
			logoStartIndex += 11
		}

		// Find the index of the closing tag
		logoEndIndex := strings.Index(pageContent, " rel=\"")
		if logoEndIndex == -1 {
			fmt.Println("No closing tag for logo found.")
			pageLogo = []byte("Not Found")
		} else {

			// Copy the substring in to a separate variable so the
			// variables with the full document data can be garbage collected
			pageLogo = []byte(pageContent[logoStartIndex:logoEndIndex])

		}

		// find title
		//find substring
		titleStartIndex := strings.Index(pageContent2, "<title>")
		if titleStartIndex == -1 {
			fmt.Println("No title element found")
			pageTitle = []byte("Not Found")
		} else {
			// The start index of the logo is the index of the first
			// character, the < symbol. We don't want to include
			// <logo> as part of the final value, so let's offset
			// the index by the number of characers in <logo>
			titleStartIndex += 7
		}

		// Find the index of the closing tag
		titleEndIndex := strings.Index(pageContent2, "</title>")
		if titleEndIndex == -1 {
			fmt.Println("No closing tag for title found.")
			pageTitle = []byte("Not Found")
		} else {

			// Copy the substring in to a separate variable so the
			// variables with the full document data can be garbage collected
			pageTitle = []byte(pageContent2[titleStartIndex:titleEndIndex])
		}

		return string(pageTitle), string(pageLogo)
	}
	return string(pageTitle), string(pageLogo)
}
