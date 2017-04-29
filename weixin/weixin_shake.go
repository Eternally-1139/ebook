package weixin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"
)

type Shake_Page struct {
	Page_id     int    `json:"page_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Page_url    string `json:"page_url,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Icon_url    string `json:"icon_url,omitempty"`
}

type lotteryinfo_req struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Onoff         int    `json:"onoff"`
	Begin_time    int64  `json:"begin_time"`
	Expire_time   int64  `json:"expire_time"`
	Sponsor_appid string `json:"sponsor_appid"`
	Total         int64  `json:"total"`
	Jump_url      string `json:"jump_url"`
	Key           string `json:"key"`
}

type lotteryinfo_res struct {
	Errcode    int    `json:"errcode"`
	Lottery_id string `json:"lottery_id"`
	Page_id    int    `json:"page_id"`
}

type hbpreorder_req struct {
	XMLName      xml.Name `xml:"xml"`
	Nonce_str    string   `xml:"nonce_str"`
	Sign         string   `xml:"sign"`
	Mch_billno   string   `xml:"mch_billno"`
	Mch_id       string   `xml:"mch_id"`
	Wxappid      string   `xml:"wxappid"`
	Send_name    string   `xml:"send_name"`
	Hb_type      string   `xml:"hb_type"`
	Total_amount int      `xml:"total_amount"`
	Total_num    int      `xml:"total_num"`
	Amt_type     string   `xml:"amt_type"`
	Wishing      string   `xml:"wishing"`
	Act_name     string   `xml:"act_name"`
	Remark       string   `xml:"remark"`
	Auth_mchid   string   `xml:"auth_mchid"`
	Auth_appid   string   `xml:"auth_appid"`
	Risk_cntl    string   `xml:"risk_cntl"`
}

type hbpreorder_res struct {
	XMLName     xml.Name `xml:"xml"`
	Sp_ticket   string   `xml:"sp_ticket"`
	Detail_id   string   `xml:"detail_id"`
	Send_time   string   `xml:"send_time"`
	Return_code string   `xml:"return_code"`
	Result_code string   `xml:"result_code"`
}

type ticket_t struct {
	Ticket string `json:"ticket"`
}

type prizebucket_req struct {
	Lottery_id      string     `json:"lottery_id"`
	Mchid           string     `json:"mchid"`
	Sponsor_appid   string     `json:"sponsor_appid"`
	Prize_info_list []ticket_t `json:"prize_info_list"`
}

type prizebucket_res struct {
	Errcode                     int        `json:"errcode"`
	Repeat_ticket_list          []ticket_t `json:"repeat_ticket_list"`
	Expire_ticket_list          []ticket_t `json:"expire_ticket_list"`
	Invalid_amount_ticket_list  []ticket_t `json:"invalid_amount_ticket_list"`
	Success_num                 int        `json:"success_num"`
	Wrong_authmchid_ticket_list []ticket_t `json:"wrong_authmchid_ticket_list"`
	Invalid_ticket_list         []ticket_t `json:"invalid_ticket_list"`
}

type Shake_device struct {
	Device_id int    `json:"device_id"`
	Uuid      string `json:"uuid"`
	Major     int    `json:"major"`
	Minor     int    `json:"minor"`
}

type Shake_device_res struct {
	Apply_id      int    `json:"apply_id"`
	Apply_time    int64  `json:"apply_time"`
	Audit_status  int    `json:"audit_status"`
	Audit_comment string `json:"audit_comment"`
	Audit_time    int64  `json:"audit_time"`
}

//申请设备ID
func (wx *Weixin) Add_shake_device(quantity int, apply_reason, comment string, poi_id int) (res Shake_device_res) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	type st_req struct {
		Quantity     int    `json:"quantity"`
		Apply_reason string `json:"apply_reason"`
		Comment      string `json:"comment,omitempty"`
		Poi_id       int    `json:"poi_id,omitempty"`
	}
	t := st_req{quantity, apply_reason, comment, poi_id}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/device/applyid", param),
		bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return res
	}

	type st_res struct {
		Data Shake_device_res `json:"data"`
	}
	var data st_res
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return res
	}
	return data.Data
}

