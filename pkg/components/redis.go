package components

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var Cache cache.Cache

func RedisInit() {
	redisconn := beego.AppConfig.String("redis_conn")
	redisport := beego.AppConfig.String("redis_port")
	redispwd := beego.AppConfig.String("redis_pwd")
	beego.Info(redisconn)
	redisConf := fmt.Sprintf(`{"key":"%s","conn":"%s:%s","dbNum":"%d","password":"%s"}`,
		"admin",
		redisconn,
		redisport,
		0,
		redispwd,
	)
	var err error
	Cache, err = cache.NewCache("redis", redisConf)
	if err != nil {
		beego.Info(redisConf)
		beego.Error("Redis connection fail:" + err.Error())
	}

}
