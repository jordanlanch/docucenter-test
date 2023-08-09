package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/docucenter-test/domain"
)

type DiscountController struct {
	DiscountUsecase domain.DiscountUsecase
}

func (pc *DiscountController) Get(c *gin.Context) {
	typeParam := c.Param("type")
	if typeParam == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing id parameter"})
		return
	}

	// Convertir string a LogisticsType
	var logisticsType domain.LogisticsType
	switch typeParam {
	case string(domain.Land):
		logisticsType = domain.Land
	case string(domain.Maritime):
		logisticsType = domain.Maritime
	default:
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid type parameter"})
		return
	}

	meter, err := pc.DiscountUsecase.GetByType(logisticsType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meter)
}

// Create - Create a new customer
func (pc *DiscountController) Create(c *gin.Context) {
	var customer domain.Discount
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.DiscountUsecase.Create(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// Modify - Update an existing customer
func (pc *DiscountController) Modify(c *gin.Context) {
	var customer domain.Discount
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.DiscountUsecase.Modify(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Remove - Delete a customer by ID
func (pc *DiscountController) Remove(c *gin.Context) {
	typeParam := c.Param("type")
	if typeParam == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing id parameter"})
		return
	}

	// Convertir string a LogisticsType
	var logisticsType domain.LogisticsType
	switch typeParam {
	case string(domain.Land):
		logisticsType = domain.Land
	case string(domain.Maritime):
		logisticsType = domain.Maritime
	default:
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid type parameter"})
		return
	}

	err := pc.DiscountUsecase.Remove(logisticsType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}
