package main

import "gin_serve/cmd"

// @title           Gin Serve Api
// @version         0.1.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/freeshineit/gin_serve

// @contact.name   API Support
// @contact.url    https://github.com/freeshineit/gin_serve
// @contact.email  xiaoshaoqq@gmail.com

// @license.name  MIT
// @license.url   https://github.com/freeshineit/gin_serve/blob/main/LICENSE

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	cmd.Execute()
}
