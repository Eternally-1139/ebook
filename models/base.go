package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"net/url"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
)

var o orm.Ormer

func init() {
	orm.RegisterModel(
		new(Address),
		new(Appointment),
		new(Carouse),
		new(Category),
		new(Collection),
		new(Message),
		new(Order),
		new(Product),
		new(User),
		new(File),
		new(Pic),
		new(ProductInfo),
		new(Comment),
		new(Weixin),
		)
}

func AddSystemAdmin() {
	var user User
	fmt.Println("please input username for system administrator")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("please input password for system administrator")
	var password string
	fmt.Scanf("%s", &password)


	user.Name = "SyS"
	user.Username = name
	user.Password = StrToMd5Sha1(password)
	user.CreateIp = "localhost"
	user.Account = 0
	user.Address = nil
	user.CreateTime = time.Now()
	user.LastLoginTime = time.Now()
	user.LastLoginIp = "localhost"
	user.Mark = 0
	user.Mobile = ""

	if _, err := user.Insert(); err != nil {
		fmt.Println("admin create error,please run this application again")
	} else {
		fmt.Println("admin create finished")
		fmt.Println("管理员账户："+user.Username)
		fmt.Println("管理员密码："+user.Password)

	}

}

func Syncbd(force bool) {
	creatDb()
	Connect()
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	if force {
		AddSystemAdmin()
	}
	fmt.Println("database init is complete.\nPlease restart the application")
}

func creatDb() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	var dns string
	var sqlstring string
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&loc=%s", db_user, db_pass, db_host, db_port, url.QueryEscape("Asia/Shanghai"))
	sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci", db_name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()
}

func Connect() {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s", db_user, db_pass, db_host, db_port, db_name, url.QueryEscape("Asia/Shanghai"))
	orm.RegisterDataBase("default", "mysql", dns)
}

func CountObjects(qs orm.QuerySeter) (int64, error) {
	cnt, err := qs.Count()
	if err != nil {
		beego.Error("models.CountObjects ", err)
		return 0, err
	}
	return cnt, err
}


func ListObjects(qs orm.QuerySeter, objs interface{}) (int64, error) {
	nums, err := qs.All(objs)
	if err != nil {
		beego.Error("models.ListObjects ", err)
		return 0, err
	}
	return nums, err
}


func StrToMd5Sha1(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	t := sha1.New();
	t.Write([]byte(rs))
	rss := hex.EncodeToString(t.Sum(nil))
	return rss
}