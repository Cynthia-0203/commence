package dal

import (
	"github.com/Cynthia/commence/app/checkout/biz/dal/mysql"
	"github.com/Cynthia/commence/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
