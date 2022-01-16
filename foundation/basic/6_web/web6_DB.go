package web

import (
	"database/sql"
	"fmt"

	// 注意这个包前边要有个下划线, 表示为了只初始化,而不使用.
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/*
想要连接mysql的准备工作:
CMD:
	1.下载安装
		想要成功引到上面的:_ "github.com/go-sql-driver/mysql", 想要做几个前置工作:
	　　执行下面两个命令：
	　　下载：go get github.com/Go-SQL-Driver/MySQL
			这一步后,在GO_PATH目录下会自动新建一个目录:pkg. 下载的东西都在里面.
			执行完这一步后,上面的import已经没有红线报错了.
	 　 安装：go install github.com/Go-SQL-Driver/MySQL
	2.导入包. 如上.
	　　import (
			　　"database/sql"
			　　_"github.com/Go-SQL-Driver/MySQL"
	　　)

fixme: 注意:!
有个地方需要注意的是，Query、Exec操作用法有些差异：
	a.Exec(update、insert、delete等无结果集返回的操作)调用完后会自动释放连接；
	b.Query(返回sql.Rows)则不会释放连接，调用完后仍然占有连接，它将连接的所属权转移给了sql.Rows，
		所以需要手动调用close归还连接，即使不用Rows也得调用rows.Close()，否则可能导致后续使用出错，如下的用法是错误的
 */

//SQL提取到外围,便于维护和DBA评审.
const queryUser = "select id, user_name from user_t where id = ?"
const insertUser = "INSERT user_t SET user_name=?, age=?,password=?"


func SimpleMysql() {
	//获取DB连接
	db := getDbConnection()
	//需要定义在外层这里.因为需要给下面使用. 否则在内层方法结束后就已经close了.
	defer db.Close()

	Query(db)

	//Insert(db)

}

/*
insert.

删除
	删除和这次的增加语法一样，只是把其中的INSERT语句改为DELETE语句
修改
	修改和这次的增加语法一样，只是把其中的INSERT语句改为UPDATE语句

 */
func Insert(db *sql.DB) {
	stmt, err := db.Prepare(insertUser)
	defer stmt.Close()
	CheckErr(err)
	res, err := stmt.Exec("李四", 12, "123321")
	CheckErr(err)
	id, err := res.LastInsertId()
	fmt.Println("@@@insert successful, the last id is :", id)
}

/*
带事务的Insert
 */
func InsertWithTransaction(db *sql.DB) {
	//事务
	tx, errBegin := db.Begin()
	CheckErr(errBegin)

	//用事务返回的tx执行. 而不是db.
	stmt, err1 := tx.Prepare(insertUser)
	defer stmt.Close()
	CheckErr(err1)
	_, err2 := stmt.Exec("李四", 12, "123321")
	CheckErr(err2)

	//err3 := tx.Commit()
	err3 := tx.Rollback()
	CheckErr(err3)
}

/*
获取DB连接

有个地方需要注意的是，Query、Exec操作用法有些差异：
	a.Exec(update、insert、delete等无结果集返回的操作)调用完后会自动释放连接；
	b.Query(返回sql.Rows)则不会释放连接，调用完后仍然占有连接，它将连接的所属权转移给了sql.Rows，
		所以需要手动调用close归还连接，即使不用Rows也得调用rows.Close()，否则可能导致后续使用出错，如下的用法是错误的
 */
func getDbConnection() *sql.DB {
	// db 是一个 sql.DB 类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	// 连接池的选项可以在 sql.DB 的方法中设置，这里为了简单省略了
	// 同一个数据库只需要调用一次Open即可
	db, errOpen := sql.Open("mysql", "root:skrlab7365@tcp(47.104.87.68:3306)/test1")
	CheckErr(errOpen)
	return db
}

/*
Query

fixme: 没关闭row? -- row.Scan的源码里关闭了. 调用row.scan不用再关闭row, 也没有API可以关闭.
 */
func Query(db *sql.DB) {
	rows, errQuery := db.Query(queryUser, 1)
	CheckErr(errQuery)
	defer rows.Close()
	var (
		id   int
		name string
	)
	// 必须要把 rows 里的内容读完，或者显式调用 Close() 方法，
	// 否则在 defer 的 rows.Close() 执行之前，连接永远不会释放
	for rows.Next() {
		err := rows.Scan(&id, &name)
		CheckErr(err)
		log.Println("@@@Query: 已取到数据:", id, name)
	}
	err := rows.Err()
	CheckErr(err)
}

func CheckErr(err error){
	if err != nil {
		panic(err)
	}
}