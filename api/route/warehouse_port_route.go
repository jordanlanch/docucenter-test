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

func NewWarehousePortRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	wpr := repository.NewWarehousePortRepository(db, domain.WarehousePortTable)
	wpc := &controller.WarehousePortController{
		WarehousePortUsecase: usecase.NewWarehousePortUsecase(wpr, timeout),
	}

	group.POST("/warehouse_ports", wpc.Create)
	group.GET("/warehouse_ports/:id", wpc.Get)
	group.PUT("/warehouse_ports/:id", wpc.Modify)
	group.DELETE("/warehouse_ports/:id", wpc.Modify)
}
