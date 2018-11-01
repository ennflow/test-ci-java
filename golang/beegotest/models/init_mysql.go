package models

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"time"
)

var SqlDB *sql.DB

type User struct {
	Id   int64
	Name string
}

func init() {
	//testmysql-qinzhao.enncloudtest.tenxcloud.cc:51316
	sqlDBObject, err := sql.Open("mysql", "root:qinzhao@tcp(testmysql-qinzhao.enncloudtest.tenxcloud.cc:51316)/xinzhiyun?charset=utf8")
	//sqlDBObject, err := sql.Open("mysql", "root:@tcp(testmysql-qinzhao.enncloudtest.tenxcloud.cc:51316)/xinzhiyun?charset=utf8")

	if err != nil {
		fmt.Println("connect mysql failed", err)
		return
	}
	err = sqlDBObject.Ping()
	if err != nil {
		fmt.Println("sqlDB ping failed", err)
		return
	}

	SqlDB = sqlDBObject
}

func InitMysql() error {
	var user User
	user.Id = time.Now().Unix()
	user.Name = "xinzhiyun"
	stmt, err := SqlDB.Prepare("insert into user(id,name)value(?,?)")
	if err != nil {
		fmt.Println("sqlDB Prapare failed", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name)
	if err != nil {
		fmt.Println("stmt execu falied", err)
		return err
	}
	return nil

}