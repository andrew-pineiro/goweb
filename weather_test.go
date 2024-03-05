package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	ApiKey  = "26b7f0abdcad4942a05175546240303"
	BaseURI = "http://api.weatherapi.com/v1/current.json?key="
)

type Location struct {
	name            string
	region          string
	country         string
	lat             float32
	lon             float32
	tz_id           string
	localtime_epoch uint64
	localtime       string
}
type Condition struct {
	text string
	icon string
	code int
}
type Current struct {
	last_updated_epoch uint64
	last_updated       string
	temp_c             float32
	temp_f             float32
	is_day             int
	condition          Condition
	wind_mph           float32
	wind_kph           float32
	wind_degree        int
	wind_dir           string
	pressure_mb        float32
	pressure_in        float32
	precip_mm          float32
	precip_in          float32
	humidity           float32
	cloud              int
	feelslike_c        float32
	feelslike_f        float32
	vis_km             float32
	vis_miles          float32
	uv                 float32
	gust_mph           float32
	gust_kph           float32
}
type Weather struct {
	location Location
	current  Current
}

//http://api.weatherapi.com/v1/current.json?key=26b7f0abdcad4942a05175546240303&q=New York&aqi=no

func test() {
	location := os.Args[1]
	url := BaseURI + ApiKey + "&q=" + url.QueryEscape(location) + "&aqi=no"
	fmt.Printf("url: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	defer resp.Body.Close()
	var weather []Weather
	//var data map[string]Weather{}
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Print(weather)
}
