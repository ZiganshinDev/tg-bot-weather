package open_weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client struct {
	http   *http.Client
	apiKey string
}

func New() (*Client, error) {
	const op = "service.open-weather.New"

	httpClient := &http.Client{Timeout: 10 * time.Second}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("%s: env: apiKey must be set", op)
	}

	return &Client{http: httpClient, apiKey: apiKey}, nil
}

func (c *Client) GetCoordinates(city string) (float64, float64, error) {
	const op = "service.open-weather.GetCoordinates"

	method := "GET"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v&APPID=%v", city, c.apiKey)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("%s; %s: %w", resp.Status, op, err)
	}

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return 0, 0, fmt.Errorf("%s: %w", op, err)
	} else {
		return location.Coord.Lat, location.Coord.Lon, nil
	}
}