//查询设备ID申请审核状态
func (wx *Weixin) Chkeck_shake_device(apply_id int) (res Shake_device_res) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	type st_req struct {
		Apply_id int `json:"apply_id"`
	}
	t := st_req{apply_id}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/device/applystatus", param),
		bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return res
	}

	type st_res struct {
		Data Shake_device_res `json:"data"`
	}
	var data st_res
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return res
	}
	return data.Data
}

//添加一个页面
func (wx *Weixin) Add_shake_page(title, description, page_url, comment, icon_url string) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	t := Shake_Page{0, title, description, page_url, comment, icon_url}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/page/add", param),
		bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	var data Shake_Page
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return 0
	}
	return data.Page_id
}

//编辑一个页面
func (wx *Weixin) Update_shake_page(page_id int, title, description, page_url, comment, icon_url string) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	t := Shake_Page{page_id, title, description, page_url, comment, icon_url}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/page/update", param),
		bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	var data Shake_Page
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return 0
	}
	return data.Page_id
}

//查询页面
func (wx *Weixin) Search_shake_page(type_id int, page_ids []int, begin, count int) []Shake_Page {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	type st_temp struct {
		Type     int   `json:"type"`
		Page_ids []int `json:"page_ids,omitempty"`
		Begin    int   `json:"begin,omitempty"`
		Count    int   `json:"count,omitempty"`
	}

	t := st_temp{type_id, page_ids, begin, count}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/page/search", param),
		bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return nil
	}
	type st_page struct {
		Pages []Shake_Page `josn:"pages"`
	}
	type st_data struct {
		Data st_page `josn:"data"`
	}
	var data st_data
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data.Data.Pages
}

//删除一个页面
func (wx *Weixin) Delete_shake_page(page_id int) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	t := Shake_Page{page_id, "", "", "", "", ""}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/page/delete", param),
		bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//上传摇一摇图片素材
func (wx *Weixin) Upload_shake_image(mediaType, filepath string) string {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["type"] = mediaType
	req, err := newfileUploadRequest(Param("https://api.weixin.qq.com/shakearound/material/add", param), nil,
		"media", filepath)
	if err != nil {
		return ""
	}
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return ""
	}
	type st_pic struct {
		Pic_url string `json:"pic_url"`
	}
	type st_res struct {
		Data st_pic `json:"data"`
	}
	var data st_res
	err = json.Unmarshal(res_body, &data)
	if err != nil {
		log.Println(err)
		return ""
	}
	return data.Data.Pic_url
}

//配置设备与页面的关联关系

//红包预下单
//func hbpreorder(wx_info *Weixin, hb_id int) int {
//	Wx_hb_info, sta := Get_Wx_hb_Info(hb_id)
//	if sta == 0 {
//		return 0
//	}
//	var h_hb hbpreorder_req
//	h_hb.Act_name = Wx_hb_info.Act_name
//	time_int64 := strconv.FormatInt(time.Now().UnixNano(), 10)
//	h_hb.Mch_billno = wx_info.Mch_id + time.Now().Format("20060102") + time_int64[len(time_int64)-10:]
//	h_hb.Mch_id = wx_info.Mch_id
//	h_hb.Nonce_str = creat_nonce_str()
//	h_hb.Remark = Wx_hb_info.Remark
//	h_hb.Send_name = Wx_hb_info.Send_name
//	h_hb.Total_amount = Wx_hb_info.Total_amount
//	h_hb.Total_num = Wx_hb_info.Total_num
//	h_hb.Wishing = Wx_hb_info.Wishing
//	h_hb.Wxappid = wx_info.Wx_appid
//	h_hb.Amt_type = Wx_hb_info.Amt_type
//	h_hb.Risk_cntl = "NORMAL"
//	h_hb.Auth_appid = "wxbf42bd79c4391863"
//	h_hb.Auth_mchid = "1000052601"
//	h_hb.Hb_type = Wx_hb_info.Hb_type
//	h_hb.Sign = Sign_v_md5(h_hb, wx_info.Pay_key)

//	data, err := xml.MarshalIndent(&h_hb, "", " ")
//	if err != nil {
//		log.Println(err)
//		return 0
//	}

