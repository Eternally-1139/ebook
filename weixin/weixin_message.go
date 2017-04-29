package weixin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"sort"
	"time"
)

//自动回复
type Auto_reply struct {
	Wx_id          string `json:"wx_id"`
	Auto_name      string `json:"auto_name"`
	Auto_type      string `json:"auto_type"`
	Auto_key       string `json:"auto_key"`
	Auto_data      string `json:"auto_data"`
	Auto_data_type string `json:"auto_data_type"`
}

//请求参数
type Request struct {
	XMLName                xml.Name `xml:"xml"`
	ToUserName             string
	FromUserName           string
	CreateTime             time.Duration
	MsgType                string
	Content                string
	MediaId                string
	Location_X, Location_Y float32
	Scale                  int
	Label                  string
	PicUrl                 string
	MsgId                  int
	Event                  string
	EventKey               string
	Ticket                 string
	Title                  string
	Description            string
}

//媒体ID
type M_MediaId struct {
	MediaId string `json:"media_id,omitempty"`
}

//文字消息
type M_text struct {
	Content string `json:"content,omitempty"`
}

//视频类
type M_Video struct {
	Title       string `json:"title,omitempty"`
	MediaId     string `json:"media_id,omitempty"`
	Description string `json:"description,omitempty"`
}

//音乐类
type M_Music struct {
	Title        string //音乐标题
	Description  string //音乐描述
	MusicURL     string //音乐链接
	HQMusicUrl   string //高质量音乐链接，WIFI环境优先使用该链接播放音乐
	ThumbMediaId string //缩略图的媒体id，通过素材管理接口上传多媒体文件，得到的id
}

//文章类
type M_Articles struct {
	Articles []Article `xml:"Articles,omitempty"`
	MediaId  string    `xml:"MediaId,omitempty";json:"media_id,omitempty"`
}

//文章类
type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

//返回参数
type ST_Response struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   string        `xml:",omitempty"`
	FromUserName string        `xml:",omitempty"`
	CreateTime   time.Duration `xml:",omitempty"`
	MsgType      string        `xml:",omitempty"`
	Content      string        `xml:",omitempty"`
	Image        *M_MediaId    `xml:",omitempty"`
	Voice        *M_MediaId    `xml:",omitempty"`
	Music        *M_Music      `xml:",omitempty"`
	Video        *M_Video      `xml:",omitempty"`

	ArticleCount int         `xml:",omitempty"`
	Articles     *M_Articles `xml:",omitempty"`
}

//创建自动回复文字类消息
func NewTextResponse(req *Request, Content string) (res interface{}, err error) {
	resp := ST_Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Text
	resp.Content = Content
	return resp, nil
}

//创建自动回复图片类消息
func NewImageResponse(req *Request, MediaId string) (res interface{}, err error) {
	resp := ST_Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Image
	resp.Image.MediaId = MediaId
	return resp, nil
}

//创建自动回复语音类消息
func NewVoiceResponse(req *Request, MediaId string) (res interface{}, err error) {
	resp := ST_Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = "voice"
	resp.Voice.MediaId = MediaId
	return resp, nil
}

//创建自动回复音乐类消息
func NewMusicResponse(req *Request, Title, Description, MusicURL, HQMusicUrl, ThumbMediaId string) (res interface{}, err error) {
	resp := ST_Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Music
	resp.Music.Title = Title
	resp.Music.Description = Description
	resp.Music.MusicURL = MusicURL
	resp.Music.HQMusicUrl = HQMusicUrl
	resp.Music.ThumbMediaId = ThumbMediaId
	return resp, nil
}

//创建自动回复视频类消息
func NewVideoResponse(req *Request, Title, Description, MediaId string) (res interface{}, err error) {
	resp := ST_Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = "video"
	resp.Video.MediaId = MediaId
	resp.Video.Title = Title
	resp.Video.Description = Description
	return resp, nil
}

