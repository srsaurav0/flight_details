package main

import (
	_ "flight-details/docs"
	_ "flight-details/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
