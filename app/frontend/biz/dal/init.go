package dal

import (
	"github.com/feeeeling/eMall/app/frontend/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
