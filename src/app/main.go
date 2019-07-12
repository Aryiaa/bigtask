package main

import (
	"app/conf"
	"app/controller"
	"app/model"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"runtime"
	"time"
)

func main()  {

	runtime.GOMAXPROCS(runtime.NumCPU())
	engine:=gin.Default()
	initApp(engine)
	runApp(engine)

}
func initApp(engine *gin.Engine) {
	//初始化配置文件
	conf.Initconf()
	initSetting()
	InitVIews(engine)
	model.Init()
	controller.Init(engine)


}
func initSetting()  {
	req.SetTimeout(10*time.Second)

}

func InitVIews(c *gin.Engine)  {
	c.LoadHTMLGlob("view/**/*")
	c.Static("/public","./public")

}

func runApp(engine *gin.Engine) {
	engine.Run(":6060")

}