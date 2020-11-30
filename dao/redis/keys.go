package redis

const (
	Prefix              = "slogin:"
	KeyUserStatus       = "user:login:"  //用户登陆状态  参数：userID
	KeyUserAtokenString = "user:atoken:" //用户的atoken 参数：userID
	KeyUserRtokenString = "user:rtoken:" //用户的rtoken 参数：userID
	KeyUserAlisaSrting  = "user:alisa:"  //用户的alisa  参数：userID
)

func GetRedisKey(key string) string {
	return Prefix + key
}
