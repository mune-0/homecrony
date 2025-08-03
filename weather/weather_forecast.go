package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherForecast struct {
	Geometry ForecastGeometry `json:"geometry"`;
	Properties ForecastProperties `json:"properties"`;
}

type ForecastGeometry struct {
	Type string `json:"type"`;
	Coordinates [][][]float64 `json:"coordinates"`;
}

type ForecastProperties struct {
	Zone string `json:"zone"`;
	Updated string `json:"updated"`;
	Periods []ForecastPeriod `json:"periods"`;
}

type ForecastPeriod struct {
	Number int `json:"number"`;
	Name string `json:"name"`;
	DetailedForecast string `json:"detailedForecast"`;
}


func GetWeatherForecast() {
	// Zone Forecast
	resp, err := http.Get("https://api.weather.gov/zones/forecast/FLZ249/forecast")
	if(err != nil) {
		fmt.Println("Err: ", err.Error())
		return;
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if(err != nil) {
		fmt.Println("Err: ", err.Error())
		return;
	}

	var forecastData WeatherForecast;
	err = json.Unmarshal(respBody, &forecastData)
	if(err != nil) {
		fmt.Println("Err: ", err.Error())
		return;
	}
	fmt.Println("Response: ", forecastData)
}