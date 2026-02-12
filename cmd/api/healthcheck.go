package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(c *gin.Context) {
	data := map[string]string{
		"status":  "ok",
		"version": version,
	}

	c.JSON(http.StatusOK, data)
}
