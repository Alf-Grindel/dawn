package dal

import (
	"fmt"
	"github.com/alf-grindel/dawn/conf"
	"github.com/alf-grindel/dawn/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	sql := conf.MySQL
	dsn := fmt.Sprintf(constants.DsnTemplate, sql.Username, sql.Password, sql.Addr, sql.Database)

	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
		})
	if err != nil {
		log.Fatalln("dal.Init: mysql connect failed, ", err)
	}
}

func Open() *gorm.DB {
	return db
}
