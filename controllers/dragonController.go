package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spacex-api-siemens/models"
	"time"
)

func getDragons() []models.Dragon {
	resp, err := http.Get("https://api.spacexdata.com/v4/dragons")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var dragons []models.Dragon
	json.Unmarshal(body, &dragons)
	return dragons
}

func DragonCount() int {
	return len(getDragons())
}

func TotalPayload() int {
	var payload int
	dragons := getDragons()
	for _, dragon := range dragons {
		payload += dragon.LaunchPayloadMass.Kg
	}
	return payload
}

func DaysSinceLaunch() map[string]int {
	timeFormat := "2006-01-02"
	m := make(map[string]int)
	dragons := getDragons()
	for _, dragon := range dragons {
		t, _ := time.Parse(timeFormat, dragon.FirstFlight)
		duration := time.Now().Sub(t)
		m[dragon.Name] = int(duration.Hours() / 24)
	}
	return m
}
