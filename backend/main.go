package main

import (
	"context"
	server "test_task/app"
	"test_task/config"
)

// @title News API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
func main() {
	config := config.NewDBConfig()

	app := server.Serve(context.Background(), *config)

	defer app.Shutdown()
	app.Listen(":8080")
}
