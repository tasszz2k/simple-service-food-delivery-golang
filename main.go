package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-service-food-delivery-golang/component"
	"simple-service-food-delivery-golang/middleware"
	"simple-service-food-delivery-golang/modules/restaurant/restauranttransport/ginrestaurant"
	ginuser "simple-service-food-delivery-golang/modules/user/transport/gin"
)

func main() {
	db, err := createConnection()
	fmt.Println(db, err)
}

func createConnection() (*gorm.DB, error) {
	// Secret key for JWT
	secretKey := os.Getenv("SYSTEM_SECRET_KEY")

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

	if err := runService(db, secretKey); err != nil {
		log.Fatalln(err)
	}

	return db, err
}

func runService(db *gorm.DB, secretKey string) error {

	appCtx := component.NewAppContext(db, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	// Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ==================== CRUD =============================

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants")
	{
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run()
}
