package main

import (
	"github.com/ParasJain0307/generate-validate-otp/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app := api.Config{
		Router: router,
	}

	app.Routes()
	router.Run(":8000")
}
