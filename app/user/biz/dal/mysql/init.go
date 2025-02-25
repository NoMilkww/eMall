package mysql

import (
	"fmt"
	"github.com/feeeeling/eMall/app/user/biz/model"
	"github.com/feeeeling/eMall/app/user/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	//mysql_username := "root"
        //mysql_password := "root"
        //mysql_host := "127.0.0.1"
        //dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/payment?charset=utf8mb4&parseTime=True&loc=Local", mysql_username, mysql_password, mysql_host)
        dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	if err != nil {
		panic(err)
	}
}
