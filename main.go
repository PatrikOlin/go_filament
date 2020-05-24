package main

import (
	"os"
	"io"
	"time"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/PatrikOlin/go_filament/models"
	"github.com/PatrikOlin/go_filament/controllers"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	
	t := time.Now()

	logpath := filepath.Join(".", "logs")
	os.MkdirAll(logpath, os.ModePerm)
	f, _ := os.OpenFile(logpath + "/spools_gin" + t.Format("20060102")  + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.Writer(f)

}

func main() {
	Init()

	r := gin.Default()

	db := models.SetupModels()

	
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
