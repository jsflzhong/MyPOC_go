// model 是专门提供RESTful接口的路由请求处理handler
// 可以使用gorm等ORM工具来处理。为了代码简洁，这里使用原生sql处理。
// 用户数据处理
package model

import (
	"awesomeProject/src/com.jsflzhong/6_web/Gin/website_1/common"
	"log"
)

type User struct {
	//如果想让json.marshal()生效,则这里的字段名必须首大.
	Id        int    `json:"uid,string,omitempty" form:"uid"`
	User_name string `json:"nick" form:"nick"`
	Password  string `json:"Password" form:"Password"`
	//注意,这个字段是int,如果想把endpoint中的入参json直接通过调用c.bind(..)来转成本struct,则需要像这样处理一下非string类型的字段.
	//否则, bind时会报错:cannot unmarshal string into Go struct field User.Age of type int
	Age int `json:"Age,string,omitempty" form:"Age`
}

//抽取SQL, 便于维护和DBA评审
const getUserById string = "select Id,User_name,Age,Password from user_t where Id = ?"
const insertUser string = "INSERT INTO user_t(User_name, Age,Password) VALUES (?, ?, ?)"
const udpateUser string = "UPDATE user_t SET User_name=?, Age=?, Password=? where id=?"
const deleteUser string = "DELETE FROM user_t where id=?"

/*
针对entity的CRUD, 都写在了entity个文件里了?
相当于简单的service层中的CRUD了.
Dao层写在了model.go那边了.
*/
func (u *User) GetUser(uid int64) (usr User, err error) {
	//u.getFromRedis(uid)

	dbw := NewDb()
	defer dbw.Db.Close()

	user := User{}
	//单行查询
	err = dbw.Db.QueryRow(getUserById, uid).
		Scan(&user.Id, &user.User_name, &user.Age, &user.Password)
	if err != nil {
		log.Printf("@@@Query data error: %v\n", err)
		return user, err
	}
	log.Println("@@@[GetUser]:", user)

	return user, nil
}

func (u *User) InsertUser(user User) (int, error) {
	dbw := NewDb()
	defer dbw.Db.Close()

	stmt, err := dbw.Db.Prepare(insertUser)
	defer stmt.Close()
	if common.CheckError(err) {
		return -1, err
	}
	//执行插入操作
	result, err := stmt.Exec(user.User_name, user.Age, user.Password)
	if common.CheckError(err) {
		return -1, err
	}
	//返回插入的id
	id, err := result.LastInsertId()
	if common.CheckError(err) {
		log.Fatalln(err)
		return -1, err
	}
	log.Println("@@@[InsertUser]success,id:", id)
	//将id类型转换(int64转int)
	return int(id), nil
}

func (u *User) UpdateUser(user User) (int, error) {
	dbw := NewDb()
	defer dbw.Db.Close()

	stmt, err := dbw.Db.Prepare(udpateUser)
	defer stmt.Close()
	if common.CheckError(err) {
		return -1, err
	}
	//执行update操作
	result, err := stmt.Exec(user.User_name, user.Age, user.Password, user.Id)
	if common.CheckError(err) {
		return -1, err
	}
	//返回影响的行数
	lines, err := result.RowsAffected()
	if common.CheckError(err) {
		log.Fatalln(err)
		return -1, err
	}
	log.Println("@@@[UpdateUser]success,lines:", lines)
	//将id类型转换(int64转int)
	return int(lines), nil
}

func (u *User) DeleteUser(id int) (int, error) {
	dbw := NewDb()
	defer dbw.Db.Close()

	stmt, err := dbw.Db.Prepare(deleteUser)
	defer stmt.Close()
	if common.CheckError(err) {
		return -1, err
	}
	//执行delete操作
	result, err := stmt.Exec(id)
	if common.CheckError(err) {
		return -1, err
	}
	//返回影响的行数
	lines, err := result.RowsAffected()
	if common.CheckError(err) {
		log.Fatalln(err)
		return -1, err
	}
	log.Println("@@@[DeleteUser]success,lines:", lines)
	//将id类型转换(int64转int)
	return int(lines), nil
}

/*
连接redis获取数据
*/
func (u *User) getFromRedis(uid int64) {
	//cache
	client := NewRedis()

	val, errRedis := client.Get("user:" + string(uid)).Result()
	if errRedis != nil {
		log.Println("GetUser from cache error:", val, errRedis)
	}
}
