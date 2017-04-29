package admin

import (
	"github.com/astaxie/beego"
	"time"
	"os"
	"io"
	"crypto/md5"
	"encoding/hex"
	"ebook/models"
)


// @router /file/upload [post]
func (this *AjaxController) Upload() {
	_, fileHeader, err := this.GetFile("file")

	if err != nil {
		beego.Debug("get file error :", err)
	}
	result := make(map[string]interface{})
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		result["stauts"] = 10001
		result["message"] = "error"
		beego.Debug("Open file error :", err)

	}

	var inFile models.File
	filename := fileHeader.Filename + string(time.Now().Unix())

	f, err := os.OpenFile("web-static/admin/upload/file/"+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	io.Copy(f, file)
	fileinfo, err := f.Stat()
	if err != nil {
		beego.Debug("file info get error :", err)
	}

	filesize := fileinfo.Size()

	hashfile, err := os.Open("web-static/admin/upload/file/" + filename)

	hash := md5.New()
	if _, err := io.Copy(hash, hashfile); err != nil {
		this.ReturnJson(10002, "get hash error")
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	hashString := hex.EncodeToString(hashInBytes)
	hashfile.Close()

	if err := os.Rename("web-static/admin/upload/file/"+filename, "web-static/admin/upload/file/"+hashString); err != nil {
		beego.Debug("file  rename error :", err)
	}

	inFile.FileName = fileHeader.Filename
	inFile.Hash = hashString
	inFile.Size = filesize
	inFile.Insert()

	result["status"] = 10000
	result["message"] = "success"
	result["hash"] = hashString
	result["src"]="/web-static/admin/upload/file/"+hashString
	result["id"] = inFile.Id
	this.Data["json"] = result
	this.ServeJSON()
	return

}