package main

import (
	_ "github.com/relato/gingogen/boilerplate/config"
	"github.com/relato/gingogen/boilerplate/handlers"
	"github.com/relato/gingogen/boilerplate/tasks"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
