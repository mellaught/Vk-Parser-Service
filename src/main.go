package main

import (
	"fmt"

	"github.com/mrKitikat/Vk-Parser-Service/src/app"
	"github.com/mrKitikat/Vk-Parser-Service/src/app/models"
	"github.com/mrKitikat/Vk-Parser-Service/src/config"
)

func main() {

	cfg := config.NewViperConfig()

	conf := &models.Config{
		Token:       cfg.GetStringSlice("vk.token"),
		Version:     cfg.GetString("vk.version"),
		URL:         cfg.GetString("vk.URL"),
		ServiceHost: cfg.GetString("server.host"),
		ServicePort: cfg.GetString("server.port"),
	}

	// Create New App with params from conf.
	app := app.NewApp(conf)

	// Start App.
	fmt.Println("App has started")
	app.Run(conf.ServiceHost + ":" + conf.ServicePort)
}
