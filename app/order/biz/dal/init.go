package dal

import (
	"github.com/Cynthia/commence/app/order/biz/dal/mysql"
	"github.com/Cynthia/commence/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
