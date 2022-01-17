package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goPOC/config"
	"log"
)

var DB *gorm.DB

/*
该方法虽然不是内置的初始化方法,但是会在main入口那边的第一行被调用,等于是起到了初始化的作用.

有个地方需要注意的是，Query、Exec操作用法有些差异：
	a.Exec(update、insert、delete等无结果集返回的操作)调用完后会自动释放连接；
	b.Query(返回sql.Rows)则不会释放连接，调用完后仍然占有连接，它将连接的所属权转移给了sql.Rows，
		所以需要手动调用close归还连接，即使不用Rows也得调用rows.Close()，否则可能导致后续使用出错，如下的用法是错误的
*/
func InitMysql() (*gorm.DB, error) {
	log.Println("@@@Init Mysql......")
	db, err := gorm.Open(config.GetConfig().Mysql.DriverName, buildDBString())
	if err == nil {
		DB = db
		//db.LogMode(true)
		db.SingularTable(true)
		log.Println("@@@Create Mysql tables......")
		CreateTables()
		log.Println("@@@Finish init Mysql")
		return db, err
	}
	return nil, err
}

func CreateTables() {
	//创建表
	//db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{}, &SmmsFile{})
	//创建索引
	//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
	CreateUser()

}

func CreateUser() {
	DB.AutoMigrate(&User{})
}

/*
用flag 加载program arguments
*/
func buildDBString() string {
	dbPassword := config.GetConfig().Mysql.Pswd
	dbUrl := config.GetConfig().Mysql.Host
	dbName := config.GetConfig().Mysql.DbName
	dbString := fmt.Sprintf("jsflzhong_1:%s@tcp(%s:3306)/%s?charset=utf8", dbPassword, dbUrl, dbName)
	log.Println("@@@[buildDBString]DB connection String:", dbString)
	return dbString
}
