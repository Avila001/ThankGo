package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// WeatherService предсказывает погоду.
type WeatherService struct{}

type MockWeatherService struct {
	name int
}

type Forecaster interface {
	Forecast() int
}

// Weather выдает текстовый прогноз погоды.
type Weather struct {
	service Forecaster
}

// Forecast сообщает ожидаемую дневную температуру на завтра.
func (ws *WeatherService) Forecast() int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(31)
	sign := rand.Intn(2)
	if sign == 1 {
		value = -value
	}
	return value
}

func (ws *MockWeatherService) Forecast() int {
	return ws.name
}

// Forecast сообщает текстовый прогноз погоды на завтра.
func (w Weather) Forecast() string {
	deg := w.service.Forecast()
	switch {
	case deg < 10:
		return "холодно"
	case deg >= 10 && deg < 15:
		return "прохладно"
	case deg >= 15 && deg < 20:
		return "идеально"
	case deg >= 20:
		return "жарко"
	}
	return "инопланетно"
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

func TestForecast(t *testing.T) {
	//service := &WeatherService{}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			service := &MockWeatherService{}
			weather := Weather{service}
			service.name = test.deg
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
