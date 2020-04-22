package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jwma/jump-jump/internal/app/config"
	"github.com/jwma/jump-jump/internal/app/models"
	"net/http"
)

//GetConfig 获取默认配置信息
func GetConfig(c *gin.Context) {
	cfg := config.GetConfig()
	landingHosts := cfg.GetStringSliceValue("landingHosts", make([]string, 0))

	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"code": 0,
		"data": gin.H{
			"config": gin.H{"landingHosts": landingHosts},
		},
	})
}

type LandingHostsParameter struct {
	Hosts []string `json:"hosts"`
}

//UpdateLandingHostsAPI 更新默认配置信息
func UpdateLandingHostsAPI() gin.HandlerFunc {
	return Authenticator(func(c *gin.Context, user *models.User) {
		if user.Role != models.RoleAdmin {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "你无权修改短链接域名",
				"code": 4999,
				"data": nil,
			})
			return
		}

		p := &LandingHostsParameter{}
		if err := c.ShouldBindJSON(p); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  err.Error(),
				"code": 4999,
				"data": nil,
			})
			return
		}

		cfg := config.GetConfig()
		_ = cfg.SetValue("landingHosts", p.Hosts)
		_ = cfg.Persist()

		c.JSON(http.StatusOK, gin.H{
			"msg":  "",
			"code": 0,
			"data": gin.H{
				"config": gin.H{"landingHosts": cfg.GetStringSliceValue("landingHosts", make([]string, 0))},
			},
		})
	})
}
