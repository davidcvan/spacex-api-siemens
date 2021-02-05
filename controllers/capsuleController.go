package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spacex-api-siemens/models"
)

//func for getting all capsules from the api
func getCapsules() []models.Capsule {
	resp, err := http.Get("https://api.spacexdata.com/v4/capsules")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var capsules []models.Capsule
	json.Unmarshal(body, &capsules)
	return capsules
}

//func for returning the total number of capsules
func CapsuleCount() int {
	return len(getCapsules())
}

//func for calculating the total water landings for all capsules
func CapsuleWaterLandings() int {
	var landings int
	capsules := getCapsules()
	for _, capsule := range capsules {
		landings += capsule.WaterLandings
	}
	return landings
}

//func for calculating the total of land landings for all capsules
func CapsuleLandLandings() int {
	var landings int
	capsules := getCapsules()
	for _, capsule := range capsules {
		landings += capsule.LandLandings
	}
	return landings
}
