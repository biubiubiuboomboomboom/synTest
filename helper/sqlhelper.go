package helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// send 

//连接
func NewSqlCon() (con *sql.DB) {
	Db, _ := sql.Open("mysql","root:123456@tcp(localhost:3306)/Test1")
	if Db==nil{
		println("连接失败")
		return
	}
	return Db
}



// 查询服务端现有最大数据

func Sql_SgetMax(conn *sql.DB) int {
	var res int
	conn.QueryRow("select max(ServiceID) from service").Scan(&res)
	return res
}

//逐条拿数据

func Sql_getOne(conn *sql.DB , i int ) int {
	var res int
	conn.QueryRow("select * from service where ServiceID > ? limit 1",i).Scan(&res)
	return res
}




//recevice 

//连接
func CNewSqlCon() (con *sql.DB) {
	Db, _ := sql.Open("mysql","root:123456@tcp(localhost:3306)/Test2?charset=utf8")
	if Db==nil{
		return
	}
	//println("连接成功")
	return Db
}



//将数据放到本地
//insert into receive(receive) values(1se
func Sql_insert(conn *sql.DB,i int) bool {
	if conn == nil{
		println(" 连接失败")
		return false
	}
	conn.Exec("insert into receive(receive) values(?)",i)
	return true
}

//查询客户端现有的最大数据

func Sql_CgetMax(conn *sql.DB) int {
	var res int
	conn.QueryRow("select max(receive) from receive").Scan(&res)
	return res
}



