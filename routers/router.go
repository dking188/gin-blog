package routers

import (
	_ "net/http"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api/version1"
	_ "gin-blog/routers/api/version1"
	"github.com/gin-gonic/gin"

)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", version1.GetTags)
		//新建标签
		apiv1.POST("/tags", version1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id",version1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", version1.DeleteTag)
	}
	return r
}
