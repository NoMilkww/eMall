package dal

import (
	"github.com/feeeeling/eMall/app/auth/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
