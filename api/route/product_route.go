package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/docucenter-test/api/controller"
	"github.com/jordanlanch/docucenter-test/bootstrap"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/repository"
	"github.com/jordanlanch/docucenter-test/usecase"
	"gorm.io/gorm"
)

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	pr := repository.NewProductRepository(db, domain.ProductTable)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
	}

	group.POST("/products", pc.Create)
	group.GET("/products/:id", pc.Get)
	group.PUT("/products/:id", pc.Modify)
	group.DELETE("/products/:id", pc.Modify)
}
