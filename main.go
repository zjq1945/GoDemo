package main

import (

	//"demo/servicecall"
	"demo/restapi/gin"
	//"demo/db/mssql"
	//"demo/db/mysql"
	//"encoding/json"
	//"demo/utility"
	//"github.com/patrickmn/go-cache"

	"fmt"
)

func main() {

	////test for consuming a restful api
	//servicecall.GetResponse()

	////test for operating on a mssql
	//mssql.DisplayVersion()

	////test for operating on a mssql with gorm

	////test for operating on a mysql
	// var c jackymysql.Consumer
	// c.Name = "Vincent"
	// c.Age = 36
	// c.Location = "HangZhou"
	// jackymysql.InsertConsumer(c)
	//fmt.Println(jackymysql.GetConsumers())
	// r, result := jackymysql.GetConsumer("Mary")
	// if result {
	// 	fmt.Println(r)
	// } else {
	// 	fmt.Println("not found")
	// }
	////delete a consumer
	// jackymysql.DeleteConsumer("Vincent")

	////test for operating on a myslq with gorm
	//jackymysql.InsertUser("Jacky", "FO")
	//jackymysql.UpdateUser("Robert", "BTS")
	//concumsers := jackymysql.GetConsumers()

	/*jsonData, err := json.Marshal(concumsers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))*/

	////test for hosting a restful api service (Gin)
	//restapi.StartTest01Services()

	////test for get settings from config file
	//fmt.Println(utility.GetMySqlConnectionString())

	////demo
	restapi.StartDemoServices()

	fmt.Println("done in main")

}
