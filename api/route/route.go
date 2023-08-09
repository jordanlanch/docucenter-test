package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/docucenter-test/api/middleware"
	"github.com/jordanlanch/docucenter-test/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB) *gin.Engine {
	router := gin.New()
	publicRouter := router.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := router.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Private APIs
	NewCustomerRouter(env, timeout, db, protectedRouter)
	NewDiscountRouter(env, timeout, db, protectedRouter)
	NewLogisticsRouter(env, timeout, db, protectedRouter)
	NewProductRouter(env, timeout, db, protectedRouter)
	NewWarehousePortRouter(env, timeout, db, protectedRouter)

	return router
}
