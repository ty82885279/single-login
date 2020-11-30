package mysql

import (
	"fmt"
	"single-login/model"
	"single-login/pkg/jwt"
)

func GetToken(userID int64) (AToken, RToken string, err error) {

	AToken, RToken, err = jwt.GenToken(userID)
	if err != nil {
		fmt.Printf("genToken err :%v\n", err)
		return "", "", err
	}
	return
}
func FindUserID(u *model.User) (userID int64, err error) {
	var user = new(model.User)

	err = Db.Debug().Where("name = ? AND psw = ? ", u.Name, u.Psw).First(user).Error
	if err != nil {
		fmt.Printf("findUser Err :%v\n", err)
		return 0, err
	}
	userID = user.ID
	return
}
