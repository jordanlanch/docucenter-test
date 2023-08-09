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

func NewCustomerRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	cr := repository.NewCustomerRepository(db, domain.CustomerTable)
	cc := &controller.CustomerController{
		CustomerUsecase: usecase.NewCustomerUsecase(cr, timeout),
	}

	group.POST("/customers", cc.Create)
	group.GET("/customers/:id", cc.Get)
	group.PUT("/customers/:id", cc.Modify)
	group.DELETE("/customers/:id", cc.Modify)
}
