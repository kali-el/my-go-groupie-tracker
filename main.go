package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	url := "https://groupietrackers.herokuapp.com/api/relation"

	type Artists struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations    string   `json:"locations"`
		ConcertDates string   `json:"concertDates"`
		Relations    string   `json:"relations"`
	}

	type locations struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	}

	type concertDates struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	}

	type relations struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}

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
