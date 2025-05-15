package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Swag(prefix string, r *gin.Engine) {
	// docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix

	// Swagger documentation route
	r.GET(prefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect /swagger to /swagger/index.html
	swagRedirectHandler := func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, prefix+"/swagger/index.html")
	}
	r.GET(prefix+"/swagger", swagRedirectHandler)
	r.GET(prefix+"/api", swagRedirectHandler)
	r.GET(prefix+"/", swagRedirectHandler)
}
