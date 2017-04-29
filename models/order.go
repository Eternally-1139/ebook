package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

/*
订单状态：Status
0:已删除
1:待支付
2:已支付
3:支付异常
支付方式：
0：微信支付
1：储值支付
2：积分混合支付
3：未选择
支付形式（payType）：
0：支付
1：租赁
*/

type Order struct {
	Id         int64
	User       *User `orm:"rel(fk)"`
	OddNumber  string `orm:"default(0)"`
	Status     int `orm:"default(1)"`
	PayMethod  int	`orm:"default(0)"`
	Price 	   float64 `orm:"default(0)"`
	PayType    int   `orm:"default(3)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	PayTime    time.Time `orm:"type(datetime);null"`
	CreateIp   string
	Address    string `orm:"null"`
}

func (this *Order) Insert () error {
	_,err:=orm.NewOrm().Insert(this)
	return err
}

func (this *Order) Read() error {
	if err:=orm.NewOrm().Read(this);err !=nil{
		return err
	}
	if err:=orm.NewOrm().Read(this.User);err !=nil{
		return err
	}

	return nil
}
func (this *Order) Delete ()(int64,error){
	return orm.NewOrm().Delete(this)
}

func (this *Order) Update ()(int64,error){
	return orm.NewOrm().Update(this)
}
