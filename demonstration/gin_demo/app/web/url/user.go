package url

import (

	"github.com/gin-gonic/gin"
	"go-integrated/demonstration/gin_demo/app/web/handler/user"
)

func userInclude(c *gin.RouterGroup) {
	baseUrl := c.Group("")

	baseUrl.GET("/:userId", user.Get)

	fansUrl := c.Group("/fans")
	fansUrl.GET("", user.Get)

}
