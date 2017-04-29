package weixin

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"
)

type T_wx_hb_info struct {
	Hb_id        int    `json:"hb_id"`
	Send_name    string `json:"send_name"` //商户名称
	Act_name     string `json:"act_name"`  //活动名称
	Client_ip    string `json:"client_ip"` //Ip地址
	Wishing      string `json:"wishing"`   //红包祝福语
	Remark       string `json:"remark"`    //备注
	Hb_type      string `json:"hb_type"`
	Amt_type     string `json:"amt_type"`
	Total_num    int    `json:"total_num"`
	Total_amount int    `json:"total_amount"`
}

type T_lotteryinfo struct {
	Wx_id        string `json:"wx_id"`
	Lottery_id   string `json:"lottery_id"`
	Use_template int    `json:"use_template"`
	Logo_url     string `json:"logo_url"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Onoff        int    `json:"onoff"`
	Begin_time   int64  `json:"begin_time"`
	Expire_time  int64  `json:"expire_time"`
	Total        int64  `json:"total"`
	Jump_url     string `json:"jump_url"`
	Page_url     string `json:"page_url"`
	Page_id      int    `json:"page_id"`
}

type T_hbpreorder struct {
	Hb_id        int    `json:"hb_id"`
	Act_name     string `json:"act_name"`
	Total_amount int    `json:"total_amount"`
	Ticket       string `json:"ticket"`
	Detail_id    int    `json:"detail_id"`
	Send_time    string `json:"send_time"`
	Status       string `json:"status"`
}

type T_shake_device struct {
	Wx_id     string `json:"wx_id"`
	Device_id int    `json:"device_id"`
	Uuid      string `json:"uuid"`
	Major     int    `json:"major"`
	Minor     int    `json:"minor"`
	Status    int    `json:"status"`
	Comment   string `json:"comment"`
	Poi_id    int    `json:"poi_id"`
}

type T_bindpage struct {
	Device_id int `json:"device_id"`
	Page_id   int `json:"page_ids"`
}

type Callback_fun func(*Weixin, ...interface{}) (interface{}, error)

type template_value struct {
	Value string `json:"value"`
}

type template_data struct {
	First    template_value `json:"first"`
	Keyword1 template_value `json:"keyword1"`
	Keyword2 template_value `json:"keyword2"`
	Keyword3 template_value `json:"keyword3"`
	Keyword4 template_value `json:"keyword4"`
	Keyword5 template_value `json:"keyword5"`
	Remark   template_value `json:"remark"`
}

type T_wxerror struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Wx_Template struct {
	Touser      string        `json:"touser"`
	Template_id string        `json:"template_id"`
	Url         string        `json:"url"`
	Data        template_data `json:"data"`
}

type Menu_sub_button struct {
	Name string `json:"name"`
	Type string `json:"type"`
	key  string `json:"key"`
	Url  string `json:"url"`
}

type Menu_button struct {
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	key        string            `json:"key"`
	Url        string            `json:"url"`
	Sub_button []Menu_sub_button `json:"sub_button"`
}

type button struct {
	Button []Menu_button `json:"button"`
}

type Wx_Menu struct {
	Menu button `json:"menu"`
}

type wx_articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Picurl      string `json:"picurl"`
	Url         string `json:"url"`
}

type Access_Token struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Errcode      string `json:"errcode"`
}

type User_token struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Errcode      int    `json:"errcode"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Errmsg       string `json:"Errmsg"`
}

type Wx_Openid_list struct {
	Openid []string `json:"openid"`
}

type ticket struct {
	Errcode    int    `json:"errcode"`
	Ticket     string `json:"ticket"`
	Errmsg     string `json:"errmsg"`
	Expires_in int    `json:"expires_in"`
}

type Weixin struct {
	Wx_id             string
	Wx_appid          string
	Wx_appsecret      string
	Wx_token          string
	Wx_encodingaeskey string
	Access_token      string
	Access_token_time int64
	Wx_status         int
	Expires_in        int
	ac_token_count    int
	FunCall           map[string]Callback_fun
	Option            map[string]string
	AutoReply         map[string]map[string]Auto_reply
	Jsapi_ticket      string
	Jsapi_token_time  int64
	_tlsConfig        *tls.Config
}

type Qrcode_Q struct {
	Expire_seconds int    `json:"expire_seconds"`
	Action_name    string `json:"action_name"`
	Action_info    scene  `json:"action_info"`
}

type scene struct {
	Scene_id  int    `json:"scene_id"`
	Scene_str string `json:"scene_str"`
}

