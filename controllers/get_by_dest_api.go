package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"flight-details/services"
)

type FlightController struct {
	web.Controller
}

// @Summary Search for flights
// @Description Search for available flights based on destination and date
// @Tags Flights
// @Accept json
// @Produce json
// @Param DestCityName query string true "Destination city name" example:"London"
// @Param timestamp query string true "Flight date" example:"2024-02-10"
// @Success 200 {array} models.Flight "List of flights"
// @Failure 400 {object} models.ErrorResponse "Bad request - Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /flights [get]
func (c *FlightController) Get() {
	destination := c.GetString("DestCityName") 
	date := c.GetString("timestamp")          

	flights, err := services.SearchFlights(destination, date)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = flights
	}
	c.ServeJSON()
}