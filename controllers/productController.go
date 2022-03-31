package controllers

import (
	"hewantani/models"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

type ProductInput struct {
	Name       string   `json:"name" binding:"required"`
	StoreId    uint     `json:"store_id" binding:"required"`
	Count      uint     `json:"count" binding:"required"`
	Price      uint     `json:"price" binding:"required"`
	Categories []string `json:"categories" binding:"required"`
}

// CreateProduct godoc
// @Summary      Create Product, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Product
// @Param        Body  body  ProductInput  true  "the body to create a Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /products [post]
func (h ProductController) CreateProduct(c *gin.Context) {
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := models.Product{}
	m.Name = input.Name
	m.StoreId = input.StoreId
	m.Count = input.Count
	m.Price = input.Price
	for _, v := range input.Categories {
		mCategory, _ := services.All.CategoryService.Find(v)
		if mCategory.Name != "" {
			m.Categories = append(m.Categories, *mCategory)
			continue
		}
		m.Categories = append(m.Categories, models.Category{Name: v})
	}

	savedProduct, err := services.All.ProductService.Save(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedProduct})
}

// UpdateProduct godoc
// @Summary      Update Product, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Product
// @Param        Body  body  ProductInput  true  "the body to update a Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /products [put]
func (h ProductController) UpdateProduct(c *gin.Context) {
	var input ProductInput

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

	m := models.Product{}
	m.Name = input.Name
	m.StoreId = input.StoreId
	m.Count = input.Count
	m.Price = input.Price
	for _, v := range input.Categories {
		mCategory, _ := services.All.CategoryService.Find(v)
		if mCategory.Name != "" {
			m.Categories = append(m.Categories, *mCategory)
			continue
		}
		m.Categories = append(m.Categories, models.Category{Name: v})
	}

	savedProduct, err := services.All.ProductService.Update(uint(id), &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedProduct})
}

// DeleteProduct godoc
// @Summary      delete product, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Product
// @Param        Body  body  ProductInput  true  "the body to delete a Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /products [delete]
func (h ProductController) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedProduct, err := services.All.ProductService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedProduct})
}

// CreateProduct godoc
// @Summary      Create Product, user role must be MERCHANT
// @Description  registering a user from public access.
// @Tags         Product
// @Param        Body  body  ProductInput  true  "the body to create a Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /products [get]
func (h ProductController) GetProducts(c *gin.Context) {

	products, err := services.All.ProductService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": products})
}
