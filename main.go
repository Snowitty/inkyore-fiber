package main

import (
	"log"

	"github.com/Snowitty/inkyore-fiber/database"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	database.ConnectDB()

}
