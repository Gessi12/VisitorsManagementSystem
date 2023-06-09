package controllers

import (
	"VisitorsManagementSystem/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Select(c *gin.Context) {
	visitorName := c.PostForm("visitor_name")
	phone := c.PostForm("phone")
	if logic.SelectVisitorName(visitorName) == "" {
		c.HTML(http.StatusUnprocessableEntity,"selectDown.html", gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	if logic.SelectPhone(visitorName) != phone {
		c.HTML(http.StatusUnprocessableEntity,"selectDown.html", gin.H{
			"code":    422,
			"message": "手机号错误",
		})
		return
	}


	s := logic.SelectVisitor(visitorName,phone)

	var users []Visitors
	/*ID uint `json:"ID"`
	VisitorName string `json:"VisitorName"`
	Sex string `json:"Sex"`
	Phone string `json:"Phone"`
	VisitId string `json:"VisitId"`
	CreatedAt time.Time `json:"CreatedAt"`
	Event string `json:"Event"`*/
	for i:=0;i<=len(s)-1;i++{
		id := s[i].ID
		visistorName := s[i].VisitorName
		sex := s[i].Sex
		phone := s[i].Phone
		visitId := s[i].VisitId
		createdAt := s[i].CreatedAt
		event := s[i].Event
		var user Visitors
		user = Visitors{ID: id,VisitorName: visistorName,Sex: sex,Phone: phone,VisitId: visitId,CreatedAt: createdAt,Event: event}

		users = append(users,user)
	}


	c.HTML(http.StatusOK,"information.html",gin.H{
		"data" :users,
		"message":"查询成功",
	})


}

func Delete(c *gin.Context) {


	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	reuslt := logic.DeleteId(id)
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
	if reuslt == true {

		c.HTML(http.StatusOK,"data.html",gin.H{
			"data" :users,
			"message":"删除成功",
		})
		return
	}else {
		c.HTML(http.StatusUnprocessableEntity,"data.html",gin.H{
			"data" :users,
			"message":"删除失败",
		})
	}
}

func Edit(c *gin.Context) {
	c.HTML(http.StatusOK,"update.html",nil)
}

func Update(c *gin.Context)  {
	visitorName := c.PostForm("visitor_name")
	visitId := c.PostForm("visit_id")
	phone := c.PostForm("phone")

	if logic.SelectVisitorName(visitorName) == "" {
		c.HTML(http.StatusUnprocessableEntity,"updatedown.html", gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	if len(phone) != 11{
		c.HTML(http.StatusUnprocessableEntity,"updatedown.html", gin.H{
			"code":    422,
			"message": "电话号码有误",
		})
		return
	}


	if len(visitId) != 18{
		c.HTML(http.StatusUnprocessableEntity,"updatedown.html", gin.H{
			"code":    422,
			"message": "身份证有误",
		})
		return
	}


	reuslt := logic.Update(visitorName,phone,visitId)
	s := logic.SelectVisitor(visitorName,phone)
	var users []Visitors
	for i:=0;i<=len(s)-1;i++{
		id := s[i].ID
		visistorName := s[i].VisitorName
		sex := s[i].Sex
		phone := s[i].Phone
		visitId := s[i].VisitId
		createdAt := s[i].CreatedAt
		event := s[i].Event
		var user Visitors
		user = Visitors{ID: id,VisitorName: visistorName,Sex: sex,Phone: phone,VisitId: visitId,CreatedAt: createdAt,Event: event}

		users = append(users,user)
	}
	if reuslt == true {
		c.HTML(http.StatusOK,"information.html",gin.H{
			"data" :users,
			"message":"更新成功",
		})
		return
	}else {
		c.HTML(http.StatusUnprocessableEntity,"information.html",gin.H{
			"data" :users,
			"message":"更新失败",
		})
	}
}

func Find(c *gin.Context) {
	c.HTML(http.StatusOK,"select.html",nil)
}


//VisitorName string `gorm:"varchar(20);not null" form:"visitor_name" json:"visitor_name"`
//Sex string `gorm:"enum('男','女');not null" form:"sex" json:"sex"`
//Phone string `gorm:"varchar(20);not null" form:"phone" json:"phone"`
//VisitId string `gorm:"varchar(20);not null" form:"visit_id" json:"visit_id"`
//Event string `gorm:"varchar(100);not null" form:"event" json:"event"`

type Visitors struct {
	ID uint `json:"ID"`
	VisitorName string `json:"VisitorName"`
	Sex string `json:"Sex"`
	Phone string `json:"Phone"`
	VisitId string `json:"VisitId"`
	CreatedAt time.Time `json:"CreatedAt"`
	Event string `json:"Event"`
}


