package jackymysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InsertConsumer(consumer Consumer) {

	db, err := sql.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("error to open:", err)
	}
	defer db.Close()

	stmtIns, err := db.Prepare("insert into `dbo`.`consumers`(`Name`,`Age`,`Location`) values(?,?,?)")
	if err != nil {
		fmt.Println("error to prepare:", err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(consumer.Name, consumer.Age, consumer.Location)
	if err != nil {
		fmt.Println("error to insert:", err)
	}
	fmt.Println("done to insert")
}

func GetConsumers() []Consumer {
	db, err := sql.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("error to open:", err)
	}
	defer db.Close()

	rows, err := db.Query("select * from consumers")
	if err != nil {
		fmt.Println("error to select:", err)
	}
	defer rows.Close()
	var rtn []Consumer
	for rows.Next() {
		var c Consumer
		err = rows.Scan(&c.ID, &c.Name, &c.Age, &c.Location)
		if err != nil {
			fmt.Println("error to scan:", err)
		}
		fmt.Println(c.ID, c.Name, c.Age, c.Location)
		rtn = append(rtn, c)

	}
	fmt.Println("done in search")
	return rtn
}