//创建自动回复文章类消息
func NewArticlesResponse(req *Request, data []Article) (res interface{}, err error) {
	resp := ST_Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = News
	resp.ArticleCount = len(data)
	Articles:=make([]Article, len(data))
	for i, a := range data {
		Articles[i]=a
	}
	m:=M_Articles{Articles:Articles}
	resp.Articles=&m
	return resp, nil
}

//处理微信自动回复消息
func (wx *Weixin) Dealwith(req *Request) (resp_b interface{}, err error) {
	if req.MsgType == Event {
		d, e := wx.event_funcall(req)
		if d != nil {
			return d, e
		}
	}

	if _, ok := wx.FunCall["all_message"]; ok {
		log.Println("call all_message")
		data, err := wx.FunCall["all_message"](wx, req)
		if err == nil {
			return data, nil
		}
	}
	return wx.auto_replay(req)
}

//查询处理结果
func (wx *Weixin) auto_replay(message *Request) (resp_b interface{}, err error) {
	switch message.MsgType {
	case "text":
		text := wx.AutoReply["text"]
		if v, o := text[message.Content]; o {
			return Creat_wx_msg(message.FromUserName, message.ToUserName, v)
		}
	}
	switch message.Event {
	case "subscribe":
		text := wx.AutoReply["subscribe"]
		if v, o := text["default"]; o {
			return Creat_wx_msg(message.FromUserName, message.ToUserName, v)
		}
	case "CLICK":
		text := wx.AutoReply["CLICK"]
		if v, o := text[message.EventKey]; o {
			return Creat_wx_msg(message.FromUserName, message.ToUserName, v)
		}
	}

	return nil, nil
}

//事件类消息回调函数处理
func (wx *Weixin) event_funcall(req *Request) (resp_b interface{}, err error) {
	if req.Event == "subscribe" {
		if _, ok := wx.FunCall["subscribe"]; ok {
			log.Println("call subscribe")
			return wx.FunCall["subscribe"](wx, req)
		}
	}
	if req.Event == "unsubscribe" {
		if _, ok := wx.FunCall["unsubscribe"]; ok {
			log.Println("call unsubscribe")
			return wx.FunCall["unsubscribe"](wx, req)

		}
	}
	if req.Event == "SCAN" {
		if _, ok := wx.FunCall["Event_SCAN"]; ok {
			log.Println("call Event_SCAN")
			return wx.FunCall["Event_SCAN"](wx, req)

		}
	}
	if req.Event == "ShakearoundLotteryBind" {
		if _, ok := wx.FunCall["all_message"]; ok {
			log.Println("call ShakearoundLotteryBind")
			return wx.FunCall["ShakearoundLotteryBind"](wx, req)

		}
	}
	if req.Event == "CLICK" {
		if _, ok := wx.FunCall["Event_CLICK"]; ok {
			log.Println("call CLICK")
			return wx.FunCall["Event_CLICK"](wx, req)
		}
	}
	if req.Event == "VIEW" {
		if _, ok := wx.FunCall["Event_VIEW"]; ok {
			log.Println("call Event_VIEW")
			return wx.FunCall["Event_VIEW"](wx, req)

		}
	}
	if _, ok := wx.FunCall["Event"]; ok {
		log.Println("call Event")
		return wx.FunCall["Event"](wx, req)

	}
	return nil, nil
}

