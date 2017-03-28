package main

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

var c redis.Conn
var err error
var db *sql.DB

type CashierData struct {
	Id              int    `json:"id"`
	Mch_id          string `json:"mch_id"`
	Openid          string `json:"openid"`
	Cashiername     string `json:"cashiername"`    //收银员名称
	Sl_cashiername  string `json:"sl_cashiername"` //上属收银员
	Cashierpassword string //收银员登录密码
	Company         string `json:"company"`        //公司名称，唯一的
	State           string `json:"state"`          //online offline，收银员在线还是下载
	Store           string `json:"store"`          //收银员门店
	Ip              string `json:"ip"`             //213.90.5 2.45:8000,收银员使用的服务器后台地址
	Key             string `json:"key"`            //收银员支付秘钥
	Role            string `json:"role"`           //账号角色，sl,company,mchid,cahsier
	Adid            string `json:"adid"`           //广告id
	Createdate      int    `json:"createdate"`     //创建日期
	Pushurl         string `json:"pushurl"`        //结果推送地址,对接其他网站页面支付接口返回的地址
	Rights          int    `json:"rights"`         //0,退款权限，1为有退款权限
	Usertype        string `json:"usertype"`       //报表里面改收银员对应的客户类型
	Address         string `json:"address"`        //店铺地址
	Bd              string `json:"bd"`             //对接负责人
	Flag            string `json:"flag"`           //为店铺打标签，比方说汽车，美容，超时，餐饮等，方便广告设置
	Rate1           int    `json:"rate1"`          //费率 0-20 当天到账费率
	Rate2           int    `json:"rate2"`          //费率 0-20 隔天到账费率
	Superior        string `json:"superior"`       //上级账号
	Rate3           int    `json:"rate3"`          //信用卡代理商的返佣
	Cardopenid      string `json:"cardopenid"`     //信用卡代理商的openid
	Recipient_name  string `json:"recipient_name"` //小微商户收款人
	Micro_mch_id    string `json:"micro_mch_id"`   //小微收款商户号

}

func init() {
	c, err = redis.Dial("tcp", "127.0.0.1:6379") //定义redis
	checkouterr(err)
	db, err = sql.Open("mysql", "cdb_outerroot:wei2016KAI@tcp(56dd3ccb68390.gz.cdb.myqcloud.com:16862)/wechatpayv2?charset=utf8")
	checkouterr(err)
}
func HgetRedis() {
	v, err := redis.String(c.Do("GET", "name")) //返回的是ASCII码
	checkouterr(err)
	fmt.Println(v)
}
func HsetRedis() {
	v, err := c.Do("SET", "name", "yyq")
	checkouterr(err)
	fmt.Println(v)
}
func FindMysql() {
	//cashier := CashierData{}
	rows, _ := db.Query("select * from cashier where mch_id=1276316801")
	checkouterr(err)
	columns, _ := rows.Columns()
	fmt.Println(len(columns))
	values := make([]sql.RawBytes, len(columns)) //定义值切片
	scanArgs := make([]interface{}, len(values)) //定义切片
	for i := range values {
		scanArgs[i] = &values[i]
	}
	r := []string{}
	m := make(map[string]string)
	//n:=[]m
	for rows.Next() {
		//j++
		//cashier:=[]CashierData{}

		rows.Scan(scanArgs...)
		for k, v := range values {
			r = append(r, string(v))
			m[columns[k]] = string(v)
		}

	}
	fmt.Println(m)
}
func RedisClose() {
	c.Close()
}
func checkouterr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func Test() {
	type A struct {
		Name, Age string
	}
	c := []A{}
	//a := A{}
	//b := []string{"Jim", "15", "Tom", "16", "Mary", "17"}
	//c[0] = A{b[0], b[1]}
	v := reflect.ValueOf(&c).Elem().Type()
	fmt.Println(v)
	v1 := reflect.TypeOf(c).Elem()
	v2 := reflect.ValueOf(v1).Elem().Type()
	fmt.Println(v1, v2)
	//for i := 0; i < reflect.ValueOf(&c).Elem().NumField(); i++ {
	//	fmt.Println(i, v.Field(i).Name)
	//}
	//v2 := reflect.ValueOf(c).Elem()
	//v2.FieldByName().SetString(b[0])

}
