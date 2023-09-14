package yandex_weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	http   *http.Client
	apiKey string
}

type WeatherGetter interface {
	GetWeather(lat float64, lon float64) error
}

func New() (*Client, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	apiKey := os.Getenv("YANDEX_API_KEY")
	if apiKey == "" {
		log.Fatal("env: apiKey must be set")
	}

	return &Client{http: httpClient, apiKey: apiKey}, nil
}

func (c *Client) GetWeather(lat float64, lon float64) (string, error) {
	const op = "service.yandex-weather.GetWeather"

	method := "GET"
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=%v&lon=%v&lang=ru_RU", lat, lon)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Add("X-Yandex-API-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s; %s: %w", resp.Status, op, err)
	}

	var weather Weather
	_ = json.NewDecoder(resp.Body).Decode(&weather)

	formatingWeatherData(&weather)

	if weather.Info.Lat == 0 && weather.Info.Lon == 0 {
		return "", err
	} else {
		weather := fmt.Sprintf("Temperature: %v\nFeel like: %v\nCondition: %v\nWind speed: %v\nWind direction: %v\nHumidity: %v\nDay Time: %v\nSeason: %v\n\nCreated with Яндекс.Погода", weather.Fact.Temp, weather.Fact.FeelsLike, weather.Fact.Condition, weather.Fact.WindSpeed, weather.Fact.WindDir, weather.Fact.Humidity, weather.Fact.Daytime, weather.Fact.Season)
		return weather, nil
	}
}

func formatingWeatherData(w *Weather) {
	if w.Fact.Daytime == "d" {
		w.Fact.Daytime = "day"
	} else {
		w.Fact.Daytime = "night"
	}
}
