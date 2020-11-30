package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var (
	ErrUserHasLogin = errors.New("用户已经登陆")
)

// 获取用户登陆状态
func GetUserStatus(userID int64) (err error) {

	UserStatusKey := GetRedisKey(KeyUserStatus + strconv.Itoa(int(userID)))

	result, err := rdb.Get(UserStatusKey).Result()

	if err != nil {
		if err == redis.Nil { //用户第一次登陆，存2个token
			fmt.Println("没有值")
			err = rdb.Set(UserStatusKey, "1", time.Duration(viper.GetInt("auth.jwt_rtoken_expire"))*time.Minute).Err()
			if err != nil {
				fmt.Printf("rdb.Set err:%v\n", err)

				return
			}
			err = nil
		}
		return
	}
	fmt.Printf("get result:%s\n", result)
	if result == "1" { //重复登陆，需要给之前的登陆的设备发消息
		fmt.Println("重复登陆，需要给之前的登陆的设备发消息")
		_ = rdb.Set(UserStatusKey, "1", time.Duration(viper.GetInt("auth.jwt_rtoken_expire"))*time.Minute).Err()
		return ErrUserHasLogin
	}
	return
}

// 保存token
func SaveToken(userID int64, AToken, RToken string) (err error) {

	UserAtokenKey := GetRedisKey(KeyUserAtokenString + strconv.Itoa(int(userID)))
	err = rdb.Set(UserAtokenKey, AToken, time.Duration(viper.GetInt("auth.jwt_atoken_expire"))*time.Minute).Err()
	if err != nil {
		fmt.Printf("set UserAtokenKey err:%v\n", err)
		return
	}
	UserRtokenKey := GetRedisKey(KeyUserRtokenString + strconv.Itoa(int(userID)))
	err = rdb.Set(UserRtokenKey, RToken, time.Duration(viper.GetInt("auth.jwt_rtoken_expire"))*time.Minute).Err()
	if err != nil {
		fmt.Printf("set UserRtokenKey err:%v\n", err)
		return
	}
	return
}

// 保存Alisa
func SaveAlisa(userID string, alisa string) bool {

	UserAlisaKey := GetRedisKey(KeyUserAlisaSrting + userID)
	err := rdb.Set(UserAlisaKey, alisa, time.Duration(viper.GetInt("auth.jwt_rtoken_expire"))*time.Minute).Err()

	if err != nil {
		fmt.Printf("save alisa err :%v\n", err)
		return false
	}
	return true
}
func UpdateAlisaTime(userID string) bool {
	UserAlisaKey := GetRedisKey(KeyUserAlisaSrting + userID)
	err := rdb.Expire(UserAlisaKey, time.Duration(viper.GetInt("auth.jwt_rtoken_expire"))*time.Minute).Err()
	if err != nil {
		fmt.Printf("save alisa err :%v\n", err)
		return false
	}
	return true
}
func UpdateLoginTime(userID string) bool {
	UserStatusKey := GetRedisKey(KeyUserStatus + userID)
	err := rdb.Expire(UserStatusKey, time.Duration(viper.GetInt("auth.jwt_rtoken_expire"))*time.Minute).Err()
	if err != nil {
		fmt.Printf("save login status err :%v\n", err)
		return false
	}
	return true
}
func GetAlisa(userID string) (alisa string, err error) {

	UserAlisaKey := GetRedisKey(KeyUserAlisaSrting + userID)
	alisa, err = rdb.Get(UserAlisaKey).Result()
	if err != nil {
		fmt.Printf("save alisa err :%v\n", err)
		return
	}
	return
}

func GetAToken(userID string) (Atoken string, err error) {

	AtokenKey := GetRedisKey(KeyUserAtokenString + userID)
	Atoken, err = rdb.Get(AtokenKey).Result()

	fmt.Printf("AtokenKey---%v\n", AtokenKey)
	fmt.Printf("UID---%v\n", userID)
	if err != nil {
		fmt.Printf("rdb.Get(AtokenKey) err :%v\n", err)
		return
	}
	return
}

func GetRToken(userID string) (Rtoken string, err error) {

	RtokenKey := GetRedisKey(KeyUserRtokenString + userID)
	Rtoken, err = rdb.Get(RtokenKey).Result()

	if err != nil {
		fmt.Printf("rdb.Get(RtokenKey) err :%v\n", err)
		return
	}
	return
}

// 用户退出登陆
func UserLogOut(userID string) (err error) {

	UserStatusKey := GetRedisKey(KeyUserStatus + userID)

	err = rdb.Del(UserStatusKey).Err() //删掉登录状态的key
	if err != nil {
		fmt.Printf("UserStatusKey.Del err:%v\n", err)
		return
	}
	UserAtokenKey := GetRedisKey(KeyUserAtokenString + userID)
	err = rdb.Del(UserAtokenKey).Err()
	if err != nil {
		fmt.Printf("UserAtokenKe.Del err:%v\n", err)
		return
	}

	UserRtokenKey := GetRedisKey(KeyUserRtokenString + userID)
	err = rdb.Del(UserRtokenKey).Err()
	if err != nil {
		fmt.Printf("UserRtokenKey.Del err:%v\n", err)
		return
	}

	UserAlisaKey := GetRedisKey(KeyUserAlisaSrting + userID)
	err = rdb.Del(UserAlisaKey).Err()
	if err != nil {
		fmt.Printf("UserAlisaKey.Del err:%v\n", err)
		return
	}

	return
}