//检查微信消息签名
func (wx_info *Weixin) WeixinCheckSignature(signature, timestamp, nonce string) bool {

	tmps := []string{wx_info.Wx_token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := str2sha1(tmpStr)
	if tmp == signature {
		return true
	}
	return false
}

type ST_filter struct {
	Is_to_all bool `json:"is_to_all"`
	Tag_id    int  `json:"tag_id,omitempty"`
}

func Creat_wx_msg(openid, from string, data Auto_reply) (res ST_Response, err error) {
	res.FromUserName = from
	res.ToUserName = openid
	res.CreateTime = time.Duration(time.Now().Unix())
	res.MsgType = data.Auto_data_type

	switch data.Auto_data_type {
	case "text":
		res.Content = data.Auto_data
	case "image":
		image := &M_MediaId{data.Auto_data}
		res.Image = image
	case "voice":
		voice := &M_MediaId{data.Auto_data}
		res.Voice = voice
	case "news_url":
		res.MsgType = "news"
		res.Articles = new(M_Articles)

		type st_temp struct {
			Title   string `json:"title"`
			Digest  string `json:"digest"`
			Img_Url string `json:"img_url"`
			Url     string `json:"url"`
		}

		var t_d []st_temp
		json.Unmarshal([]byte(data.Auto_data), &t_d)

		res.ArticleCount = len(t_d)
		for _, d := range t_d {
			temp := Article{xml.Name{}, d.Title, d.Digest, d.Img_Url, d.Url}
			res.Articles.Articles = append(res.Articles.Articles, temp)
		}

		//	case "news":
		//		news_list, _ := Sel_Wx_news_replay("main", data.Auto_data)
		//		res.ArticleCount = len(news_list)
		//		res.Articles = new(M_Articles)
		//		for _, news := range news_list {
		//			temp := Article{xml.Name{}, news.News_title, news.News_digest, news.Material_url, news.News_url}
		//			res.Articles.Articles = append(res.Articles.Articles, temp)
		//		}
	default:
		return res, errors.New("未知联系")
	}
	return res, nil
}

//type	媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb），news，即图文消息
//errcode	错误码
//errmsg	错误信息
//msg_id	消息发送任务的ID
//msg_data_id
type Send_resp struct {
	Type        string `json:"type"`
	Msg_id      int64  `json:"msg_id"`
	Msg_data_id int64  `json:"msg_data_id"`
}

//https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func (wx *Weixin) Send_all(Tag_id int, Msgtype, Content string) (s_data Send_resp, err error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	var Filter ST_filter
	if Tag_id == 0 {
		Filter.Is_to_all = true
	} else {
		Filter.Is_to_all = false
		Filter.Tag_id = Tag_id
	}
	var data interface{}
	switch Msgtype {
	case "mpnews":
		type st_temp struct {
			Filter  ST_filter `json:"filter"`
			Mpnews  M_MediaId `json:"mpnews"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{Filter, M_MediaId{Content}, Msgtype}
		data = t
	case "text":
		type st_temp struct {
			Filter  ST_filter `json:"filter"`
			Text    M_text    `json:"text"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{Filter, M_text{Content}, Msgtype}
		data = t
	case "image":
		type st_temp struct {
			Filter  ST_filter `json:"filter"`
			Image   M_MediaId `json:"image"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{Filter, M_MediaId{Content}, Msgtype}
		data = t
	case "voice":
		type st_temp struct {
			Filter  ST_filter `json:"filter"`
			Voice   M_MediaId `json:"voice"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{Filter, M_MediaId{Content}, Msgtype}
		data = t
	case "mpvideo":
		type st_temp struct {
			Filter  ST_filter `json:"filter"`
			Mpvideo M_MediaId `json:"mpvideo"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{Filter, M_MediaId{Content}, Msgtype}
		data = t
	default:
		return s_data, errors.New("暂不支持该类型")
	}

	d, _ := json.Marshal(data)
	log.Println(string(d))
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/message/mass/sendall", param),
		bytes.NewReader(d))
	if err != nil {
		log.Println(err)
		return s_data, err
	}
	resp, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err, string(resp))
		return s_data, err
	}
	err = json.Unmarshal(resp, &s_data)
	if err != nil {
		log.Println(err)
		return s_data, err
	}
	return s_data, nil
}

//https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func (wx *Weixin) Send_list(user_lisr []string, Msgtype, Content string) (s_data Send_resp, err error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	var data interface{}
	switch Msgtype {
	case "mpnews":
		type st_temp struct {
			Touser  []string  `json:"touser"`
			Mpnews  M_MediaId `json:"mpnews"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{user_lisr, M_MediaId{Content}, Msgtype}
		data = t
	case "text":
		type st_temp struct {
			Touser  []string `json:"touser"`
			Text    M_text   `json:"text"`
			Msgtype string   `json:"msgtype"`
		}
		t := st_temp{user_lisr, M_text{Content}, Msgtype}
		data = t
	case "image":
		type st_temp struct {
			Touser  []string  `json:"touser"`
			Image   M_MediaId `json:"image"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{user_lisr, M_MediaId{Content}, Msgtype}
		data = t
	case "voice":
		type st_temp struct {
			Touser  []string  `json:"touser"`
			Voice   M_MediaId `json:"voice"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{user_lisr, M_MediaId{Content}, Msgtype}
		data = t
	case "mpvideo":
		type st_temp struct {
			Touser  []string  `json:"touser"`
			Mpvideo M_MediaId `json:"mpvideo"`
			Msgtype string    `json:"msgtype"`
		}
		t := st_temp{user_lisr, M_MediaId{Content}, Msgtype}
		data = t
	default:
		return s_data, errors.New("暂不支持该类型")
	}

	d, _ := json.Marshal(data)
	log.Println(string(d))
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/message/mass/send", param),
		bytes.NewReader(d))
	if err != nil {
		log.Println(err)
		return s_data, err
	}
	resp, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return s_data, err
	}
	err = json.Unmarshal(resp, &s_data)
	if err != nil {
		log.Println(err)
		return s_data, err
	}
	return s_data, nil
}

//https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func (wx *Weixin) Delete_mass(msg_id string) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	t := M_MediaId{msg_id}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/message/mass/delete", param),
		bytes.NewReader(d))
	if err != nil {
		log.Println(err)
		return 0
	}
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1

}

