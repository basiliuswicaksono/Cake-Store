package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/basiliuswicaksono/Cake-Store/params"
	"github.com/basiliuswicaksono/Cake-Store/repositories"
)

type CakeService struct {
	cakeRepo repositories.CakeRepo
}

func NewCakeService(cakeRepo repositories.CakeRepo) *CakeService {
	return &CakeService{
		cakeRepo,
	}
}

func (c *CakeService) GetListOfCakes() *params.Response {
	cakes, err := c.cakeRepo.GetListOfCakes()
	if err != nil {
		log.Println("ERROR:", err)
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: cakes,
	}
}

func (c *CakeService) GetCakeDetail(cakeId int) *params.Response {
	cake, err := c.cakeRepo.GetCakeDetail(cakeId)
	if err != nil {
		log.Println("ERROR:", err)
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	if cake == nil {
		return &params.Response{
			Status:  http.StatusOK,
			Payload: "Record not found",
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: cake,
	}
}

func (c *CakeService) AddNewCake(req params.CakeRequest) *params.Response {
	err := c.cakeRepo.AddNewCake(req)
	if err != nil {
		log.Println("ERROR:", err)
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: "New cake was added.",
	}
}

func (c *CakeService) UpdateCake(cakeId int, req params.UpdateCakeRequest) *params.Response {
	err := c.cakeRepo.UpdateCake(cakeId, req)
	if err != nil {
		log.Println("ERROR:", err)
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: fmt.Sprintf("Cake with ID = %d was updated", cakeId),
	}
}

func (c *CakeService) DeleteCake(cakeId int) *params.Response {
	err := c.cakeRepo.DeleteCake(cakeId)
	if err != nil {
		log.Println("ERROR:", err)
		return &params.Response{
			Status: http.StatusInternalServerError,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: fmt.Sprintf("Cake with ID = %d was deleted", cakeId),
	}
}
