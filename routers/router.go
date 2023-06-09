package routers

import (
	"VisitorsManagementSystem/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	V1 := router.Group("/admin")
	{
		V1.POST("/login",controllers.Login)
		V1.GET("/getLogin",controllers.GetLogin)
		V1.POST("/select",controllers.Select)
		V1.GET("/find",controllers.Find)
		V1.GET("/delete/:id",controllers.Delete)
		V1.GET("/edit",controllers.Edit)
		V1.POST("/update",controllers.Update)
	}
	V2 := router.Group("/visitor")
	{
		V2.POST("/add",controllers.AddVisitor)
		V2.GET("/getAdd",controllers.GetAdd)
	}
	return router
}
