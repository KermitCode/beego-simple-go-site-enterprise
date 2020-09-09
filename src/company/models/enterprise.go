package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Getdata() map[int]map[string]string {

	db, _ := sql.Open("mysql", "user:password@tcp(114.215.10.24:3306)/04007CN?charset=utf8")
	rows2, _ := db.Query("SELECT id,category,title,leader,leaderbaike,area,description,sortnum_china," +
		"sortyear,updateTime,website FROM qd_enterprise where objectcp=1 order by score desc")
	//返回所有列
	cols, _ := rows2.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	db.Close()

	return result

}
