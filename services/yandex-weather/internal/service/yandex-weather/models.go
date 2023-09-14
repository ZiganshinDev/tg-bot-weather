package yandex_weather

import "time"

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
