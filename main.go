package main

import (
	"log"
	"mealshare/models"
	"mealshare/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

)

func main() {
	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
    configuration, err := LoadConfig(".")
    if err != nil {
		log.Fatal("Error loading .env file")
	}
    
    // Initialize DB
    models.InitDB(configuration)

    // Load the routes
    routes.AuthRoutes(r)

    // Run the server
    r.Run(":8080")
}
 
func LoadConfig(path string) (config models.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}