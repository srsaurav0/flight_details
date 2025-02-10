package routers

import (
	"flight-details/controllers"

	"github.com/beego/beego/v2/server/web"
)
 
func init() {
	web.Router("/flights", &controllers.FlightController{})
	web.Router("/swagger/*", &controllers.SwaggerController{})
}