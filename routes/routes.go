package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang-test/controller"
	"golang-test/middleware"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowAllOrigins = true
    corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

    // To be able to send tokens to the server.
    corsConfig.AllowCredentials = true
    // OPTIONS method for ReactJS
    corsConfig.AddAllowMethods("OPTIONS")

    r.Use(cors.New(corsConfig))

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    route := r.Group("/api/v1")

    // login
    route.POST("/register", controller.Register)
    route.POST("/login", controller.Login)
    
    // blog
    blogMiddlewareRoute := route.Group("/blogs")
    blogMiddlewareRoute.GET("", controller.GetAllBlog)
    blogMiddlewareRoute.GET("/:id", controller.GetBlogById)
	blogMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	{
		blogMiddlewareRoute.POST("", controller.CreateBlog)
		blogMiddlewareRoute.PUT("", controller.UpdateBlog)
		blogMiddlewareRoute.DELETE("/:id", controller.DeleteBlog)
	}
    return r
}