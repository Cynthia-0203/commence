package dal

import (
	"github.com/Cynthia/commence/app/cart/biz/dal/mysql"
	"github.com/Cynthia/commence/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
