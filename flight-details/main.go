
package main

import (
	_ "flight_details/routers"
	_ "flight_details/docs" 
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
