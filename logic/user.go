package logic

import (
	"fmt"
	"gorm.io/gorm"
	"single-login/dao/mysql"
	"single-login/dao/redis"
	"single-login/model"
)

func FindUser(u *model.User) (userID int64, AToken, RToken string, err error) {

	//先通过mysql查找userID
	userID, err = mysql.FindUserID(u)
	if err == gorm.ErrRecordNotFound {
		return
	}
	//根据userID去redis查看用户的登陆状态
	err = redis.GetUserStatus(userID)
	if err != nil {
		if err == redis.ErrUserHasLogin {
			fmt.Println("用户已经登录")

			//获取别名给发送消息
			SendMsgByAlisa(userID)
			err = nil
		}
		if err != nil {
			return 0, "", "", err
		}
	}

	// 生成token
	AToken, RToken, err = mysql.GetToken(userID)
	if err != nil {
		fmt.Printf("get token err:%v\n", err)
		return
	}
	//atoken,rtoken保存至redis
	err = redis.SaveToken(userID, AToken, RToken)
	if err != nil {
		return
	}
	//

	return
}
