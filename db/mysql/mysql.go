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

func DeleteConsumer(consumerName string) {

	db, err := sql.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("error to open:", err)
	}
	defer db.Close()

	stmtDel, err := db.Prepare("delete from consumers where Name = ?")
	if err != nil {
		fmt.Println("error to prepare:", err)
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(consumerName)
	if err != nil {
		fmt.Println("error to delete:", err)
	}
	fmt.Println("done to delete")
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

func GetConsumer(consumerName string) (Consumer, bool) {
	db, err := sql.Open("mysql", "jacky@mysql-test02:zaq1xsw2CDE#@tcp(mysql-test02.mysql.database.azure.com:3306)/dbo?parseTime=true")
	if err != nil {
		fmt.Println("error to open:", err)
	}
	defer db.Close()
	var rtn Consumer
	err = db.QueryRow("select * from consumers where name=?", consumerName).Scan(&rtn.ID, &rtn.Name, &rtn.Age, &rtn.Location)
	if err != nil {
		fmt.Println("error to scan:", err)
		return rtn, false
	} else {
		return rtn, true
	}
}