type Qrcode_S struct {
	Ticket         string `json:"ticket"`
	Expire_seconds int    `json:"expire_seconds"`
	Url            string `json:"url"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

//新建一个微信对象
func New(Wx_id, Wx_appid, Wx_appsecret, Wx_token, Wx_encodingaeskey, Access_token string,
	Access_token_time int64, Wx_status int, Option map[string]string, FunCall map[string]Callback_fun) (*Weixin, error) {
	weixin := new(Weixin)
	weixin.Wx_id = Wx_id
	weixin.Wx_token = Wx_token
	weixin.Wx_appid = Wx_appid
	weixin.Wx_appsecret = Wx_appsecret
	weixin.Wx_encodingaeskey = Wx_encodingaeskey
	weixin.Access_token = Access_token
	weixin.Access_token_time = Access_token_time
	weixin.Wx_status = Wx_status
	weixin.Expires_in = 3600
	weixin.FunCall = FunCall
	weixin.AutoReply = nil
	weixin.Option = Option
	if time.Now().Unix()-weixin.Access_token_time > 3600 {
		weixin.get_access_token()
	} else {
		check_access_token(weixin)
	}
	if get_wx_js_jsapi_ticket(weixin) == 0 {
		get_wx_js_jsapi_ticket(weixin)
	}
	go defend_wxinfo(weixin)
	return weixin, nil
}

//生成二维码参数
func (wx *Weixin) Create_qrcode(expire_seconds, scene_id int, scene_str string) (Qrcode_S, error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	var q Qrcode_Q
	if expire_seconds != 0 {
		q.Action_name = "QR_SCENE"
	} else if scene_str != "" {
		q.Action_name = "QR_LIMIT_STR_SCENE"
	} else {
		q.Action_name = "QR_LIMIT_SCENE"
	}
	q.Action_info.Scene_id = scene_id
	q.Action_info.Scene_str = scene_str
	q.Expire_seconds = expire_seconds

	d, err := json.Marshal(q)

	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/qrcode/create", param), bytes.NewReader(d))
	var qs Qrcode_S
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return qs, err
	}

	err = json.Unmarshal(res_body, &qs)
	return qs, err
}

//微信对象access_token维护
func defend_wxinfo(wx *Weixin) {
	for {
		if time.Now().Unix()-wx.Access_token_time > 3600 {
			wx.get_access_token()
			wx.Access_token_time = time.Now().Unix()
			wx.ac_token_count = 0
		} else {
			check_access_token(wx)
		}
		if time.Now().Unix()-wx.Jsapi_token_time > 3600 {
			if get_wx_js_jsapi_ticket(wx) == 0 {
				get_wx_js_jsapi_ticket(wx)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}

//检查微信access_token有效性
func check_access_token(wx *Weixin) int {

	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token="+
		wx.Access_token, nil)
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	//	log.Println(string(res))
	return 1
}

//生成请求url参数
func Param(url_base string, P map[string]string) string {
	for k, v := range P {
		if strings.Index(url_base, "?") != -1 {
			url_base += "&"
		} else {
			url_base += "?"
		}
		url_base += k
		url_base += "="
		url_base += url.QueryEscape(v)
	}
	return url_base
}

//获取js的jsapi_ticket
func get_wx_js_jsapi_ticket(wx *Weixin) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["type"] = "jsapi"
	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/ticket/getticket", param), nil)

	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	var temp_tick ticket
	err = json.Unmarshal(res_body, &temp_tick)
	if err != nil {
		log.Println(err)
		return 0
	} else if temp_tick.Errcode == 0 {
		wx.Jsapi_token_time = time.Now().Unix()
		wx.Jsapi_ticket = temp_tick.Ticket
		if _, ok := wx.FunCall["jsapi_ticket"]; ok {
			log.Println("call jsapi_ticket")
			wx.FunCall["jsapi_ticket"](wx)
		}
		return 1
	}
	return 0
}

//创建jsapi_ticket签名
func (wx *Weixin) Create_signature(url, noncestr string, timestamp int64, data map[string]interface{}) string {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["url"] = url
	data["noncestr"] = noncestr
	data["jsapi_ticket"] = wx.Jsapi_ticket
	data["timestamp"] = strconv.FormatInt(timestamp, 10)
	return Sign_v_sha1(data)
}

func Create_signature(wx_info *Weixin, noncestr, timestamp, url string) string {
	var str1 string
	str1 = "jsapi_ticket=" + wx_info.Jsapi_ticket + "&noncestr=" + noncestr + "&timestamp=" + timestamp + "&url=" + url
	return str2sha1(str1)
}

func Sign_v_sha1(data map[string]interface{}) string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str1 := ""
	for i := range keys {
		val := data[keys[i]]
		var str string
		switch val.(type) {
		case string:
			str = val.(string)
		case bool:
			str = strconv.FormatBool(val.(bool))
		case int:
			str = strconv.Itoa(val.(int))
		case int64:
			str = strconv.FormatInt(val.(int64), 10)
		case []byte:
			str = string(val.([]byte))
		default:
			continue
		}
		if len(str) == 0 {
			continue
		}
		if len(str1) != 0 {
			str1 += "&"
		}
		str1 += keys[i] + "=" + str
	}
	t := sha1.New()
	io.WriteString(t, str1)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//获取access_token
func (wx *Weixin) get_access_token() int {
	if wx.Wx_status == 0 || (time.Now().Unix()-wx.Access_token_time < 60 && wx.ac_token_count > 5) {
		return -1
	}

	param := make(map[string]string)
	param["grant_type"] = "client_credential"
	param["appid"] = wx.Wx_appid
	param["secret"] = wx.Wx_appsecret

	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/token", param), nil)

	res_body, err := requset(wx, req, -1)
	if err != nil {
		log.Println(err)
		return -1
	}
	log.Println(string(res_body))
	if err != nil {
		log.Println(err)
	}
	var ac_token Access_Token
	ac_token.Errcode = ""
	err = json.Unmarshal(res_body, &ac_token)
	if err != nil {
		log.Println(err)
		return -1
	}
	if ac_token.Errcode != "" {
		return -1
	}

	if time.Now().Unix()-wx.Access_token_time < 60 {
		wx.ac_token_count++
	} else {
		wx.ac_token_count = 0
	}

	wx.Access_token_time = time.Now().Unix()
	wx.Access_token = ac_token.Access_token
	wx.Expires_in = ac_token.Expires_in
	if _, ok := wx.FunCall["updatetoken"]; ok {
		wx.FunCall["updatetoken"](wx)
	}
	return 0
}

//string转sha1
func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//微信MD5签名
func Sign_v_md5(data interface{}, key string) string {
	k := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	var keys []string
	m := make(map[string]interface{})
	for i := 0; i < k.NumField(); i++ {
		chKey := k.Field(i).Tag.Get("xml")
		t_str := strings.Split(chKey, ",")
		keys = append(keys, t_str[0])
		m[t_str[0]] = v.Field(i).Interface()
	}
	sort.Strings(keys)
	str1 := ""
	for i := range keys {
		val := m[keys[i]]
		var str string
		switch val.(type) {
		case string:
			str = val.(string)
		case bool:
			str = strconv.FormatBool(val.(bool))
		case int:
			str = strconv.Itoa(val.(int))
		case int64:
			str = strconv.FormatInt(val.(int64), 10)
		case []byte:
			str = string(val.([]byte))
		default:
			continue
		}
		if len(str) == 0 {
			continue
		}
		if len(str1) != 0 {
			str1 += "&"
		}
		str1 += keys[i] + "=" + str
	}
	str1 += "&key=" + key
	t := md5.New()
	io.WriteString(t, str1)
	return fmt.Sprintf("%X", t.Sum(nil))
}

func DecodeRequest(data []byte) (req *Request, err error) {
	req = &Request{}
	if err = xml.Unmarshal(data, req); err != nil {
		return
	}
	req.CreateTime *= time.Second
	return
}

//发送微信请求
//flag = -1  忽略Access_token
//flag = 0  重新获取Access_token 并返回错误
//flag = 1  重新获取Access_token 并继续进行操作
func requset(wx *Weixin, req *http.Request, flag int) ([]byte, error) {
	if flag != -1 && time.Now().Unix()-wx.Access_token_time > int64(wx.Expires_in) {
		if wx.get_access_token() != 0 {
			return nil, errors.New("获取ac_token出错")
		}
		if flag == 0 {
			return nil, errors.New("ac_token出错")
		}
	}

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errcode T_wxerror
	err = json.Unmarshal(res_body, &errcode)
	if err == nil && errcode.Errcode != 0 {
		if errcode.Errcode == 40001 {
			go wx.get_access_token()
		}
		return res_body, errors.New(errcode.Errmsg)
	}

	return res_body, nil
}

// 随机字符串
/*Random = 0  // 纯数字
Random = 1  // 小写字母
Random = 2  // 大写字母
Random   = 3  // 数字、大小写字母*/
func Random_str(size int, Random int) string {
	iRandom, Randoms, result := Random, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := Random > 2 || Random < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			iRandom = rand.Intn(3)
		}
		scope, base := Randoms[iRandom][0], Randoms[iRandom][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
