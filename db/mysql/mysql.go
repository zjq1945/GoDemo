package jackymysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type Employee struct {
	ID     int
	Name   string
	Depart string
}

func InsertUser(name string, depart string) {
	db, err := gorm.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo")
	if err != nil {
		fmt.Println("err:", err.Error)
	}
	fmt.Println("opened with NO error")
	defer db.Close()

	db.AutoMigrate(&Employee{})

	employee := Employee{Name: name, Depart: depart}

	db.Create(&employee)
}

func UpdateUser(name string, departToUpdate string) {
	db, err := gorm.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("err:", err.Error)
	}
	fmt.Println("opened with NO error")
	defer db.Close()

	db.AutoMigrate(&Employee{})

	var employee Employee
	db.Model(&employee).Where("Name=?", name).Update("Depart", departToUpdate)

}

func FindUser(name string) {
	db, err := gorm.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("err:", err.Error)
	}
	fmt.Println("opened with NO error")
	defer db.Close()

	e := &Employee{}
	db.Where("Name=?", name).Find(&e)
	fmt.Println(e)

}

func FindUsers() {
	db, err := gorm.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("err:", err.Error)
	}
	fmt.Println("opened with NO error")
	defer db.Close()

	var employees []Employee

	db.Find(&employees)
	fmt.Println(&employees)

}
