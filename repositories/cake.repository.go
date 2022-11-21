package repositories

import (
	"database/sql"
	"fmt"

	"github.com/basiliuswicaksono/Cake-Store/models"
	"github.com/basiliuswicaksono/Cake-Store/params"
)

type CakeRepo interface {
	GetListOfCakes() ([]*models.Cake, error)
	GetCakeDetail(id int) (*models.Cake, error)
	AddNewCake(req params.CakeRequest) error
	UpdateCake(id int, req params.UpdateCakeRequest) error
	DeleteCake(id int) error
}

type cakeRepo struct {
	db *sql.DB
}

func NewCakeRepo(db *sql.DB) CakeRepo {
	return &cakeRepo{db}
}

func (c *cakeRepo) GetListOfCakes() ([]*models.Cake, error) {
	var cakes []*models.Cake = []*models.Cake{}

	result, err := c.db.Query("SELECT * from cakes order by rating asc, title asc")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var cake models.Cake
		err := result.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cakes = append(cakes, &cake)
	}

	return cakes, err
}

func (c *cakeRepo) GetCakeDetail(cakeId int) (*models.Cake, error) {
	var cake models.Cake

	result, err := c.db.Query("SELECT * FROM cakes WHERE id = ?", cakeId)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	if cake.ID == 0 {
		return nil, err
	}

	return &cake, err
}

func (c *cakeRepo) AddNewCake(req params.CakeRequest) error {
	stmt, err := c.db.Prepare("INSERT INTO cakes(title, description, rating, image) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(req.Title, req.Description, req.Rating, req.Image)
	if err != nil {
		return err
	}

	return nil
}

func (c *cakeRepo) UpdateCake(cakeId int, req params.UpdateCakeRequest) error {
	found, _ := c.GetCakeDetail(cakeId)
	if found == nil {
		return fmt.Errorf("cake with id %d, not found", cakeId)
	}

	stmt, err := c.db.Prepare("UPDATE cakes SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?")
	if err != nil {
		return err
	}

	if req.Title != "" {
		found.Title = req.Title
	}
	if req.Description != "" {
		found.Description = req.Description
	}
	if req.Rating != 0.0 {
		found.Rating = req.Rating
	}
	if req.Image != "" {
		found.Image = req.Image
	}

	_, err = stmt.Exec(found.Title, found.Description, found.Rating, found.Image, cakeId)
	if err != nil {
		return err
	}

	return nil
}

func (c *cakeRepo) DeleteCake(cakeId int) error {
	found, _ := c.GetCakeDetail(cakeId)
	if found == nil {
		return fmt.Errorf("cake with id %d, not found", cakeId)
	}

	stmt, err := c.db.Prepare("DELETE FROM cakes WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cakeId)
	if err != nil {
		return err
	}

	return nil
}