//	res, err := SecurePost("https://api.mch.weixin.qq.com/mmpaymkttransfers/hbpreorder", data, "text/xml")
//	if err != nil {
//		beego.Warn(res, err)
//		return 0
//	}
//	res_body := make([]byte, 1024)
//	res.Body.Read(res_body)
//	beego.Warn(string(res_body))
//	var h_res hbpreorder_res
//	if err = xml.Unmarshal(res_body, &h_res); err != nil {
//		log.Println(err)
//		return 0
//	}
//	if h_res.Return_code == "SUCCESS" && h_res.Result_code == "SUCCESS" {
//		add_hbpreorder(hb_id, h_res.Sp_ticket, h_res.Detail_id, h_res.Send_time)
//		return 1
//	}
//	return 0
//}

//创建活动
//func creat_lottery(wx_info *Weixin, lottery lotteryinfo_req, use_template int, logo_url string) int {
//	lottery.Key = wx_info.Pay_key
//	lottery.Sponsor_appid = wx_info.Wx_appid
//	data, err := json.Marshal(&lottery)
//	if err != nil {
//		return 0
//	}
//	res_body, err := post("https://api.weixin.qq.com/shakearound/lottery/addlotteryinfo?access_token="+
//		wx_info.Access_token+"&use_template="+strconv.Itoa(use_template)+"&logo_url="+logo_url, data, "application/json")
//	if err != nil {
//		beego.Warn(err)
//		return 0
//	}
//	var l_res lotteryinfo_res
//	if err = json.Unmarshal(res_body, &l_res); err != nil {
//		log.Println(err)
//		return 0
//	}
//	if l_res.Errcode == 0 {
//		add_lotteryinfo(wx_info.Wx_id, use_template, logo_url, "", lottery, l_res)
//		return 1
//	}
//	return 0
//}

////录入红包信息
//func setprizebucket(wx_info *Weixin, lottery_id string, prize_info_list []string) int {
//	var req_data prizebucket_req
//	req_data.Lottery_id = lottery_id
//	req_data.Mchid = wx_info.Mch_id
//	req_data.Sponsor_appid = wx_info.Wx_appid
//	for _, v := range prize_info_list {
//		var t ticket_t
//		t.Ticket = v
//		req_data.Prize_info_list = append(req_data.Prize_info_list, t)
//	}
//	data, err := json.Marshal(&req_data)
//	if err != nil {
//		return 0
//	}
//	log.Println(string(data))
//	res_body, err := post("https://api.weixin.qq.com/shakearound/lottery/setprizebucket?access_token="+wx_info.Access_token, data, "application/json")
//	if err != nil {
//		beego.Warn(err)
//		return 0
//	}
//	var p_res prizebucket_res
//	if err = json.Unmarshal(res_body, &p_res); err != nil {
//		log.Println(err)
//		return 0
//	}
//	if p_res.Errcode == 0 {
//		if p_res.Success_num != len(prize_info_list) {
//			if len(p_res.Expire_ticket_list) != 0 {
//				prize_info_list = del_ticket(p_res.Expire_ticket_list, prize_info_list)
//			}
//			if len(p_res.Invalid_amount_ticket_list) != 0 {
//				prize_info_list = del_ticket(p_res.Invalid_amount_ticket_list, prize_info_list)
//			}
//			if len(p_res.Invalid_ticket_list) != 0 {
//				prize_info_list = del_ticket(p_res.Invalid_ticket_list, prize_info_list)
//			}
//			if len(p_res.Repeat_ticket_list) != 0 {
//				prize_info_list = del_ticket(p_res.Repeat_ticket_list, prize_info_list)
//			}
//			if len(p_res.Wrong_authmchid_ticket_list) != 0 {
//				prize_info_list = del_ticket(p_res.Wrong_authmchid_ticket_list, prize_info_list)
//			}
//		}
//		for _, v := range prize_info_list {
//			set_prizebucket(lottery_id, v)
//		}
//		return len(prize_info_list)
//		//		add_lotteryinfo(use_template, logo_url, lottery, l_res)
//	}
//	return 0
//}

//设置红包活动抽奖开关
func (wx *Weixin) Setlotteryswitch(lottery_id string, onoff int) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["lottery_id"] = lottery_id
	param["onoff"] = strconv.Itoa(onoff)

	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/lottery/setlotteryswitch", param),
		nil)
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

