package test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/basiliuswicaksono/Cake-Store/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	return router
// }

// // cakeRepo1 := repositories.NewCakeRepo(db.DB)
// // cakeService1 := services.NewCakeService(cakeRepo1)
// // cakeController1 := controllers.NewCakeContoller(*cakeService1)

// func TestGetCakeHandler(t *testing.T) {
// 	r := SetUpRouter()
// 	r.GET("/cakes", cakeController.GetListOfCakes)
// 	req, _ := http.NewRequest("GET", "/cakes", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var cakes []*models.Cake
// 	json.Unmarshal(w.Body.Bytes(), &cakes)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEmpty(t, cakes)
// }
