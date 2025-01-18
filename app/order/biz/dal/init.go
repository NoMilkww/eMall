package dal

import (
	"github.com/feeeeling/eMall/app/order/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
