package dal

import (
	"github.com/feeeeling/eMall/app/product/biz/dal/mysql"
	"github.com/feeeeling/eMall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