//申请设备ID
func (wx *Weixin) d_applyid(quantity int, apply_reason, comment, poi_id string) int {
	type t_req struct {
		Quantity     int    `json:"quantity"`
		Apply_reason string `json:"apply_reason"`
		Comment      string `json:"comment"`
		Poi_id       string `json:"poi_id"`
	}
	t_r := t_req{quantity, apply_reason, comment, poi_id}
	data, _ := json.Marshal(t_r)
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/shakearound/device/applyid", param),
		bytes.NewReader(data))
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

	//	res_body, err := post("https://api.weixin.qq.com/shakearound/device/applyid?access_token="+wx_info.Access_token, data, "application/json")

	//	if err != nil {
	//		log.Println(err)
	//	}
	//	type m_Errcode struct {
	//		Errcode  int    `json:"errcode"`
	//		Errmsg   string `json:"errmsg"`
	//		Apply_id int    `json:"apply_id"`
	//	}
	//	var t m_Errcode
	//	json.Unmarshal(res_body, &t)
	//	err = json.Unmarshal(res_body, &t)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	if t.Errcode == 0 {
	//		return 1
	//	}
	//	return 0

}

//查询设备ID申请审核状态
func d_applystatus(wx_info *Weixin, apply_id int) int {
	//	res_body, err := post("https://api.weixin.qq.com/shakearound/device/applystatus?access_token="+wx_info.Access_token, []byte(`{"apply_id": `+strconv.Itoa(apply_id)+`}`), "application/json")

	//	if err != nil {
	//		log.Println(err)
	//	}
	//	type m_Errcode struct {
	//		Errcode       int    `json:"errcode"`
	//		Errmsg        string `json:"errmsg"`
	//		apply_time    int64
	//		audit_comment string
	//		audit_status  int
	//		audit_time    int64
	//	}
	//	var t m_Errcode
	//	json.Unmarshal(res_body, &t)
	//	err = json.Unmarshal(res_body, &t)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	if t.Errcode == 0 {
	//		return 1
	//	}
	return 0
}

//编辑设备信息https://api.weixin.qq.com/shakearound/device/update?access_token=ACCESS_TOKEN
//func d_update(wx_info *Weixin, device shakea_device, comment string) int {
//	type t_req struct {
//		Device_identifier shakea_device `json:"device_identifier"`
//		Comment           string        `json:"comment"`
//	}
//	t_r := t_req{device, comment}
//	data, _ := json.Marshal(t_r)
//	res_body, err := post("https://api.weixin.qq.com/shakearound/device/applyid?access_token="+wx_info.Access_token, data, "application/json")
//	if err != nil {
//		log.Println(err)
//	}
//	type m_Errcode struct {
//		Errcode  int    `json:"errcode"`
//		Errmsg   string `json:"errmsg"`
//		Apply_id int    `json:"apply_id"`
//	}
//	var t m_Errcode
//	json.Unmarshal(res_body, &t)
//	err = json.Unmarshal(res_body, &t)
//	if err != nil {
//		log.Println(err)
//	}
//	if t.Errcode == 0 {
//		return 1
//	}
//	return 0
//}

//配置设备与门店的关联关系https://api.weixin.qq.com/shakearound/device/bindlocation?access_token=ACCESS_TOKEN
//func d_bindlocation(wx_info *Weixin, device shakea_device, poi_id int) int {
//	type t_req struct {
//		Device_identifier shakea_device `json:"device_identifier"`
//		poi_id            int           `json:"comment"`
//	}
//	t_r := t_req{device, poi_id}
//	data, _ := json.Marshal(t_r)
//	res_body, err := post("https://api.weixin.qq.com/shakearound/device/applyid?access_token="+wx_info.Access_token, data, "application/json")
//	if err != nil {
//		log.Println(err)
//	}
//	type m_Errcode struct {
//		Errcode int    `json:"errcode"`
//		Errmsg  string `json:"errmsg"`
//	}
//	var t m_Errcode
//	err = json.Unmarshal(res_body, &t)
//	if err != nil {
//		log.Println(err)
//	}
//	if t.Errcode == 0 {
//		return 1
//	}
//	return 0
//}

