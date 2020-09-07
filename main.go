package main

import (
	"log"
	"time"

	"gotrading/app/controllers"
	"gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	df, _ := models.GetAllCandle(config.Config.ProductCode, time.Minute, 365)
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
