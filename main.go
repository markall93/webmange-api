package main

import (
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage"
	"time"
)

func main() {
	time.LoadLocation("Asia/Shanghai")
	webmanage.Init("")
	select {}

	//routers := router.InitRouter()
	//routers.Run("172.168.0.44:8087")
}