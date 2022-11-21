package main

import (
	"log"

	"github.com/basiliuswicaksono/Cake-Store/config"
	"github.com/basiliuswicaksono/Cake-Store/controllers"
	"github.com/basiliuswicaksono/Cake-Store/repositories"
	"github.com/basiliuswicaksono/Cake-Store/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Run() error {
	// responsible for initializing and starting
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}

	// db.DB.Close()

	err = db.MigrateDB()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	route := gin.Default()

	cakeRepo := repositories.NewCakeRepo(db.DB)
	cakeService := services.NewCakeService(cakeRepo)
	cakeController := controllers.NewCakeContoller(*cakeService)

	mainRouter := route.Group("/")
	{
		mainRouter.GET("/cakes", cakeController.GetListOfCakes)
		mainRouter.GET("/cakes/:id", cakeController.GetCakeDetail)
		mainRouter.POST("/cakes", cakeController.AddNewCake)
		mainRouter.PATCH("/cakes/:id", cakeController.UpdateCake)
		mainRouter.DELETE("/cakes/:id", cakeController.DeleteCake)
	}

	route.Run(":4000")

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
