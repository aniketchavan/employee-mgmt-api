package main

import (
	"fmt"
	"log"
	"strconv"

	"example.com/employee-mgmt/config"
	"example.com/employee-mgmt/handlers"

	"example.com/employee-mgmt/persistence"
	"github.com/gin-gonic/gin"
)

func main() {

	// Read the Config YAML file
	config.InitializeConfigurations()

	// Check if port number is valid
	mySqlPort := 3306
	if portValue, isPresent := config.ConfigSet.Properties["MySqlPort"]; isPresent {
		portValue, err := strconv.Atoi(portValue)
		if err != nil {
			log.Fatalf("error while parsing DB port: %v", err)
		}
		mySqlPort = portValue
	}

	// Initialize MySQL DB connections
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.ConfigSet.Properties["MySqlUsername"], config.ConfigSet.Properties["MySqlPassword"], config.ConfigSet.Properties["MySqlHostname"], mySqlPort, config.ConfigSet.Properties["MySqlDbName"])
	if err := persistence.InitMySQL(dsn); err != nil {
		// To terminate the application if connection details are incorrect or any other error exists
		log.Fatalf("Could not connect to MySQL: %v", err)
	}

	// Initialize Redis
	persistence.InitRedis(config.ConfigSet.Properties["RedisServerURL"])

	r := gin.Default()

	// Define the upload route
	r.POST("/upload", handlers.UploadEmployeeData)

	// Define the get employee route
	r.GET("/employee/:email", handlers.GetEmployee)

	// Define the update employee route
	r.PUT("/employee/:email", handlers.UpdateEmployee)

	// Define the update employee route
	r.DELETE("/employee/:email", handlers.DeleteEmployee)

	// Start the server on port 8080
	r.Run(":8080")
}