//配置设备与页面的关联关系https://api.weixin.qq.com/shakearound/device/bindpage?access_token=ACCESS_TOKEN
func d_bindpage(wx_info *Weixin, device_id int, page_ids []int, Append int) int {
	//	type t_req struct {
	//		Device_identifier shakea_device `json:"device_identifier"`
	//		Page_ids          []int         `json:"page_ids"`
	//		Bind              int           `json:"bind"`
	//		Append            int           `json:"append"`
	//	}
	//	t_r := t_req{shakea_device{Device_id: device_id}, page_ids, 1, Append}
	//	d, _ := json.Marshal(t_r)
	//	res_body, err := post("https://api.weixin.qq.com/shakearound/device/bindpage?access_token="+wx_info.Access_token,
	//		d, "application/json")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	type m_Errcode struct {
	//		Errcode int    `json:"errcode"`
	//		Errmsg  string `json:"errmsg"`
	//	}
	//	var t m_Errcode
	//	err = json.Unmarshal(res_body, &t)
	//	if err != nil {
	//		log.Println(err, ":", res_body)
	//		return 0
	//	}
	//	if t.Errcode == 0 {
	//		return 1
	//	}
	return 0
}

//查询设备列表https://api.weixin.qq.com/shakearound/device/search?access_token=ACCESS_TOKEN
func d_search(wx_info *Weixin) {

}

//编辑页面信息https://api.weixin.qq.com/shakearound/page/update?access_token=ACCESS_TOKEN
func update_page(wx_info *Weixin, page_id int, title, description, page_url, comment, icon_url string) int {
	//	type t_req struct {
	//		Page_id     int    `json:"page_id"`
	//		Title       string `json:"title"`
	//		Description string `json:"description"`
	//		Page_url    string `json:"page_url"`
	//		Comment     string `json:"comment"`
	//		Icon_url    string `json:"icon_url"`
	//	}
	//	t_r := t_req{page_id, title, description, page_url, comment, icon_url}
	//	data, _ := json.Marshal(t_r)
	//	res_body, err := post("https://api.weixin.qq.com/shakearound/page/update?access_token="+wx_info.Access_token,
	//		data, "application/json")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	type m_Errcode struct {
	//		Errcode int    `json:"errcode"`
	//		Errmsg  string `json:"errmsg"`
	//	}
	//	var t m_Errcode
	//	err = json.Unmarshal(res_body, &t)
	//	if err != nil {
	//		log.Println(err, ":", res_body)
	//		return 0
	//	}
	//	if t.Errcode == 0 {
	//		return 1
	//	}
	return 0
}

//删除页面https://api.weixin.qq.com/shakearound/page/delete?access_token=ACCESS_TOKEN
func delete_page(wx_info *Weixin, page_id int) int {
	//	type t_req struct {
	//		Page_id int `json:"page_id"`
	//	}
	//	t_r := t_req{page_id}
	//	data, _ := json.Marshal(t_r)
	//	res_body, err := post("https://api.weixin.qq.com/shakearound/page/delete?access_token="+wx_info.Access_token,
	//		data, "application/json")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	type m_Errcode struct {
	//		Errcode int    `json:"errcode"`
	//		Errmsg  string `json:"errmsg"`
	//	}
	//	var t m_Errcode
	//	err = json.Unmarshal(res_body, &t)
	//	if err != nil {
	//		log.Println(err, ":", res_body)
	//		return 0
	//	}
	//	if t.Errcode == 0 {
	//		return 1
	//	}
	return 0
}

func del_ticket(list []ticket_t, ticket []string) []string {
	for _, v := range list {
		for i, t := range ticket {
			if t == v.Ticket {
				if i+1 != len(ticket) {
					ticket = append(ticket[:i], ticket[i+1:]...)
				} else {
					ticket = ticket[:i]
				}
			}
		}
	}
	return ticket
}

