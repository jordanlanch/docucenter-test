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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserTable)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
