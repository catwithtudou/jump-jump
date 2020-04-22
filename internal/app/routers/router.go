package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jwma/jump-jump/internal/app/handlers"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//todo:gin中环境切换
	if gin.Mode() == gin.DebugMode { // 开发环境下，开启 CORS
		//todo:gin中自动解决跨域插件
		corsCfg := cors.DefaultConfig()
		corsCfg.AllowAllOrigins = true
		corsCfg.AddAllowHeaders("Authorization")
		r.Use(cors.New(corsCfg))
	}

	r.LoadHTMLFiles("./web/admin/index.html")
	r.StaticFS("/static", http.Dir("./web/admin/static"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	user:=r.Group("/v1/user")
	{
		user.POST("/login", handlers.Login)
		user.Use(handlers.JWTAuthenticatorMiddleware())
		user.GET("/info",  handlers.GetUserInfoAPI())
		user.POST("/logout", handlers.LogoutAPI())
		user.PATCH("/change-password", handlers.ChangePasswordAPI())
	}



	config:=r.Group("/v1/config")
	{
		config.GET("/", handlers.GetConfig)
		config.Use(handlers.JWTAuthenticatorMiddleware())
		config.PATCH("/",handlers.UpdateLandingHostsAPI())
	}



	shortLinkAPI := r.Group("/v1/short-link")
	shortLinkAPI.Use(handlers.JWTAuthenticatorMiddleware())
	shortLinkAPI.GET("/", handlers.ListShortLinksAPI())
	shortLinkAPI.GET("/:id", handlers.GetShortLinkAPI())
	shortLinkAPI.POST("/", handlers.CreateShortLinkAPI())
	shortLinkAPI.PATCH("/:id", handlers.UpdateShortLinkAPI())
	shortLinkAPI.DELETE("/:id", handlers.DeleteShortLinkAPI())
	shortLinkAPI.GET("/:id/*action", handlers.ShortLinkActionAPI())

	return r
}

//SetupLandingRouter 根据主机id访问不同主机
func SetupLandingRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/:id", handlers.Redirect)

	return r
}
