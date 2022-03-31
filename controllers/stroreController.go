package controllers

import (
	"hewantani/models"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
}

type StoreInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

// CreateStore godoc
// @Summary      Create Store, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Store
// @Param        Body  body  StoreInput  true  "the body to create a Store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /stores [post]
func (h StoreController) CreateStore(c *gin.Context) {
	var input StoreInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Store{}
	s.Name = input.Name
	s.Description = input.Description
	s.Address = input.Address
	s.UserId = c.MustGet("user_id").(uint)

	savedStore, err := services.All.StoreService.Save(&s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := map[string]string{
		"name":        savedStore.Name,
		"description": savedStore.Description,
		"address":     savedStore.Address,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}

// GetStores godoc
// @Summary      get stores, anyone can use this
// @Description  registering a user from public access.
// @Tags         Store
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /Stores [get]
func (h StoreController) GetStores(c *gin.Context) {
	data, err := services.All.StoreService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// UpdateStore godoc
// @Summary      Update Store, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Store
// @Param        Body  body  StoreInput  true  "the body to update a Store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /stores [put]
func (h StoreController) UpdateStore(c *gin.Context) {
	var input StoreInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := models.Store{}
	m.Name = input.Name
	m.Description = input.Description
	m.Address = input.Address
	m.UserId = c.MustGet("user_id").(uint)

	savedStore, err := services.All.StoreService.Update(uint(id), &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedStore})
}

// DeleteStore godoc
// @Summary      delete Store, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Store
// @Param        Body  body  StoreInput  true  "the body to delete a Store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /stores [delete]
func (h StoreController) DeleteStore(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedStore, err := services.All.StoreService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedStore})
}
