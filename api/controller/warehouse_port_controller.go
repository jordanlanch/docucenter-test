package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/invopop/jsonschema"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stoewer/go-strcase"
	"github.com/xeipuuv/gojsonschema"
)

type WarehousePortController struct {
	WarehousePortUsecase domain.WarehousePortUsecase
}

func (mc *WarehousePortController) Fetch(c *gin.Context) {
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

	filters := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		if k != "limit" && k != "offset" {
			filters[k] = v[0]
		}
	}

	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	warehousesPorts, err := mc.WarehousePortUsecase.GetMany(pagination, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, warehousesPorts)
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

	warehousePort, err := pc.WarehousePortUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, warehousePort)
}

// Create - Create a new warehousePort
func (pc *WarehousePortController) Create(c *gin.Context) {
	var warehousePort domain.WarehousesAndPorts

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	warehousePortSchemaReflect := r.Reflect(domain.WarehousesAndPorts{})
	schemaLoader := gojsonschema.NewGoLoader(&warehousePortSchemaReflect)

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	documentLoader := gojsonschema.NewBytesLoader(requestBody)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if !result.Valid() {
		err = fmt.Errorf("[ERROR] invalid payload %+v", result.Errors())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &warehousePort); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newWarehousePort, err := pc.WarehousePortUsecase.Create(&warehousePort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newWarehousePort)
}

// Modify - Update an existing warehousePort
func (pc *WarehousePortController) Modify(c *gin.Context) {
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

	var warehousePort domain.WarehousesAndPorts

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	warehousePortSchemaReflect := r.Reflect(domain.WarehousesAndPorts{})
	schemaLoader := gojsonschema.NewGoLoader(&warehousePortSchemaReflect)

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	documentLoader := gojsonschema.NewBytesLoader(requestBody)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if !result.Valid() {
		err = fmt.Errorf("[ERROR] invalid payload %+v", result.Errors())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &warehousePort); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updateWarehousePort, err := pc.WarehousePortUsecase.Modify(&warehousePort, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateWarehousePort)
}

// Remove - Delete a warehousePort by ID
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