//func (this *Weixin) Add_lottery() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	var lottery lotteryinfo_req
//	lottery.Title = this.GetString("title")
//	lottery.Desc = this.GetString("desc")
//	lottery.Onoff, _ = this.GetInt("onoff")
//	loc, _ := time.LoadLocation("Asia/Shanghai")
//	b_t, err := time.ParseInLocation("2006-01-02 15:04:05", this.GetString("begin_time"), loc)
//	if err != nil {
//		jsondata["status"] = 0
//		jsondata["info"] = "开始时间格式错误"
//		return
//	}
//	lottery.Begin_time = b_t.Unix()
//	e_t, err := time.ParseInLocation("2006-01-02 15:04:05", this.GetString("expire_time"), loc)
//	if err != nil {
//		jsondata["status"] = 0
//		jsondata["info"] = "结束时间格式错误"
//		return
//	}
//	lottery.Expire_time = e_t.Unix()
//	lottery.Total, _ = this.GetInt64("total")
//	lottery.Jump_url = this.GetString("Jump_url")
//	use_template, err := this.GetInt("use_template")
//	if use_template != 1 && use_template != 2 {
//		jsondata["status"] = 0
//		return
//	}
//	logo_url := this.GetString("logo_url")
//	jsondata["status"] = creat_lottery(wx_info, lottery, use_template, logo_url)
//}

//func (this *Weixin) Set_lottery() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()

//	lottery_id := this.GetString("lottery_id")
//	onoff, _ := this.GetInt("onoff")
//	beego.Warn(onoff, lottery_id)
//	if setlotteryswitch(wx_info, lottery_id, onoff) == 1 {
//		set_lottery(wx_info.Wx_id, lottery_id, onoff)
//		jsondata["status"] = 1
//		return
//	}
//	jsondata["status"] = 0
//}

//func (this *Weixin) Mod_page() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	title := this.GetString("title")
//	description := this.GetString("des")
//	comment := this.GetString("comment")
//	icon_url := this.GetString("logo_url")
//	page_id, _ := this.GetInt("page_id")

//	if update_page(wx_info, page_id, title, description, "", comment, icon_url) == 1 {
//		mod_page(wx_info.Wx_id, page_id, title, description, "", comment, icon_url)
//		jsondata["status"] = 1
//	}
//	jsondata["status"] = 0

//}

//func (this *Weixin) Del_page() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()

//	page_id, _ := this.GetInt("page_id")
//	if delete_page(wx_info, page_id) == 1 {
//		del_page(wx_info.Wx_id, page_id)
//		jsondata["status"] = 1
//		return
//	}
//	jsondata["status"] = 0
//}

//func (this *Weixin) Get_All_Lottery() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	jsondata["aaData"], jsondata["status"] = get_all_lottery(wx_info.Wx_id)
//}

//func (this *Weixin) Get_Lottery() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	lottery_id := this.GetString("lottery_id")
//	jsondata["data"], jsondata["status"] = get_lottery(wx_info.Wx_id, lottery_id)
//}

//func (this *Weixin) Set_prizebucket() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	tickets := this.GetStrings("ticket")
//	lottery_id := this.GetString("lottery_id")
//	count := setprizebucket(wx_info, lottery_id, tickets)
//	if count == 0 {
//		jsondata["status"] = 0
//	} else {
//		jsondata["status"] = 1
//		jsondata["succ"] = count
//	}

//}

//func (this *Weixin) Get_All_Hbpreorder() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	hb_tpl, _ := get_all_hb_tpl(wx_info.Wx_id)
//	all_hb, _ := get_all_hbpreorder()
//	var res_d []T_hbpreorder
//	for _, hb := range all_hb {
//		for _, t := range hb_tpl {
//			if t.Hb_id == hb.Hb_id {
//				res_d = append(res_d, hb)
//			}
//		}
//	}
//	jsondata["aaData"] = res_d
//	jsondata["status"] = 1
//}

//func (this *Weixin) Get_Can_Hbpreorder() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()

//	hb_tpl, _ := get_all_hb_tpl(wx_info.Wx_id)
//	all_hb, _ := get_all_can_hbpreorder()
//	var res_d []T_hbpreorder
//	for _, hb := range all_hb {
//		for _, t := range hb_tpl {
//			if t.Hb_id == hb.Hb_id {
//				res_d = append(res_d, hb)
//			}
//		}
//	}
//	jsondata["aaData"] = res_d
//	jsondata["status"] = 1
//}

//func (this *Weixin) Hb_preorder() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	hb_id, _ := this.GetInt("hb_id")
//	count, _ := this.GetInt("count")
//	flag := 0
//	for i := 0; i < count; i++ {
//		if hbpreorder(wx_info, hb_id) == 1 {
//			flag++
//		}
//	}
//	jsondata["status"] = 1
//	jsondata["succ"] = flag
//}

