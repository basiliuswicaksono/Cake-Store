package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/basiliuswicaksono/Cake-Store/params"
	"github.com/basiliuswicaksono/Cake-Store/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CakeController struct {
	cakeService services.CakeService
}

func NewCakeContoller(cakeService services.CakeService) *CakeController {
	return &CakeController{cakeService}
}

func (c *CakeController) GetListOfCakes(ctx *gin.Context) {
	result := c.cakeService.GetListOfCakes()
	ctx.JSON(result.Status, result.Payload)
}

func (c *CakeController) GetCakeDetail(ctx *gin.Context) {
	cakeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": "Invalid id format",
			}})
		return
	}

	result := c.cakeService.GetCakeDetail(cakeId)
	ctx.JSON(result.Status, result.Payload)
}

func (c *CakeController) AddNewCake(ctx *gin.Context) {
	var request params.CakeRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := c.cakeService.AddNewCake(request)
	ctx.JSON(result.Status, result.Payload)
}

func (c *CakeController) UpdateCake(ctx *gin.Context) {
	cakeId, err := strconv.Atoi(ctx.Param("id"))
	var request params.UpdateCakeRequest

	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})
		return
	}

	result := c.cakeService.UpdateCake(cakeId, request)
	ctx.JSON(result.Status, result.Payload)
}

func (c *CakeController) DeleteCake(ctx *gin.Context) {
	cakeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": "Invalid id format",
			}})
		return
	}

	result := c.cakeService.DeleteCake(cakeId)
	ctx.JSON(result.Status, result.Payload)
}
