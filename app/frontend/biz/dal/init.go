package dal

import (
	"github.com/Cynthia/commence/app/frontend/biz/dal/mysql"
	"github.com/Cynthia/commence/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
