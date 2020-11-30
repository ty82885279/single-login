package logic

import (
	"fmt"
	"github.com/spf13/viper"
	jpushclient "github.com/ylywyn/jpush-api-go-client"
	"single-login/dao/redis"
	"strconv"
)

func SendMsgByAlisa(userID int64) bool {

	//获取alisa
	alisa, err := redis.GetAlisa(strconv.Itoa(int(userID)))
	if err != nil {
		fmt.Printf("get Alisa err:%v\n", err)
		return false
	}

	//通过alisa发消息
	SendMsg(alisa)

	return true
}

func SendMsg(alisa string) {
	var pf jpushclient.Platform
	pf.Add(jpushclient.IOS)

	var ad jpushclient.Audience
	ad.SetAlias([]string{alisa})
	var msg jpushclient.Message
	msg.Content = `重复登陆`
	var op jpushclient.Option
	op.ApnsProduction = true

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetMessage(&msg)
	payload.SetAudience(&ad)
	payload.SetOptions(&op)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	push := jpushclient.NewPushClient(viper.GetString("jpush.master"), viper.GetString("jpush.appkey"))
	str, err := push.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	} else {
		fmt.Printf("ok:%s\n", str)
	}
}
