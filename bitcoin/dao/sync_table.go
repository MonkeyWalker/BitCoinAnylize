package dao

import(
	"github.com/jinzhu/gorm"
	"../../globeconstant"
)


func SyncTable(userName string, pwd string,v interface{})  {
	db, err := gorm.Open(globeconstant.MYSQL, userName+":"+pwd+"@/douban?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if (db.HasTable(v)) {
		db.AutoMigrate(v) //当你在module 中添加了新的字段后,会同步到表中,
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB,charset=utf8").
			CreateTable(v)
	}
}