//func (this *Weixin) Bindpage() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	device_id, _ := this.GetInt("device_id")
//	page_strs := this.GetStrings("page_id[]")
//	log.Println(page_strs)
//	flag, _ := this.GetInt("flag")
//	var page_ids []int
//	for _, page_str := range page_strs {
//		page_id, _ := strconv.Atoi(page_str)
//		if page_id != 0 {
//			page_ids = append(page_ids, page_id)
//		}
//	}
//	if d_bindpage(wx_info, device_id, page_ids, flag) == 1 {
//		bindpage(device_id, page_ids, flag)
//		jsondata["status"] = 1
//		return
//	}
//	jsondata["status"] = 0
//}

//func (this *Weixin) Add_shake_device() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	device_id, _ := this.GetInt("device_id")
//	uuid := this.GetString("uuid")
//	major, _ := this.GetInt("major")
//	minor, _ := this.GetInt("minor")
//	status, _ := this.GetInt("status")
//	comment := this.GetString("comment")
//	poi_id, _ := this.GetInt("poi_id")
//	jsondata["status"] = add_shake_device(wx_info.Wx_id, device_id, uuid, major, minor, status, comment, poi_id)
//}

//func (this *Weixin) Get_All_ShakeDevice() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	jsondata["aaData"], jsondata["status"] = get_all_shakedevice(wx_info.Wx_id)
//}

//func (this *Weixin) Get_ShakeDevice() {
//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	wx_info, ok := Wx_info[weixin_id]
//	if !ok {
//		this.Abort("404")
//		return
//	}
//	jsondata := make(map[string]interface{})
//	defer func() {
//		this.Data["json"] = jsondata
//		this.ServeJSON()
//	}()
//	device_id, _ := this.GetInt("device_id")
//	jsondata["data"], jsondata["status"] = get_shakedevice(wx_info.Wx_id, device_id)
//}

//func (this *Weixin) Upload_device() {
//	//	weixin_id := this.Ctx.Input.Param(":weixin_id")
//	//	wx_info, ok := Wx_info[weixin_id]
//	//	if !ok {
//	//		this.Abort("404")
//	//		return
//	//	}
//	//	jsondata := make(map[string]interface{})
//	//	defer func() {
//	//		this.Data["json"] = jsondata
//	//		this.ServeJSON()
//	//	}()
//	//	File_in, _, err := this.GetFile("file")
//	//	File_in.Close()
//	//	if err != nil {
//	//		log.Println(err)
//	//		jsondata["status"] = 0

//	//	}
//	//	var x_data []byte
//	//	File_in.Read(x_data)

//	//	xlFile, err := xlsx.OpenBinary(x_data)
//	//	if err != nil {
//	//		log.Println(err)
//	//		jsondata["status"] = 0
//	//		return
//	//	}
//	//	for _, sheet := range xlFile.Sheets {
//	//		for i, row := range sheet.Rows {
//	//			if i < 9 {
//	//				continue
//	//			}
//	//			var (
//	//				device_id int
//	//				uuid      string
//	//				major     int
//	//				minor     int
//	//				comment   string
//	//				poi_id    int
//	//			)
//	//			for j, cell := range row.Cells {
//	//				switch j {
//	//				case 0:
//	//					device_id, _ = cell.Int()
//	//				case 1:
//	//					uuid = cell.String()
//	//				case 2:
//	//					major, _ = cell.Int()
//	//				case 3:
//	//					minor, _ = cell.Int()
//	//				case 4:
//	//					comment = cell.String()
//	//				case 5:
//	//					poi_id, _ = cell.Int()
//	//				}
//	//			}
//	//			add_shake_device(wx_info.Wx_id, device_id, uuid, major, minor, 0, comment, poi_id)
//	//		}
//	//	}

//}

//最好修改成requset
//func post(url string, Content []byte, Content_Type string) ([]byte, error) {
//	client := &http.Client{}
//	req, err := client.Post(url, Content_Type, bytes.NewBuffer(Content))
//	if err != nil {
//		return nil, err
//	}
//	log.Println(url, string(Content))
//	var res_body []byte
//	for {
//		body := make([]byte, 1024)
//		n, err := req.Body.Read(body)
//		res_body = append(res_body, body[:n]...)
//		if err != nil {
//			break
//		}
//	}
//	log.Println(string(res_body))
//	return res_body, nil
//}
