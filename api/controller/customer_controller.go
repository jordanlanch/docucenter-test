package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type CustomerController struct {
	CustomerUsecase domain.CustomerUsecase
}

// func (mc *CustomerController) Fetch(c *gin.Context) {
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

// 	meters, err := mc.CustomerUsecase.GetManyCustomers(pagination)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, meters)
// }

func (mc *CustomerController) Get(c *gin.Context) {
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

	meter, err := mc.CustomerUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meter)
}

// Create - Create a new customer
func (mc *CustomerController) Create(c *gin.Context) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := mc.CustomerUsecase.Create(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// Modify - Update an existing customer
func (mc *CustomerController) Modify(c *gin.Context) {
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := mc.CustomerUsecase.Modify(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Remove - Delete a customer by ID
func (mc *CustomerController) Remove(c *gin.Context) {
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

	err = mc.CustomerUsecase.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
