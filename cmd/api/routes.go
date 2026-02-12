package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	// Healthcheck
	router.GET("/v1/healthcheck", app.healthcheckHandler)

	// Bioskop routes
	router.POST("/bioskop", app.createBioskopHandler)
	router.GET("/bioskop", app.getBioskopHandler)
	router.PUT("/bioskop/:id", app.updateBioskopHandler)
	router.DELETE("/bioskop/:id", app.deleteBioskopHandler)

	return router
}
