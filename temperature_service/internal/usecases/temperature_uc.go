package usecases

import (
	"context"
	"github.com/haseebh/weatherapp_temperature/internal/entities/repository"
	"github.com/haseebh/weatherapp_temperature/pkg/config"
	"log"
	"time"
)

type TemperatureUseCase interface {
	FetchAndStoreTemperature(ctx context.Context, location string, startTime time.Time, endTime time.Time) error
	GetTemperatureByPeriod(ctx context.Context, location string, start, end time.Time) ([]*repository.Temperature, error)
}

type temperatureUC struct {
	tempRepo   repository.TemperatureRepository
	weatherSvc WeatherUseCase
}

func NewTemperatureUseCase(repo repository.TemperatureRepository) TemperatureUseCase {

	cfg := config.LoadConfig()
	return &temperatureUC{
		tempRepo:   repo,
		weatherSvc: NewWeatherService(cfg.WeatherAPIKey, cfg.WeatherAPIBaseURL),
	}
}

func (uc *temperatureUC) FetchAndStoreTemperature(ctx context.Context, location string, startTime time.Time, endTime time.Time) error {
	current := startTime
	for current.Before(endTime) || current.Equal(endTime) {
		err := uc.storeTemperature(ctx, location, current)
		if err != nil {
			log.Print(err)
		}
		current = current.AddDate(0, 0, 1) // Move to the next day
	}
	return nil
	/*	batchSize := 100
		var wg sync.WaitGroup
		var mu sync.Mutex
		var errorsResults []error

		// Calculate duration between start and end dates
		duration := endTime.Sub(startTime)
		// Calculate total number of days
		totalDays := int(duration.Hours() / 24)
		for i := 0; i < totalDays; i += batchSize {
			wg.Add(1)
			end := i + batchSize
			if end > totalDays {
				end = totalDays
			}
			sem := semaphore.NewWeighted(int64(batchSize))
			go func(wg *sync.WaitGroup, errors *[]error, mu *sync.Mutex) {
				defer wg.Done()
				dateStr, _ := time.Parse("2006-01-02", startTime.Format("2006-01-02"))
				sem.Acquire(context.TODO(), 1)
				go func(sem *semaphore.Weighted, dateStr time.Time) {
					defer sem.Release(1)
					err := uc.storeTemperature(ctx, location, dateStr)
					if err != nil {
						mu.Lock()
						*errors = append(*errors, err)
						mu.Unlock()
						log.Print(err)
					}

				}(sem, dateStr)
				startTime = startTime.AddDate(0, 0, 1) // Move to the next day
			}(&wg, &errorsResults, &mu)
		}
		wg.Wait()

		errResult := ""
		for _, err := range errorsResults {
			errResult += err.Error()
		}
		if errResult != "" {
			return errors.New(errResult)
		}
		return nil*/
}
func (uc *temperatureUC) storeTemperature(ctx context.Context, location string, date time.Time) error {
	temperatureResult, err := uc.tempRepo.GetTemperature(location, date)
	if err == nil && temperatureResult == nil {
		temperature, err := uc.weatherSvc.FetchTemperature(location, date)
		if err != nil {
			return err
		}
		_, err = uc.tempRepo.InsertTemperature(ctx, temperature)
		return err
	}
	return nil
}
func (uc *temperatureUC) GetTemperatureByPeriod(ctx context.Context, location string, start, end time.Time) ([]*repository.Temperature, error) {
	start, _ = time.Parse("2006-01-02", start.Format("2006-01-02"))
	end, _ = time.Parse("2006-01-02", end.Format("2006-01-02"))
	return uc.tempRepo.GetTemperaturesByPeriod(location, start, end)
}
