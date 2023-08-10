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

func NewLogisticsRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	llr := repository.NewLogisticsRepository(db, domain.LogisticsTable)
	dr := repository.NewDiscountRepository(db, domain.DiscountTable)
	llc := &controller.LogisticsController{
		LogisticsUsecase: usecase.NewLogisticsUsecase(llr, dr, timeout),
	}

	group.POST("/logistics", llc.Create)
	group.GET("/logistics", llc.Fetch)
	group.GET("/logistics/:id", llc.Get)
	group.PUT("/logistics/:id", llc.Modify)
	group.DELETE("/logistics/:id", llc.Remove)
}
