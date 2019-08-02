package jackymysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type Employee struct {
	gorm.Model
	Id     int    `gorm:"Type:int;PRIMARY_KEY;AUTO_INCREMENT"`
	Name   string `gorm:"Type:varchar(100);"`
	Depart string `gorm:"Type:varchar(100)"`
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
