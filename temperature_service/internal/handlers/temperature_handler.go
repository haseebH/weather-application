package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haseebh/weatherapp_temperature/internal/usecases"
)

type TemperatureHandler struct {
	temperatureUseCase usecases.TemperatureUseCase
}

func NewTemperatureHandler(tempUC usecases.TemperatureUseCase) *TemperatureHandler {
	return &TemperatureHandler{temperatureUseCase: tempUC}
}

func (tc *TemperatureHandler) FetchTemperature(c *gin.Context) {
	location := c.Param("location")
	sDateStr := c.Query("startDate")
	sDate, err := time.Parse("2006-01-02", sDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}
	eDateStr := c.Query("endDate")
	eDate, err := time.Parse("2006-01-02", eDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}
	if err := tc.temperatureUseCase.FetchAndStoreTemperature(c, location, sDate, eDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "temperature fetched and stored"})
}

func (tc *TemperatureHandler) GetTemperatureByPeriod(c *gin.Context) {
	location := c.Param("location")
	period := c.Param("period")

	var start, end time.Time
	end = time.Now().UTC()

	switch period {
	case "month":
		start = end.AddDate(0, -1, 0).UTC()
	case "year":
		start = end.AddDate(-1, 0, 0).UTC()
	case "3years":
		start = end.AddDate(-3, 0, 0).UTC()
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid period"})
		return
	}

	temperatures, err := tc.temperatureUseCase.GetTemperatureByPeriod(c, location, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, temperatures)
}
