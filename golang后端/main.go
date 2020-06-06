package main

import (
	"gin/controllers"
	"gin/middleware/jwt"
	"gin/models"
	"gin/routers/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	r.POST("/auth", api.GetAuth)
	r.POST("/Fileupload", controllers.Fileupload)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/enterAssets", controllers.EnterAssets)
		apiv1.POST("/execTask", controllers.ExecTask)
		apiv1.POST("/getBaseInfo", controllers.GetBaseInfo)
		apiv1.POST("/download", controllers.DownloadInfo)
		apiv1.POST("/RemovePlugin", controllers.RemovePlugin)
		apiv1.POST("/GetPocNameAll", controllers.GetPocNameAll)
		apiv1.POST("/SingletonCheck", controllers.SingletonCheck)
		apiv1.POST("/SingletonCheckResult", controllers.SingletonCheckResult)
		apiv1.POST("/SelectSingkeAllInfo", controllers.SelectSingkeAllInfo)
		apiv1.POST("/getAllAssetsInfo", controllers.GetAllAssetsInfo)
		//apiv1.POST(
		apiv1.POST("/GetBugInfo", controllers.GetBugInfo)
		apiv1.POST("/getAssetsAddress", controllers.GetAssetsAddress)
		apiv1.POST("/removeAssets", controllers.RemoveAssets)
		apiv1.POST("/GetPocInfo", controllers.GetPocInfo)

	}
	r.Run(":80")
}
