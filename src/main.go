package main

import (
	"VkParser/src/app"
	"VkParser/src/app/models"
	"VkParser/src/config"
	"fmt"
)

func main() {

	cfg := config.NewViperConfig()

	conf := &models.Config{
		Token:       cfg.GetString("vk.token"),
		Version:     cfg.GetString("vk.version"),
		URL:         cfg.GetString("vk.URL"),
		ServiceHost: cfg.GetString("server.host"),
		ServicePort: cfg.GetString("server.port"),
		Timeout:     cfg.GetInt("timeouts.balance"),
	}

	// Create New App with params from conf.
	app := app.NewApp(conf)

	// Start App.
	fmt.Println("App has started")
	app.Run(conf.ServiceHost + ":" + conf.ServicePort)
}
