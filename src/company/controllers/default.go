package controllers

import (
	"company/models"
	//"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Yesshow"] = 1

	//row := make(map[string]string)
	//fmt.Println(row)
	c.Data["enterprise"] = models.Getdata()
	c.Data["snum"] = 1
	//fmt.Println(c)
	c.TplName = "index.tpl"

	// id := 0
	// title := ""
	// category := ""
	// leader := ""
	// leaderbaike := ""
	// area := ""
	// description := ""
	// sortnum_china := 0
	// sortyear := ""
	// updateTime := 0
	// var enterprise []type = make([]type, 15)
	// for rows.Next() {
	// 	rows.Scan(&id, &category, &title, &leader, &leaderbaike, &area, &description, &sortnum_china, &sortyear, &updateTime)
	// 	fmt.Println(id, title)
	// 	enterprise
	// }

}

//func getEnterpriseList() [] {

//data := getEnterpriseList()
//检查是否连接成功数据库

//return data
// stmt, err := db.Prepare("INSERT INTO user_info SET username=?,departname=?,create_time=?")
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// res, err := stmt.Exec("nulige", "商务部", "2019-1-28")
// id, err := res.LastInsertId()
// if err != nil {
// 	panic(err)
// }

// fmt.Println(id)
//}
