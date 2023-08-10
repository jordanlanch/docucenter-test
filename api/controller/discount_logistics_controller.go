package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type DiscountController struct {
	DiscountUsecase domain.DiscountUsecase
}

func (pc *DiscountController) Fetch(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit parameter"})
		return
	}

	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid offset parameter"})
		return
	}

	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	discounts, err := pc.DiscountUsecase.GetMany(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, discounts)
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

	// Captura y convierte el parámetro quantity
	quantityParam := c.Param("quantity")
	quantity, err := strconv.Atoi(quantityParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid quantity parameter"})
		return
	}

	discount, err := pc.DiscountUsecase.GetByTypeAndQuantity(logisticsType, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, discount)
}

// Create - Create a new customer
func (pc *DiscountController) Create(c *gin.Context) {
	var discount domain.Discount
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newDiscount,err := pc.DiscountUsecase.Create(&discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newDiscount)
}

// Modify - Update an existing customer
func (pc *DiscountController) Modify(c *gin.Context) {
	var discount domain.Discount
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updateDiscount, err := pc.DiscountUsecase.Modify(&discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateDiscount)
}

// Remove - Delete a customer by ID
func (pc *DiscountController) Remove(c *gin.Context) {
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

	err = pc.DiscountUsecase.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}
