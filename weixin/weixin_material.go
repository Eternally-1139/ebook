package weixin

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type media_res struct {
	Media_id string `josn:"media_id"`
	Url      string `json:"url"`
}

//图文消息
type T_News struct {
	Title              string `json:"title"`
	Thumb_media_id     string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Show_cover_pic     int    `json:"show_cover_pic"`
	Content            string `json:"content"`
	Url                string `json:"url"`
	Content_source_url string `json:"content_source_url"`
}

//图文信息返回
type ST_News struct {
	T_News
	Url       string `json:"url"`
	Thumb_url string `json:"thumb_url"`
}

//媒体返回
type T_Media struct {
	Type       string `json:"type"`
	Media_id   string `json:"media_id"`
	Created_at int64  `json:"created_at"`
	Url        string `json:"url"`
}

type ST_material_list struct {
	Total_count int           `json:"total_count"`
	Item_count  int           `json:"item_count"`
	Item        []ST_material `json:"item"`
	Content     []T_News      `json:"content"`
}

type ST_material struct {
	Media_id    string `json:"media_id"`
	Name        string `json:"name,omitempty"`
	Update_time int64  `json:"update_time,omitempty"`
	Url         string `json:"url,omitempty"`
}

type ST_news_item struct {
	Media_id    string          `json:"media_id"`
	Content     ST_news_content `json:"content"`
	Update_time int64           `json:"update_time"`
}

type ST_news_content struct {
	News_item []T_News `json:"news_item"`
}

type ST_res_news struct {
	Total_count int            `json:"total_count"`
	Item_count  int            `json:"item_count"`
	Item        []ST_news_item `json:"item"`
}

//{
//  "total_count": TOTAL_COUNT,
//  "item_count": ITEM_COUNT,
//  "item": [{
//      "media_id": MEDIA_ID,
//      "content": {
//          "news_item": [{
//              "title": TITLE,
//              "thumb_media_id": THUMB_MEDIA_ID,
//              "show_cover_pic": SHOW_COVER_PIC(0 / 1),
//              "author": AUTHOR,
//              "digest": DIGEST,
//              "content": CONTENT,
//              "url": URL,
//              "content_source_url": CONTETN_SOURCE_URL
//          },
//          //多图文消息会在此处有多篇文章
//          ]
//       },
//       "update_time": UPDATE_TIME
//   },
//   //可能有多个图文消息item结构
// ]
//}

//增加临时资源
func (wx *Weixin) Add_temp_material(mediaType, filepath string) (data T_Media, err error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["type"] = mediaType
	req, err := newfileUploadRequest(Param("https://api.weixin.qq.com/cgi-bin/media/upload", param), nil,
		"media", filepath)
	if err != nil {
		return
	}
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

//获取临时的媒体资源
func (wx *Weixin) Get_temp_material(basepath, media_id string) (string, error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["media_id"] = media_id

	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/media/get", param), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	file_name := ""
	Disposition := resp.Header.Get("Content-Disposition")
	if Disposition != "" {
		if strings.Index(Disposition, `attachment; filename="`) == 0 {
			file_name = Disposition[len(`attachment; filename="`) : len(Disposition)-1]
			log.Println(file_name)
		}
	}
	if file_name == "" {
		return "", errors.New("not a file")
	}
	file_name = path.Join(basepath, file_name)
	file, err := os.Create(file_name)
	if err != nil {
		log.Println(err)
		return file_name, err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err)
		return file_name, err
	}
	return file_name, err
}

//增加图文消息
func (wx *Weixin) Add_news(news []T_News) (data T_Media, err error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	type articles struct {
		Articles []T_News `json:"articles"`
	}
	t := articles{news}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/material/add_news", param),
		bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
	//	return res_body, err
}

//修改图文消息
func (wx *Weixin) Update_news(Media_id string, Index int, news T_News) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	type articles struct {
		Media_id string `json:"media_id"`
		Index    int    `json:"index"`
		Articles T_News `json:"articles"`
	}
	t := articles{Media_id, Index, news}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/material/update_news", param),
		bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		return 0
	}
	return 1
}

//增加图文消息图片
func (wx *Weixin) Uploadimg(filepath string) (data T_Media, err error) {
	//	https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	req, err := newfileUploadRequest(Param("https://api.weixin.qq.com/cgi-bin/media/uploadimg", param),
		nil, "media", filepath)
	if err != nil {
		return
	}
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

//增加资源
func (wx *Weixin) Add_material(mediaType, filepath string) (data T_Media, err error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["type"] = mediaType

	req, err := newfileUploadRequest(Param("https://api.weixin.qq.com/cgi-bin/material/add_material", param),
		nil, "media", filepath)
	if err != nil {
		return
	}
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

//删除资源
func (wx *Weixin) Del_material(media_id string) int {
	type st_temp struct {
		Media_id string `json:"media_id"`
	}
	t := st_temp{media_id}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token="+
		wx.Access_token, bytes.NewReader(d))

	_, err = requset(wx, req, 0)
	if err == nil {
		return 1
	}
	return 0
}

//获取除文章和视频类型外的媒体资源
func (wx *Weixin) Get_material(basepath, media_id string) (string, error) {
	//	https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN
	//	http.Get()
	type st_temp struct {
		Media_id string `json:"media_id"`
	}
	t := st_temp{media_id}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token="+
		wx.Access_token, bytes.NewReader(d))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	file_name := ""
	Disposition := resp.Header.Get("Content-Disposition")
	if Disposition != "" {
		if strings.Index(Disposition, `attachment; filename="`) == 0 {
			file_name = Disposition[len(`attachment; filename="`) : len(Disposition)-1]
			log.Println(file_name)
		}
	}
	if file_name == "" {
		return "", errors.New("not a file")
	}
	file_name = path.Join(basepath, file_name)
	file, err := os.Create(file_name)
	if err != nil {
		log.Println(err)
		return file_name, err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err)
		return file_name, err
	}
	return file_name, err
}

//获取图文资源
func (wx *Weixin) Get_materials_news(Type string, offset, count int) (data ST_res_news, err error) {
	type st_temp struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Count  int    `json:"count"`
	}
	t := st_temp{Type, offset, count}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token="+
		wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, err
}

//获取资源
func (wx *Weixin) Get_materials(Type string, offset, count int) (data ST_material_list, err error) {
	type st_temp struct {
		Type   string `json:"type"`
		Offset int    `json:"offset"`
		Count  int    `json:"count"`
	}
	t := st_temp{Type, offset, count}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token="+
		wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	//	log.Println(string(res_body))
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, err
}

//获取除文章类型外的资源列表
func (wx *Weixin) Get_all_materials(Type string) (materials []ST_material, err error) {
	offset := 0
	for {
		data, err := wx.Get_materials(Type, offset, 20)
		if err != nil {
			return materials, err
		}
		offset += data.Item_count
		materials = append(materials, data.Item...)
		if offset >= data.Total_count {
			break
		}
	}
	return materials, nil
}

//获取文章类型资源列表
func (wx *Weixin) Get_all_materials_news() (materials []ST_news_item, err error) {
	offset := 0
	for {
		data, err := wx.Get_materials_news("news", offset, offset+20)
		if err != nil {
			return materials, err
		}
		offset += data.Item_count
		materials = append(materials, data.Item...)
		if offset >= data.Total_count {
			break
		}
	}
	return materials, nil
}

//微信文件上传
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}
