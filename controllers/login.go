package controllers

import (
	"VisitorsManagementSystem/logic"
	"VisitorsManagementSystem/middleware"
	"VisitorsManagementSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//登录
func Login(ctx *gin.Context) {
	//获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//数据验证
	if len(username) <= 0 {
		ctx.HTML(http.StatusUnprocessableEntity,"bad.html", gin.H{
			"code":    422,
			"message": "名字不能为空",
		})
		return
	}
	//验证是否有该用户
	if username != logic.SelectAdmin(username) {
		ctx.HTML(http.StatusUnprocessableEntity,"bad.html", gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	if len(password) < 6 {
		ctx.HTML(http.StatusUnprocessableEntity,"bad.html", gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}


	//判断密码是否正确
	if password != logic.SelectPassword(username) {
		ctx.HTML(http.StatusUnprocessableEntity,"bad.html", gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}


	user := models.Admin{
		Username: username,
		Password: password,
	}

	token ,err := middleware.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"code":"500",
			"msg":"系统异常",
		})

		log.Panicf("token generate error:%v",err)
		return
	}


	//返回结果
	result := logic.AllData()

	var users []Users

	for i:=0;i<=len(result)-1;i++{
		id := result[i].ID
		visistorName := result[i].VisitorName
		sex := result[i].Sex
		createdAt := result[i].CreatedAt
		event := result[i].Event
		var user Users
		user = Users{ID: id,VisitorName: visistorName,Sex: sex,CreatedAt: createdAt,Event: event}

		users = append(users,user)
	}

	ctx.HTML(http.StatusOK,"data.html",gin.H{
		"code":"200",
		"data":users,
		"msg":gin.H{"token":token},
		"message":"登录成功",
	})

}

func GetLogin(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}

type Users struct {
	ID uint `json:"ID"`
	VisitorName string `json:"VisitorName"`
	Sex string `json:"Sex"`
	CreatedAt time.Time `json:"CreatedAt"`
	Event string `json:"Event"`
}