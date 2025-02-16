package dal

import (
	"github.com/feeeeling/eMall/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
