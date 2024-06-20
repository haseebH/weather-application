package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/haseebh/weatherapp_temperature/internal/entities/repository"
	"net/http"
	"time"
)

type WeatherUseCase interface {
	FetchTemperature(location string, sDate time.Time) (*repository.Temperature, error)
}
type weatherService struct {
	APIKey  string
	APIBase string
}

func NewWeatherService(apiKey, apiBase string) WeatherUseCase {
	return &weatherService{
		APIKey:  apiKey,
		APIBase: apiBase,
	}
}

// FetchTemperature
// todo: convert to startTime and Endtime to make sure we call weather api efficiently
// todo: integrate with location to coordinates service
func (ws *weatherService) FetchTemperature(location string, sDate time.Time) (*repository.Temperature, error) {
	startDateStr := sDate.Format("2006-01-02")
	lat, long := getCoordinates(location)
	url := fmt.Sprintf("%s?latitude=%f&longitude=%f&daily=apparent_temperature_max&timezone=GMT&start_date=%s&end_date=%s",
		ws.APIBase,
		lat,
		long,
		startDateStr,
		startDateStr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	result := struct {
		Latitude             float64 `json:"latitude"`
		Longitude            float64 `json:"longitude"`
		GenerationtimeMs     float64 `json:"generationtime_ms"`
		UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
		Timezone             string  `json:"timezone"`
		TimezoneAbbreviation string  `json:"timezone_abbreviation"`
		Elevation            float64 `json:"elevation"`
		DailyUnits           struct {
			Time             string `json:"time"`
			Temperature2MMax string `json:"temperature_2m_max"`
			Temperature2MMin string `json:"temperature_2m_min"`
		} `json:"daily_units"`
		Daily struct {
			Time                   []string  `json:"time"`
			ApparentTemperatureMax []float64 `json:"apparent_temperature_max"`
		} `json:"daily"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	for i, _ := range result.Daily.Time {
		date, err := time.Parse("2006-01-02", result.Daily.Time[i])
		if err != nil {
			return nil, err
		}
		data := repository.Temperature{
			Location: location,
			Value:    result.Daily.ApparentTemperatureMax[i],
			Date:     date,
		}
		return &data, nil
	}
	return nil, nil
}

func getCoordinates(location string) (float64, float64) {

	return 52.52, 13.419998
}
