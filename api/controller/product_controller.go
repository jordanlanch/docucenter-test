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

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) Fetch(c *gin.Context) {
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

	products, err := pc.ProductUsecase.GetMany(pagination, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

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

	product, err := pc.ProductUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Create - Create a new product
func (pc *ProductController) Create(c *gin.Context) {
	var product domain.Product

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	productSchemaReflect := r.Reflect(domain.Product{})
	schemaLoader := gojsonschema.NewGoLoader(&productSchemaReflect)

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

	if err := json.Unmarshal(requestBody, &product); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newProduct, err := pc.ProductUsecase.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

// Modify - Update an existing product
func (pc *ProductController) Modify(c *gin.Context) {
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

	var product domain.Product

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	productSchemaReflect := r.Reflect(domain.Product{})
	schemaLoader := gojsonschema.NewGoLoader(&productSchemaReflect)

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

	if err := json.Unmarshal(requestBody, &product); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updateProduct, err := pc.ProductUsecase.Modify(&product, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateProduct)
}

// Remove - Delete a product by ID
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
