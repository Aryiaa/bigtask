package controller

import (
	"corecd-v2/src/app/service/project"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectController struct {

}

func (t *ProjectController) InitRouters(c *gin.Engine)  {
	project:=c.Group("/project")
	{
		project.GET("/index",t.Index)
		project.GET("/create",t.Create)
		project.POST("/create",t.DoCreate)
	}

}

func (t *ProjectController) Index(c *gin.Context)  {
	pages,err:=project.PageList(c)

	fmt.Println("*******查询所有项目")
	fmt.Println(pages)
	c.HTML(http.StatusOK, "project/index.html", gin.H{
		"title": "项目首页",
		"appErr":err,
		"pages":pages,
	})

}

func (t *ProjectController) Create(c *gin.Context)  {

}

func (t *ProjectController) DoCreate(c *gin.Context)  {

}