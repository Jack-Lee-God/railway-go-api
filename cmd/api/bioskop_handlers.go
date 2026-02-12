package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"goapi.railway.app/internal/models"
)

func (app *application) createBioskopHandler(c *gin.Context) {
	nama := c.PostForm("nama")
	lokasi := c.PostForm("lokasi")
	ratingStr := c.PostForm("rating")

	if nama == "" || lokasi == "" || ratingStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama, Lokasi, dan Rating wajib diisi",
		})
		return
	}

	rating, err := strconv.ParseFloat(ratingStr, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rating harus berupa angka",
		})
		return
	}

	bioskop := models.Bioskop{
		Nama:   nama,
		Lokasi: lokasi,
		Rating: float32(rating),
	}

	if app.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}

	app.db.Create(&bioskop)

	c.JSON(http.StatusCreated, bioskop)
}

func (app *application) getBioskopHandler(c *gin.Context) {
	var bioskops []models.Bioskop

	if app.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}

	result := app.db.Find(&bioskops)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data",
		})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

func (app *application) updateBioskopHandler(c *gin.Context) {
	id := c.Param("id")

	var bioskop models.Bioskop

	if app.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}

	if err := app.db.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data bioskop tidak ditemukan",
		})
		return
	}

	nama := c.PostForm("nama")
	lokasi := c.PostForm("lokasi")
	ratingStr := c.PostForm("rating")

	if nama == "" || lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama and Lokasi tidak boleh kosong",
		})
		return
	}

	var rating float32 = bioskop.Rating
	if ratingStr != "" {
		value, err := strconv.ParseFloat(ratingStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Rating harus berupa angka",
			})
			return
		}
		rating = float32(value)
	}

	bioskop.Nama = nama
	bioskop.Lokasi = lokasi
	bioskop.Rating = rating

	app.db.Save(&bioskop)

	c.JSON(http.StatusOK, bioskop)
}

func (app *application) deleteBioskopHandler(c *gin.Context) {
	id := c.Param("id")

	var bioskop models.Bioskop

	if app.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}

	if err := app.db.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data bioskop tidak ditemukan",
		})
		return
	}

	if err := app.db.Delete(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menghapus data bioskop",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop deleted successfully",
	})
}
