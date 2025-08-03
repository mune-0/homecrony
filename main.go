package main

import (
	"fmt"

	"org.mnzr/crony/weather"
)


func main() {
	fmt.Println("HomeCrony is starting!")
	weather.GetWeatherForecast()
}