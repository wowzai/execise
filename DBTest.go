package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	//连接数据库
	db, err := sql.Open("mysql", "user:pwd@/wowzai?charset=utf8")
	checkErr(err)

	//插入数据
	/*
		以下三种插入写法都可以
	*/
	/*
		stmt,err := db.Prepare("insert userinfo set uname=?,uage=?,ucity=?,usex=?")
		checkErr(err)

		res,err := stmt.Exec("wowzai",24,"Hangzhou","M")
		checkErr(err)


		stmt,err := db.Prepare(`insert userinfo set uname='wowzai',uage=24,
							ucity='Hangzhou',usex='M'`)
	*/

	stmt, err := db.Prepare(`insert userinfo values(?,?,?,?)`)
	checkErr(err)

	res, err := stmt.Exec("wowzai", 24, "Hangzhou", "M")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	res, err = stmt.Exec("ding", 25, "Hangzhou", "M")
	checkErr(err)

	id, err = res.LastInsertId()
	checkErr(err)

	fmt.Println("id =", id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set uage=? where uname=?")
	checkErr(err)

	res, err = stmt.Exec(25, "wowzai")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect =", affect)

	//查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uname, ucity, usex string
		var uage int
		err = rows.Scan(&uname, &uage, &ucity, &usex)
		checkErr(err)
		fmt.Printf("uname=%s,uage=%d,ucity=%s,usex=%s\n", uname, uage, ucity, usex)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uname=?")
	checkErr(err)

	res, err = stmt.Exec("wowzai")
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("affect =", affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
