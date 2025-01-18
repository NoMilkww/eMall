package dal

import (
	"github.com/feeeeling/eMall/app/payment/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
