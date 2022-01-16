package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

/*
安装:
go get -u github.com/jinzhu/gorm

数据模型定义#
表名，列名如何对应结构体
在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。

即，如果有一个user表，那么如果你定义的结构体名为：User，gorm会默认表名为users而不是user。

例如有如下表结构定义：

CREATE TABLE `areas` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `area_id` int(11) NOT NULL COMMENT '区县id',
  `area_name` varchar(45) NOT NULL COMMENT '区县名',
  `city_id` int(11) NOT NULL COMMENT '城市id',
  `city_name` varchar(45) NOT NULL COMMENT '城市名称',
  `province_id` int(11) NOT NULL COMMENT '省份id',
  `province_name` varchar(45) NOT NULL COMMENT '省份名称',
  `area_status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '该条区域信息是否可用 ： 1:可用  2：不可用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='区域表'
那么对应的结构体定义如下：

type Area struct {
	Id int
	AreaId int
	AreaName string
	CityId int
	CityName string
	ProvinceId int
	ProvinceName string
	AreaStatus int
	CreatedAt time.Time
	UpdatedAt time.Time
}
如何全局禁用表名复数呢？

可以在创建数据库连接的时候设置如下参数：

// 全局禁用表名复数
db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
这样的话，表名默认即为结构体的首字母小写形式。
*/

func main() {
	defer db.Close()

	//insert
	db.Create(&User{Name: "zhangSan", Age: 18,Sex:1,Phone:"13333333333"})

	//query
	var user User
	db.First(&user, 1)// 查询id为1的user
	fmt.Println("@@@user1:",user)
	db.First(&user, "name=?","zhangSan") //查询name为zhangsan的user.
	fmt.Println("@@@user2:",user)

	//update
	db.Model(&user).Update("phone", "13111111111")
	db.First(&user, 1)
	fmt.Println("@@@user1:",user)

	//delete
	//db.Delete(&user)
}

type User struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

var db *gorm.DB

func init() {
	var err error
	dbPassword := flag.String("db_password", "nil", "The password of db connection")
	dbUrl := flag.String("db_url", "nil", "The url of db connection")
	flag.Parse()

	dbString := fmt.Sprintf("root:%s@tcp(%s:3306)/test2", *dbPassword, *dbUrl)
	log.Println("@@@DBUrl=",dbString)
	//":="只在当前方法内有效, 无法给全局变量赋值.需要改成"="
	//否则下面引用全局变量时会报错: invalid memory address or nil pointer dereference
	db, err = gorm.Open("mysql", dbString)
	if err != nil {
		panic(err)
	}

	//设置全局表名禁用复数(否则表明是结构体名的复数形式.例结构体是:user,则表名是users)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// 自动迁移模式.(自动根据struct来在DB中创建表)
	// 自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据
	//db.AutoMigrate(&User{})
	//defer db.Close()
}

//插入数据
func (user *User) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	//db.Table("user").Create(user)
	db.Create(user)
}
