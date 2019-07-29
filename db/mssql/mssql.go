package mssql

import (
	"database/sql"
	"fmt"
	"log"
)

func Test() {

}

func GetSimpleTest() string {

	condb, err := sql.Open("mssql", "server=jackytest01.database.windows.net;user id=jacky;password=zaq1xsw2CDE#")
	if err != nil {
		fmt.Println("error")
	}

	defer condb.Close()

	stmt, err := condb.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Println(somenumber)
	fmt.Println(somechars)

	return "Jacky"
}

func SelectVersion() string {

	condb, err := sql.Open("mssql", "server=jackytest01.database.windows.net;user id=jacky;password=zaq1xsw2CDE#")
	if err != nil {
		fmt.Println("error")
	}

	defer condb.Close()

	var (
		sqlversion string
	)

	rows, err := condb.Query("select @@version")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&sqlversion)
	}

	return sqlversion
}
