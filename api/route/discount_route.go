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

func NewDiscountRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	cr := repository.NewDiscountRepository(db, domain.DiscountTable)
	cc := &controller.DiscountController{
		DiscountUsecase: usecase.NewDiscountUsecase(cr, timeout),
	}

	group.POST("/discounts", cc.Create)
	group.GET("/discounts", cc.Fetch)
	group.GET("/discounts/:type/:quantity", cc.Get)
	group.PUT("/discounts/:id", cc.Modify)
	group.DELETE("/discounts/:id", cc.Remove)
}
