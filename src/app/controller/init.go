package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	InitRouters(c *gin.Engine)
}
var controllerList=[] Controller{
	new(ProjectController),
	new(LoginController),
}
func Init(c *gin.Engine)  {
 for _,v:=range controllerList{
 	v.InitRouters(c)

	}
}

