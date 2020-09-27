package main

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/PatrikOlin/go_filament/controllers"
	"github.com/PatrikOlin/go_filament/models"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)

	t := time.Now()

	logpath := filepath.Join(".", "logs")
	os.MkdirAll(logpath, os.ModePerm)
	f, _ := os.OpenFile(logpath+"/spools_gin"+t.Format("20060102")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	r.Use(cors.Default())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("v1/")
	{
		v1.GET("/spools", controllers.GetAllSpools)
		v1.GET("/spools/:tag", controllers.FindSpoolByTag)
		v1.POST("/spools", controllers.CreateSpool)
		// r.GET("/spools/:id", controllers.FindSpoolById)
		v1.PATCH("/spools/:tag", controllers.UpdateSpool)
		v1.DELETE("/spools/:tag", controllers.DeleteSpool)

		v1.GET("/brands", controllers.GetAllBrands)
		v1.GET("/brands/:name", controllers.FindBrandByName)
		v1.POST("/brands", controllers.CreateBrand)
		v1.PATCH("/brands/:name", controllers.UpdateBrand)
		v1.DELETE("/brands/:name", controllers.DeleteBrand)
	}

	r.Run()
}
