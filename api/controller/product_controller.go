package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

// func (pc *ProductController) Fetch(c *gin.Context) {
// 	limitStr := c.DefaultQuery("limit", "10")
// 	limit, err := strconv.Atoi(limitStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit parameter"})
// 		return
// 	}

// 	offsetStr := c.DefaultQuery("offset", "0")
// 	offset, err := strconv.Atoi(offsetStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid offset parameter"})
// 		return
// 	}

// 	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

// 	meters, err := pc.ProductUsecase.GetManyProducts(pagination)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, meters)
// }

func (pc *ProductController) Get(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing id parameter"})
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	meter, err := pc.ProductUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meter)
}

// Create - Create a new customer
func (pc *ProductController) Create(c *gin.Context) {
	var customer domain.Product
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.ProductUsecase.Create(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// Modify - Update an existing customer
func (pc *ProductController) Modify(c *gin.Context) {
	var customer domain.Product
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.ProductUsecase.Modify(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Remove - Delete a customer by ID
func (pc *ProductController) Remove(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing id parameter"})
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = pc.ProductUsecase.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
