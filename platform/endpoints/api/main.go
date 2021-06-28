package main

import (
	"gochat/platform/endpoints/api/swagger/restapi"
	"gochat/platform/endpoints/api/swagger/restapi/operations"
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
)

//go:generate swagger generate server --target ./swagger/ --name gochat --spec ./swagger/swagger.yml --principal models.Principal --exclude-main

func main() {
	// Load the Embedded Swagger File
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Create a new API
	api := operations.NewGochatAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Configure the API
	server.ConfigureAPI()

	// Assign the Port Number
	port, err := strconv.Atoi(os.Getenv("API_SERVICE_PORT"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	server.Port = port

	// Start the Server
	log.Println("Starting API Server...")
	if err := server.Serve(); err != nil {
		log.Fatalln("Failed to start the API Server: ", err.Error())
	}
}
