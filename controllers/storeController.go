package controllers

import (
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
}

type StoreInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required,url"`
}

// CreateStore godoc
// @Summary      Create Store, user role must be MERCHANT
// @Description  Create Store
// @Tags         Store
// @Param        Body  body  StoreInput  true  "the body to create a Store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Store}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /stores [post]
func (h StoreController) CreateStore(c *gin.Context) {
	var input StoreInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	s := models.Store{}
	s.Name = input.Name
	s.Description = input.Description
	s.Address = input.Address
	s.UserId = c.MustGet("user_id").(uint)
	s.ImageUrl = input.ImageUrl
	data, err := services.All.StoreService.Save(&s)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}

// GetStores godoc
// @Summary      get stores, anyone can use this
// @Description  get stores
// @Tags         Store
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Store}
// @Success      500  {object}  models._Err
// @Router       /stores [get]
func (h StoreController) GetStores(c *gin.Context) {
	data, err := services.All.StoreService.FindAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// UpdateStore godoc
// @Summary      Update Store, user role must be MERCHANT and must own the store
// @Description  update store
// @Tags         Store
// @Param id path string true "store id"
// @Param        Body  body  StoreInput  true  "the body to update a Store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Store}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /stores/{id} [put]
func (h StoreController) UpdateStore(c *gin.Context) {
	var input StoreInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	storeId := c.MustGet("store_id").(uint)
	found := c.MustGet("found").(*models.Store)

	found.Name = input.Name
	found.Description = input.Description
	found.Address = input.Address
	found.UserId = c.MustGet("user_id").(uint)
	found.ImageUrl = input.ImageUrl

	savedStore, err := services.All.StoreService.Update(uint(storeId), found)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedStore})
}

// DeleteStore godoc
// @Summary      delete Store, user role must be MERCHANT and must own the store
// @Description  delete store
// @Tags         Store
// @Param id path string true "store id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /stores/{id} [delete]
func (h StoreController) DeleteStore(c *gin.Context) {
	storeId := c.MustGet("store_id").(uint)

	err := services.All.StoreService.Delete(uint(storeId))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success deleted store"})
}
