package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type WarehousePortController struct {
	WarehousePortUsecase domain.WarehousePortUsecase
}

func (pc *WarehousePortController) Get(c *gin.Context) {
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

	meter, err := pc.WarehousePortUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meter)
}

// Create - Create a new customer
func (pc *WarehousePortController) Create(c *gin.Context) {
	var customer domain.WarehousesAndPorts
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.WarehousePortUsecase.Create(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// Modify - Update an existing customer
func (pc *WarehousePortController) Modify(c *gin.Context) {
	var customer domain.WarehousesAndPorts
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.WarehousePortUsecase.Modify(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Remove - Delete a customer by ID
func (pc *WarehousePortController) Remove(c *gin.Context) {
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

	err = pc.WarehousePortUsecase.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "WarehousesAndPorts deleted successfully"})
}
