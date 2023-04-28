package locationfind

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Location struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Client struct {
	http   *http.Client
	apiKey string
	city   string
}

func NewClient(city string) *Client {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	return &Client{httpClient, apiKey, city}
}

func (c *Client) fetchLocation() (*Location, error) {
	method := "GET"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v&APPID=%v", c.city, c.apiKey)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return &Location{}, nil
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return &Location{}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return &Location{}, nil
	}

	var location Location
	_ = json.NewDecoder(resp.Body).Decode(&location)

	return &location, nil
}

func ReflectLocation(c *Client) (float64, float64) {
	res, err := c.fetchLocation()
	if err != nil {
		log.Println(err)
	}

	return res.Coord.Lat, res.Coord.Lon
}
