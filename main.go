package main

import (

	//"demo/servicecall"
	//"demo/restapi/gin"
	//"demo/db/mssql"
	//"demo/db/mysql"
	"fmt"
)

func main() {

	//test for consuming a restful api
	//servicecall.GetResponse()

	//test for operating on a mssql
	//mssql.DisplayVersion()

	//test for operating on a mssql with gorm

	//test for operating on a mysql
	/*var c jackymysql.Consumer
	c.Name = "Vincent"
	c.Age = 36
	c.Location = "HangZhou"
	jackymysql.InsertConsumer(c)*/
	//jackymysql.GetConsumers()

	//test for operating on a myslq with gorm
	//jackymysql.InsertUser("Jacky", "FO")
	//jackymysql.UpdateUser("Robert", "BTS")
	//jackymysql.FindUsers()

	//test for hosting a restful api service (Gin)
	//restapi.StartTest01Services()

	//test for get settings from config file

	//demo

	fmt.Println("done in main")

}
