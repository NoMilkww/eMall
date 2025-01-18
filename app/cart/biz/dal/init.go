package dal

import (
	"github.com/feeeeling/eMall/app/cart/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
