package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/basiliuswicaksono/Cake-Store/config"
	"github.com/basiliuswicaksono/Cake-Store/controllers"
	"github.com/basiliuswicaksono/Cake-Store/models"
	"github.com/basiliuswicaksono/Cake-Store/params"
	"github.com/basiliuswicaksono/Cake-Store/repositories"
	"github.com/basiliuswicaksono/Cake-Store/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDb() config.Store {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Migrate(db config.Store) {
	err := db.MigrateDB()
	if err != nil {
		log.Println("Failed to run migrations")
		panic(err.Error())
	}
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetupController(db config.Store) *controllers.CakeController {
	cakeRepo1 := repositories.NewCakeRepo(db.DB)
	cakeService1 := services.NewCakeService(cakeRepo1)
	cakeController1 := controllers.NewCakeContoller(*cakeService1)

	return cakeController1
}

func TestAddNewCake(t *testing.T) {
	db := SetupDb()
	Migrate(db)
	cakeController := SetupController(db)

	cake := params.CakeRequest{
		Title:       "title test",
		Description: "description test",
		Rating:      5.0,
		Image:       "image test",
	}
	data, err := json.Marshal(cake)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)

	r := SetUpRouter()
	r.POST("/cakes", cakeController.AddNewCake)
	req, _ := http.NewRequest("POST", "/cakes", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetListOfCakes(t *testing.T) {
	db := SetupDb()
	cakeController := SetupController(db)

	r := SetUpRouter()
	r.GET("/cakes", cakeController.GetListOfCakes)
	req, _ := http.NewRequest("GET", "/cakes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var cakes []models.Cake
	json.Unmarshal(w.Body.Bytes(), &cakes)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, cakes)
}
