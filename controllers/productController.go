package controllers

import (
	"hewantani/httperror"
	"hewantani/models"
	"hewantani/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

type ProductInput struct {
	Name        string   `json:"name" binding:"required"`
	StoreId     uint     `json:"store_id" binding:"required"`
	Count       uint     `json:"count" binding:"required"`
	Price       uint     `json:"price" binding:"required"`
	ImageUrl    string   `json:"image_url" binding:"required,url"`
	Description string   `json:"description" binding:"required"`
	Categories  []string `json:"categories" binding:"required"`
}

// CreateProduct godoc
// @Summary      Create Product, user role must be MERCHANT
// @Description  create product
// @Tags         Product
// @Param        Body  body  ProductInput  true  "the body to create a Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Product}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /products [post]
func (h ProductController) CreateProduct(c *gin.Context) {
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	m := models.Product{}
	m.Name = input.Name
	m.StoreId = input.StoreId
	m.Count = input.Count
	m.Price = input.Price
	m.ImageUrl = input.ImageUrl
	m.Description = input.Description
	m.UserId = c.MustGet("user_id").(uint)
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
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedProduct})
}

type updateProductInput struct {
	Name        string   `json:"name" binding:"required"`
	Count       uint     `json:"count" binding:"required"`
	Price       uint     `json:"price" binding:"required"`
	ImageUrl    string   `json:"image_url" binding:"required,url"`
	Description string   `json:"description" binding:"required"`
	Categories  []string `json:"categories" binding:"required"`
}

// UpdateProduct godoc
// @Summary      Update Product, user role must be MERCHANT
// @Description  update  product
// @Tags         Product
// @Param id path string true "product id"
// @Param        Body  body  updateProductInput  true  "the body to update a Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Product}
// @Success      400  {object}  models._Err
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /products/{id} [put]
func (h ProductController) UpdateProduct(c *gin.Context) {
	var input updateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	productId := c.MustGet("product_id").(uint)
	found := c.MustGet("found").(*models.Product)

	found.Name = input.Name
	found.Count = input.Count
	found.Price = input.Price
	found.ImageUrl = input.ImageUrl
	found.Description = input.Description
	for _, v := range input.Categories {
		mCategory, _ := services.All.CategoryService.Find(v)
		if mCategory.Name != "" {
			found.Categories = append(found.Categories, *mCategory)
			continue
		}
		found.Categories = append(found.Categories, models.Category{Name: v})
	}

	savedProduct, err := services.All.ProductService.Update(productId, found)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedProduct})
}

// DeleteProduct godoc
// @Summary      delete product, user role must be MERCHANT
// @Description  delete product
// @Tags         Product
// @Param id path string true "product id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce      json
// @Success      200  {object}  models._Res
// @Success      401  {object}  models._Err
// @Success      403  {object}  models._Err
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /products/{id} [delete]
func (h ProductController) DeleteProduct(c *gin.Context) {
	productId := c.MustGet("product_id").(uint)

	err := services.All.ProductService.Delete(productId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success deleted product"})
}

// GetProducts godoc
// @Summary      get products, anyone can access
// @Description  get products
// @Tags         Product
// @Param        categories    query     []string  false  "filter by categories"  collectionFormat(multi)
// @Param        keyword    query     string  false  "filter by keyword"
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Product}
// @Success      500  {object}  models._Err
// @Router       /products [get]
func (h ProductController) GetProducts(c *gin.Context) {
	categories := c.QueryArray("categories")
	keyword := c.Query("keyword")
	products, err := services.All.ProductService.FindAll(categories, keyword)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": products})
}

// GetProduct godoc
// @Summary      get product, anyone can access
// @Description  get product by id
// @Tags         Product
// @Param id path string true "product id"
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Product}
// @Success      404  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /products/{id} [get]
func (h ProductController) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	productId, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}
	products, err := services.All.ProductService.FindById(uint(productId))
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": products})
}
