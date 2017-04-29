package models

import (
	"time"
	"strings"
	"github.com/astaxie/beego/orm"
	"os"
)

type File struct {
	Id         int64
	FileName   string
	Type 	   string `orm:"null"`
	Hash       string `orm:"size(64)"`
	Size       int64
	UploadTime time.Time `orm:"type(datetime);auto_now_add"`
}

func (this *File) Insert() (int64, error) {
	if len(strings.Split(this.FileName,"."))>1{
		this.Type=strings.Split(this.FileName,".")[1]
	}
	return orm.NewOrm().Insert(this)
}

func (this *File) FindByHash() error {
	return orm.NewOrm().QueryTable("file").Filter("hash", this.Hash).One(this)
}


func (this *File) Read() error{
	return orm.NewOrm().QueryTable("file").Filter("id",this.Id).One(this)
}

func (this *File) Update() (int64,error){
	return orm.NewOrm().Update(this)
}

func (this *File) Delete()(int64,error) {
	i,err:=orm.NewOrm().Delete(this)
	if err!=nil{
		return 0,err
	}

	count,_:=orm.NewOrm().QueryTable("file").Filter("hash",this.Hash).Count()
	if  count< 1{
		filename:="web-static/admin/upload/file/"+this.Hash
		if err:=os.Remove(filename);err!=nil{
			return 0,err
		}

	}
	return i,err

}
