package mssql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func SimpleTest() string {

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
	ctx := context.Background()
	err = condb.QueryRowContext(ctx, "select @@version").Scan(&sqlversion)
	if err != nil {
		fmt.Println(err)
	}

	return sqlversion
}
