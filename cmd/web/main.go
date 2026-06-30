package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func main() {

	url := "https://groupietrackers.herokuapp.com/api"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		apibytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(apibytes))
	}

}
