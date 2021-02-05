package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spacex-api-siemens/models"
)

//func for getting all capsules from the api
func getCrew() []models.Crew {
	resp, err := http.Get("https://api.spacexdata.com/v4/crew")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var crews []models.Crew
	json.Unmarshal(body, &crews)
	return crews
}

func CrewCount() int {
	return len(getCrew())
}

func CrewNames() []string {
	var names []string
	crews := getCrew()
	for _, crew := range crews {
		names = append(names, crew.Name)
	}
	return names
}

func LaunchCount() int {
	var count int
	crews := getCrew()
	for _, crew := range crews {
		count += len(crew.Launches)
	}
	return count
}
