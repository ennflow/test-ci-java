package models

import(
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
)

var ZKConn *zk.Conn

func GetZkClient(){
	z,_,err:=zk.Connect([]string{"zookeeper-qinzhao.enncloudtest.tenxcloud.cc:40673"},10*time.Second)
	if err!=nil{
		fmt.Println("get zk failed",err)
		return
	}

	defer z.Close()

	//err=z.AddAuth("http",[]byte("qinzhao"))
	//if err!=nil{
	//	fmt.Println("zk error=",err)
	//	return
	//}

	if z!=nil{
		ZKConn=z
	}
	return

}

func TestZk(){

	//defer ZKConn.Close()

	//err:=ZKConn.AddAuth("http",[]byte("qinzhao"))
	//if err!=nil{
	//	fmt.Println("zk error=",err)
	//	return
	//}
	//
	fmt.Println(ZKConn.State().String())

	//children,stat,_,err:=ZKConn.ChildrenW("/")
	//if err!=nil{
	//	fmt.Println("zk childrenW failed",err)
	//	return
	//}
	//
	//fmt.Println("children",children,stat)



}
