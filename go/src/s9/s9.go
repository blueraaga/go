package main

import ("os"
		"fmt"
		"net/http"
		"io/ioutil")


func main() {
	fmt.Println("Initiating URL GET request")
	
	resp, err := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	/* check if there was no error, else program returns nil pointer
	   that leads to error with ioutil.ReadAll further down the line */
	if err != nil {
		fmt.Printf("Error found\n") // debug error
		fmt.Println(resp)            // debug error
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Requesting over")
	
	fmt.Println("Pulling out response body")
	
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}