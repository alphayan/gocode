package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var c redis.Conn //定义redis连接
var err error
var db *sql.DB //定义mysql连接

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

func initRedis() {
	c, err = redis.Dial("tcp", "192.168.0.105:6379") //定义redis
	checkouterr(err)
}
func HgetRedis() {
	v, err := redis.String(c.Do("GET", "name")) //返回的是ASCII码,需要用string转型
	checkouterr(err)
	fmt.Println(v)
}
func HsetRedis() {
	v, err := c.Do("SET", "name", "yyq")
	checkouterr(err)
	fmt.Println(v)
}
func Hmset(a interface{}) {
	var b []interface{}
	args, _ := a.([]map[string]string) //类型断言
	for i, j := range args {
		//fmt.Println(i, j)
		argss := []string{"mp" + strconv.Itoa(i)}
		for k, v := range j {
			argss = append(argss, k, v) //将每个map转换成[]string
		}
		//fmt.Println(argss)
		b = append(b, argss) //将string切片转换成切片的切片
	}

	for bi, _ := range b {
		d := make([]interface{}, 0)
		for _, bv := range b[bi].([]string) { //因为HMSET的参数为...interface{}型所以需要转换[]string到[]interface{}
			d = append(d, bv)
		}

		va, err := c.Do("HMSET", d...)
		checkouterr(err)
		fmt.Println(va)
	}
}
func Hgetall() { //取出map
	v, err := redis.Strings(c.Do("KEYS", "*")) //获取所有的key
	checkouterr(err)
	fmt.Println(v)
	for i := 0; i < 14; i++ {
		v, err := redis.Strings(c.Do("HGETALL", "mp"+strconv.Itoa(i)))
		checkouterr(err)
		fmt.Println("第", i, "次:", v)
	}

}
func FindMysql() interface{} {
	//cashier := CashierData{}
	rows, _ := db.Query("select * from cashier where mch_id=1276316801")
	checkouterr(err)
	columns, _ := rows.Columns()
	fmt.Println(len(columns))
	values := make([]sql.RawBytes, len(columns)) //定义切片，用了存放最后的查询结果，类型为[]byte
	scanArgs := make([]interface{}, len(values)) //定义切片，用了对应每个切片内元素的地址
	for i := range values {
		scanArgs[i] = &values[i] //应该是查询只能使用指针型，将两个切片对应起来
	}
	r := []string{}
	m := make(map[string]string)
	n := make([]map[string]string, 0)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for k, v := range values {
			r = append(r, string(v))
			m[columns[k]] = string(v)
		}
		n = append(n, m) //可以通过反射将结果存到结构体内部，反射速度慢，要分类插入，代码量巨大，可以直接使用orm

	}
	//fmt.Println(n)
	return n
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
	//c := []A{}
	a := A{}
	//b := []string{"Jim", "15", "Tom", "16", "Mary", "17"}

	v := reflect.ValueOf(a)
	fmt.Println(v)
	v1 := reflect.Indirect(v)
	fmt.Println(v1)
}
