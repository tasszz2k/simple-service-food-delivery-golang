package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-service-golang-04/modules/restaurant/restaurantmodel"
	"simple-service-golang-04/modules/restaurant/restauranttransport/ginrestaurant"
	"strconv"
)

func main() {
	db, err := createConnection()
	fmt.Println(db, err)
}

func createConnection() (*gorm.DB, error) {
	// Connect to MySQL
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

	return db, err
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ==================== CRUD =============================
	restaurants := r.Group("/restaurants")
	{
		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "id must be an integer",
				})
				return
			}

			var data restaurantmodel.Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "restaurant not found",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": data})

		})

		restaurants.GET("", func(c *gin.Context) {
			var data []restaurantmodel.Restaurant

			type Filter struct {
				CityId int `json:"city_id" form:"city_id"`
			}

			var filter Filter

			c.ShouldBind(&filter)

			newDb := db

			if filter.CityId > 0 {
				newDb = newDb.Where("city_id = ?", filter.CityId)
			}

			if err := newDb.Find(&data).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "restaurants not found",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": data})

		})

		restaurants.POST("", ginrestaurant.CreateRestaurant(db))

		restaurants.PUT("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "id must be an integer",
				})
				return
			}

			var data restaurantmodel.RestaurantUpdate

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "restaurant not found",
				})
				return
			}

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})

		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "id must be an integer",
				})
				return
			}

			var data restaurantmodel.Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "restaurant not found",
				})
				return
			}

			if err := db.Delete(&data).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": data})
		})

	}

	return r.Run()
}