//https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
func (wx *Weixin) Preview_mass(openid, wxname, Msgtype, Content string) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	var data interface{}
	switch Msgtype {
	case "mpnews":
		type st_temp struct {
			Touser   string    `json:"touser,omitempty"`
			Towxname string    `json:"towxname,omitempty"`
			Mpnews   M_MediaId `json:"mpnews"`
			Msgtype  string    `json:"msgtype"`
		}
		t := st_temp{openid, wxname, M_MediaId{Content}, Msgtype}
		data = t
	case "text":
		type st_temp struct {
			Touser   string `json:"touser,omitempty"`
			Towxname string `json:"towxname,omitempty"`
			Text     M_text `json:"text"`
			Msgtype  string `json:"msgtype"`
		}
		t := st_temp{openid, wxname, M_text{Content}, Msgtype}
		data = t
	case "image":
		type st_temp struct {
			Touser   string    `json:"touser,omitempty"`
			Towxname string    `json:"towxname,omitempty"`
			Image    M_MediaId `json:"image"`
			Msgtype  string    `json:"msgtype"`
		}
		t := st_temp{openid, wxname, M_MediaId{Content}, Msgtype}
		data = t
	case "voice":
		type st_temp struct {
			Touser   string    `json:"touser,omitempty"`
			Towxname string    `json:"towxname,omitempty"`
			Voice    M_MediaId `json:"voice"`
			Msgtype  string    `json:"msgtype"`
		}
		t := st_temp{openid, wxname, M_MediaId{Content}, Msgtype}
		data = t
	case "mpvideo":
		type st_temp struct {
			Touser   string    `json:"touser,omitempty"`
			Towxname string    `json:"towxname,omitempty"`
			Mpvideo  M_MediaId `json:"mpvideo"`
			Msgtype  string    `json:"msgtype"`
		}
		t := st_temp{openid, wxname, M_MediaId{Content}, Msgtype}
		data = t
	default:
		return 0
	}

	d, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/message/mass/preview", param),
		bytes.NewReader(d))
	if err != nil {
		log.Println(err)
		return 0
	}
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1

}

//https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
func (wx *Weixin) Send_message(data interface{}) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	d, _ := json.Marshal(data)
	req, err := http.NewRequest("POST",
		Param("https://api.weixin.qq.com/cgi-bin/message/custom/send", param),
		bytes.NewReader(d))
	if err != nil {
		log.Println(err)
		return 0
	}
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1

}
