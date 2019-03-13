//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:23
//==================================
package jorm

import (
	"fmt"
	"gitlab.iguiyu.com/park/struct/model"
	"log"
	"testing"
	"time"
)

/*
create table contact
(
  id              int auto_increment
    primary key,
  name            varchar(10) null comment '姓名',
  age             int         null comment '年龄',
  gender          char        null comment '性别<男、女>',
  phone_number    varchar(15) null comment '电话号码',
  qq_number       varchar(15) null comment 'QQ号码',
  wx_number       varchar(20) null comment '微信号码',
  home_address    varchar(50) null comment '家庭住址',
  company_address varchar(50) null comment '公司地址'
);
*/

type Contact struct {
	RealName    string `json:"real_name" jorm:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
}

func TestCallProcedure(t *testing.T) {
	//err := InitMySQL("root:Ming521.@tcp(jerry.igoogle.ink:3306)/db_test?charset=utf8")
	err := InitMySQL("developer:Iloveguiyu2018!@tcp(rm-uf6sl3y5zl5mku48jho.mysql.rds.aliyuncs.com:3306)/lock_test?charset=utf8")
	if err != nil {
		fmt.Println("err:", err)
	}
	session := MySQL().NewSession()

	//contact := new(Contact)
	//columns := []string{"name", "age", "phone_number", "home_address"}
	//
	//_, err = session.Where("name = ?", "付明明").Cols(columns...).Get(contact)
	//if err != nil {
	//	fmt.Println("err2:", err)
	//} else {
	//	fmt.Println("contact:", contact)
	//}
	//
	//sql, i := session.LastSQL()
	//log.Println("sql:", sql)
	//log.Println("i:", i)

	vip := new(model.SiteVip)
	has, err := session.Table("site_vip").Where("plate_number = ?", "沪HZ5690").
		And("start_time <= ?", time.Now().Format("2006-01-02 15:04:05")).Get(vip)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("has:", has)
	log.Println(vip)

	//poleLog := new(model.PoleLog)
	//_, err = session.Table("pole_log").Where("plate_number = ?", "沪HZ5690").Desc("id").Get(poleLog)
	//if err != nil {
	//	errorLog := new(model.ErrorLog)
	//	errorLog.Project = "测试项目"
	//	errorLog.Method = "TestCallProcedure"
	//	errorLog.Param = "sadfds"
	//	sql, value := session.LastSQL()
	//	errorLog.ErrorSql = fmt.Sprintf("%v : %v", sql, value)
	//	errorLog.ErrotMsg = err.Error()
	//	errorLog.CreateTime = time.Now()
	//	//写入数据库
	//	session.Table("error_log").InsertOne(errorLog)
	//}
	//log.Println("poleLog:", poleLog)

	//log := new(model.PoleLog)
	////_, err = MySQL().ID(172134).Get(log)
	//_, err = MySQL().ID(14).Get(log)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//
	//if log.EndTime.IsZero() {
	//	fmt.Println("endtime: null")
	//} else {
	//	fmt.Println("endtime:", log.EndTime)
	//}

	//result, err := CallProcedure("query_student", 1, 9).InParams("付明明").Query()
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//for _, v := range result {
	//	fmt.Println(v)
	//}

	//contact := new(Contact)
	//err = CallProcedure("query_student", 1, 9).InParams("付明明").Get(contact)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//fmt.Println("contact:", contact)
	//
	//contactList := make([]*Contact, 0)
	//err = CallProcedure("query_student", 1, 9).InParams("付明明").Find(&contactList)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//fmt.Println("contactList:", contactList)

}

func TestDbSearch(t *testing.T) {
	err := InitMySQL("developer:Iloveguiyu2018!@tcp(rm-uf6sl3y5zl5mku48jho.mysql.rds.aliyuncs.com:3306)/lock_test?charset=utf8")
	if err != nil {
		fmt.Println("err:", err)
	}

	poleLogList := make([]model.PoleLog, 0)
	err = MySQL().Where("plate_number = ?", "苏E6PS87").And("start_time < ?", "2019-02-14 13:31:44").And("pay_status = 0 or status = 'IN'").Find(&poleLogList)
	if len(poleLogList) > 0 {
		for _, v := range poleLogList {
			log.Println("VVV：", v)
		}
	}
}
