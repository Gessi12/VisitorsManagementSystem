package controllers

import  (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/logic"
	"VisitorsManagementSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
func AddVisitor(c *gin.Context) {
	visitorName := c.PostForm("visitor_name")
	sex := c.PostForm("sex")
	phone := c.PostForm("phone")
	visitId := c.PostForm("visit_id")
	event := c.PostForm("event")

	if len(visitorName) == 0{
		c.HTML(http.StatusUnprocessableEntity,"registerdown.html", gin.H{
			"code":    422,
			"message" : "用户名不能为空",
		})
		return
	}

	if sex != "男" && sex != "女"{
			c.HTML(http.StatusUnprocessableEntity,"registerdown.html", gin.H{
				"code":    422,
				"message" : "性别只能为男女",
			})
			return
		}

	if len(phone) != 11{
		c.HTML(http.StatusUnprocessableEntity,"registerdown.html", gin.H{
			"code":    422,
			"message": "电话号码有误",
		})
		return
	}


	if len(visitId) != 18{
		c.HTML(http.StatusUnprocessableEntity,"registerdown.html", gin.H{
			"code":    422,
			"message": "身份证有误",
		})
		return
	}

	if len(event) == 0 {
		c.HTML(http.StatusUnprocessableEntity,"registerdown.html", gin.H{
			"code":    422,
			"message": "事由不能为空",
		})
		return
	}

	newVisitor := models.Visitor{
		VisitorName: visitorName,
		Sex: sex,
		Phone: phone,
		VisitId: visitId,
		Event: event,
	}
	dao.InitDB().Create(&newVisitor)

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
	//返回结果
	c.HTML(http.StatusOK,"data.html",gin.H{
		"code":"200",
		"data":users,
		"message":"登记成功",
	})
}

func GetAdd(c *gin.Context) {
	c.HTML(http.StatusOK,"input.html",nil)
}






