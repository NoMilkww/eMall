package dal

import (
	"github.com/feeeeling/eMall/app/user/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
