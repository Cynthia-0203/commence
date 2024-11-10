package dal

import (
	"github.com/Cynthia/commence/app/payment/biz/dal/mysql"
	// "github.com/Cynthia/commence/app/payment/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
