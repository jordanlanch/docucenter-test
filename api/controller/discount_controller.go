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

	filters := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		if k != "limit" && k != "offset" {
			filters[k] = v[0]
		}
	}

	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	discounts, err := pc.DiscountUsecase.GetMany(pagination, filters)
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

// Create - Create a new discount
func (pc *DiscountController) Create(c *gin.Context) {
	var discount domain.Discount

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	discountSchemaReflect := r.Reflect(domain.Discount{})
	schemaLoader := gojsonschema.NewGoLoader(&discountSchemaReflect)

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

	if err := json.Unmarshal(requestBody, &discount); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newDiscount, err := pc.DiscountUsecase.Create(&discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newDiscount)
}

// Modify - Update an existing discount
func (pc *DiscountController) Modify(c *gin.Context) {
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

	var discount domain.Discount

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	discountSchemaReflect := r.Reflect(domain.Discount{})
	schemaLoader := gojsonschema.NewGoLoader(&discountSchemaReflect)

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

	if err := json.Unmarshal(requestBody, &discount); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updateDiscount, err := pc.DiscountUsecase.Modify(&discount, id)
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
