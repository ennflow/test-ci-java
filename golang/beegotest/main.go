package main

import (
	_ "golang/beegotest/routers"
	"github.com/astaxie/beego"
	//"golang/beegotest/models"
	//"fmt"
	//"os"
)

func main() {
	//err:=models.InitMysql()
	//if err!=nil{
	//	fmt.Fprintln(os.Stdout,"faile=",err)
	//	return
	//}
	//
	//models.InitRedisClient()
	//result,err:=models.SetValue()
	//if err!=nil{
	//	fmt.Fprintln(os.Stdout,"redis set value failed=",err)
	//	return
	//}
	//fmt.Println("result=",result)
	//value,err:=models.GetValue()
	//if err!=nil{
	//	fmt.Fprintln(os.Stdout,"redis get value failed=",err)
	//	return
	//}
	//fmt.Println("value=",value)
	//models.GetZkClient()
	//
	//models.TestZk()

	beego.Run()
}

