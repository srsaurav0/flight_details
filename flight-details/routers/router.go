package routers

import (
	"flight-details/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// Route for flight search
	web.Router("/api/v1/flights/search", &controllers.FlightController{}, "get:GetByAllParams")

	// Register the Swagger UI endpoint
	web.Router("/swagger/*", &controllers.SwaggerController{})
}
