package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	locationfind "github.com/ZiganshinDev/telebot/internal/service/locationFind"
)

type Weather struct {
	Now   int       `json:"now"`
	NowDt time.Time `json:"now_dt"`
	Info  struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
		URL string  `json:"url"`
	} `json:"info"`
	Fact struct {
		Temp       int     `json:"temp"`
		FeelsLike  int     `json:"feels_like"`
		Icon       string  `json:"icon"`
		Condition  string  `json:"condition"`
		WindSpeed  int     `json:"wind_speed"`
		WindGust   float64 `json:"wind_gust"`
		WindDir    string  `json:"wind_dir"`
		PressureMm int     `json:"pressure_mm"`
		PressurePa int     `json:"pressure_pa"`
		Humidity   int     `json:"humidity"`
		Daytime    string  `json:"daytime"`
		Polar      bool    `json:"polar"`
		Season     string  `json:"season"`
		ObsTime    int     `json:"obs_time"`
	} `json:"fact"`
	Forecast struct {
		Date     string `json:"date"`
		DateTs   int    `json:"date_ts"`
		Week     int    `json:"week"`
		Sunrise  string `json:"sunrise"`
		Sunset   string `json:"sunset"`
		MoonCode int    `json:"moon_code"`
		MoonText string `json:"moon_text"`
		Parts    []struct {
			PartName   string  `json:"part_name"`
			TempMin    int     `json:"temp_min"`
			TempMax    int     `json:"temp_max"`
			TempAvg    int     `json:"temp_avg"`
			FeelsLike  int     `json:"feels_like"`
			Icon       string  `json:"icon"`
			Condition  string  `json:"condition"`
			Daytime    string  `json:"daytime"`
			Polar      bool    `json:"polar"`
			WindSpeed  float64 `json:"wind_speed"`
			WindGust   int     `json:"wind_gust"`
			WindDir    string  `json:"wind_dir"`
			PressureMm int     `json:"pressure_mm"`
			PressurePa int     `json:"pressure_pa"`
			Humidity   int     `json:"humidity"`
			PrecMm     int     `json:"prec_mm"`
			PrecPeriod int     `json:"prec_period"`
			PrecProb   int     `json:"prec_prob"`
		} `json:"parts"`
	} `json:"forecast"`
}

type Client struct {
	http   *http.Client
	apiKey string
	city   string
}

func NewClient(city string) *Client {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	apiKey := os.Getenv("YANDEX_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	return &Client{httpClient, apiKey, city}
}

func formatingWeatherData(w *Weather) {
	if w.Fact.Daytime == "d" {
		w.Fact.Daytime = "day"
	} else {
		w.Fact.Daytime = "night"
	}
}

func (c *Client) fetchWeather() (*Weather, error) {
	lat, lon := locationfind.ReflectLocation(locationfind.NewClient(c.city))

	method := "GET"
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=%f&lon=%f&lang=ru_RU", lat, lon)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Yandex-API-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return &Weather{}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return &Weather{}, nil
	}

	var weather Weather
	_ = json.NewDecoder(resp.Body).Decode(&weather)

	formatingWeatherData(&weather)

	return &weather, nil
}

func ReflectWeather(c *Client) string {
	res, err := c.fetchWeather()
	if err != nil {
		fmt.Println(err)
	}

	if res.Info.Lat == 0 && res.Info.Lon == 0 {
		return fmt.Sprintf("This city does not exist or was not found `%v`", c.city)
	} else {
		return fmt.Sprintf("Now in %v: \nTemperature: %v\nFeel like: %v\nCondition: %v\nWind speed: %v\nWind direction: %v\nHumidity: %v\nDay Time: %v\nSeason: %v\n\nCreated with Яндекс.Погода", c.city, res.Fact.Temp, res.Fact.FeelsLike, res.Fact.Condition, res.Fact.WindSpeed, res.Fact.WindDir, res.Fact.Humidity, res.Fact.Daytime, res.Fact.Season)
	}
}
