package dal

import (
	"github.com/feeeeling/eMall/app/checkout/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
