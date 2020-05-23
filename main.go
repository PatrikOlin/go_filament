package main

import (
	"os"
	"io"
	"time"
	"github.com/gin-gonic/gin"

	"github.com/PatrikOlin/go_filament/models"
	"github.com/PatrikOlin/go_filament/controllers"
)


func main() {
	r := gin.Default()

	db := models.SetupModels()

	t := time.Now()
	f, _ := os.Create("spools_gin" + t.Format(time.RFC3339)  + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/spools", controllers.FindSpools)
	r.POST("/spools", controllers.CreateSpool)
	// r.GET("/spools/:id", controllers.FindSpoolById)
	r.GET("/spools/:tag", controllers.FindSpoolByTag)
	r.PATCH("/spools/:tag", controllers.UpdateSpool)
	r.DELETE("/spools/:tag", controllers.DeleteSpool)

	r.Run()
}
