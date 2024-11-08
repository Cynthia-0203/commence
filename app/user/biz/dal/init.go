package dal

import (
	"github.com/Cynthia/commence/app/user/biz/dal/mysql"
	"github.com/Cynthia/commence/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
