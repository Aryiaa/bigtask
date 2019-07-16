package user

import (
	"app/model"
	"corecd-v2/src/app/library/sessions"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//创建用户
//func Create(c *gin.Context) (*model.BigdataUser,error) {
//
//}
func SearchUserByID(validuser ValidateUser)  (model.BigdataUser,error) {
	//fmt.Println("查看用户是否存在")
	var user model.BigdataUser
	ID:=validuser.ID
	result:=model.DB().First(&user,ID)
	if result.RecordNotFound()==true{
		//fmt.Println(" 不存在则添加")
		fmt.Println(validuser.Email)
		bigdatauser:=model.BigdataUser{ID:validuser.ID,Name:validuser.Name,AvatarUrl:validuser.AvatarUrl,Email:validuser.Email}
		model.DB().Create(&bigdatauser)
		//fmt.Println(" 添加数据库成功")
		return bigdatauser,nil
	}
	if err:=result.Error;err!=nil{
		return user,err
	}
	//c存在则更新登录时间
	model.DB().Model(&user).Update("UpdatedAt",time.Now())
	//fmt.Println(" 更新时间成功")
	return user,nil
}


//处理登录信息，如果存在更新登录时间，不存砸则保存到数据库
func HandleUserInfo(info string) error{
	var validuser ValidateUser
	fmt.Println("处理登录信息，如果存在更新登录时间，不存砸则保存到数据库")
	json.Unmarshal([]byte(info),&validuser)
	user,err:=SearchUserByID(validuser)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	//存在则更新数据
	fmt.Println("****")
	fmt.Println(user)
	return nil
}

//判断是否登录
func IsLogin(c *gin.Context)  bool {
	return sessions.Uid(c)!=0

}