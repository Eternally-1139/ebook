package weixin

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

type ST_openid struct {
	Openid string `json:"openid"`
}

type ST_groups struct {
	Groupsid   int    `json:"id"`
	Groupsname string `json:"name"`
}

type ST_tag struct {
	Tagid   int    `json:"id,omitempty"`
	Tagname string `json:"name,omitempty"`
}

type Wx_User_info struct {
	Errcode        int    `json:"errcode"`
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	Headimgurl     string `json:"headimgurl"`
	Subscribe_time int64  `json:"subscribe_time"`
	Remark         string `json:"remark"`
	Groupid        int    `json:"groupid"`
	Tagid_list     []int  `json:"tagid_list"`
	Unionid        string `json:"unionid"`
}

type Wx_User_list struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	Data        Wx_Openid_list
	Next_openid string `json:"next_openid"`
	Errcode     int    `json:"errcode"`
}

//获取用户user_token
func (wx *Weixin) Get_user_token(Code string) User_token {

	param := make(map[string]string)
	param["appid"] = wx.Wx_appid
	param["secret"] = wx.Wx_appsecret
	param["code"] = Code
	param["grant_type"] = "authorization_code"
	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/sns/oauth2/access_token", param), nil)

	res_body, err := requset(wx, req, 1)
	if err != nil {
		log.Println(err)
	}
	var u_token User_token
	err = json.Unmarshal(res_body, &u_token)
	if err != nil {
		log.Println(err)
	}
	return u_token
}

//获取用户信息by user_token
func (wx *Weixin) Get_wx_user_info(access_token, openid string) (user_info Wx_User_info, err error) {
	param := make(map[string]string)
	param["access_token"] = access_token
	param["openid"] = openid

	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/sns/userinfo?lang=zh_CN", param), nil)
	res_body, err := requset(wx, req, 1)
	if err != nil {
		log.Println(err)
		return user_info, err
	}
	err = json.Unmarshal(res_body, &user_info)
	if err != nil {
		log.Println(err)
		return user_info, err
	}
	return user_info, nil
}

//获取用户信息by ac_token
func (wx *Weixin) Get_wx_user(openid string) (user_info Wx_User_info, err error) {

	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["openid"] = openid
	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/user/info?lang=zh_CN", param), nil)
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return user_info, err
	}
	err = json.Unmarshal(res_body, &user_info)
	if err != nil {
		log.Println(err)
		return user_info, err
	}
	return user_info, nil
}

//批量获取用户信息
func (wx *Weixin) Get_wx_users(openid_s []string) (user_info []Wx_User_info, err error) {
	var openids []ST_openid
	for index, openid := range openid_s {
		if (index+1)%99 == 0 {
			user, err := wx.Get_wx_users_100(openids)
			if err != nil {
				return user, err
			}
			user_info = append(user_info, user...)
			openids = openids[0:0]
		}
		openids = append(openids, ST_openid{openid})
	}
	user, err := wx.Get_wx_users_100(openids)
	if err != nil {
		return user, err
	}
	user_info = append(user_info, user...)
	return
}

func (wx *Weixin) Get_wx_users_100(openids []ST_openid) (user_info []Wx_User_info, err error) {
	if len(openids) > 100 {
		log.Println("数据太多")
		return nil, errors.New("数据太多")
	}
	type t_user_list struct {
		User_list []ST_openid `json:"user_list"`
	}
	t := t_user_list{openids}
	d, err := json.Marshal(t)

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token="+wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return user_info, err
	}

	type users struct {
		User_info_list []Wx_User_info `json:"user_info_list"`
	}
	var user_list users
	temp_data := strings.Replace(string(res_body), "\x1c", "", -1)
	temp_data = strings.Replace(temp_data, "\x0e", "", -1)
	err = json.Unmarshal([]byte(temp_data), &user_list)
	if err != nil {
		log.Println(err, string(res_body))
		//		return user_info, err
	}
	return user_list.User_info_list, nil
}

//获取用户openid列表
func (wx *Weixin) Get_user_list() []string {
	var openid_list []string
	next_openid := ""
loop:
	param := make(map[string]string)
	param["access_token"] = wx.Access_token
	param["next_openid"] = next_openid
	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/user/get", param), nil)
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return openid_list
	}
	var user_list Wx_User_list
	err = json.Unmarshal(res_body, &user_list)
	if err != nil || user_list.Errcode != 0 {
		log.Println(string(res_body))
		return openid_list
	}
	openid_list = append(openid_list, user_list.Data.Openid...)
	if user_list.Total > len(openid_list) {
		next_openid = user_list.Next_openid
		goto loop
	}
	return openid_list
}

//获取所有用户的所有信息
func (wx *Weixin) Get_all_user_info() (user_info []Wx_User_info, err error) {
	return wx.Get_wx_users(wx.Get_user_list())
}

//创建分组
func (wx *Weixin) Create_groups(name string) (groupid int) {

	type groups struct {
		Group ST_groups `json:"group"`
	}
	t := groups{ST_groups{0, name}}
	d, _ := json.Marshal(t)

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/groups/create?access_token="+wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	err = json.Unmarshal(res_body, &t)
	if err != nil {
		log.Println(err)
		return 0
	}
	return t.Group.Groupsid
}

//获取所有分组
func (wx *Weixin) Get_groups() (data []ST_groups, err error) {
	type groups struct {
		Groups []ST_groups `json:"groups"`
	}
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/groups/get?access_token="+wx.Access_token, nil)
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	var t_data groups
	err = json.Unmarshal(res_body, &t_data)
	if err != nil {
		log.Println(err)
		return t_data.Groups, err
	}
	return t_data.Groups, nil
}

//获取用户所在分组
func (wx *Weixin) Get_groupsid(openid string) int {
	type t_str struct {
		Openid string `json:"openid"`
	}
	t := t_str{openid}
	d, _ := json.Marshal(t)

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/groups/getid?access_token="+wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	type t_groupid struct {
		Groupid int `json:"groupid"`
	}
	var t_g t_groupid
	err = json.Unmarshal(res_body, &t_g)
	if err != nil {
		log.Println(err)
		return 0
	}
	return t_g.Groupid
}

//更新分组名
func (wx *Weixin) Update_group(name string, groupid int) int {
	type groups struct {
		Group ST_groups `json:"group"`
	}
	t := groups{ST_groups{groupid, name}}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/groups/update?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//移动用户到分组
func (wx *Weixin) Move_group(openid string, groupid int) int {
	type t_openid struct {
		Openid     string `json:"openid"`
		To_groupid int    `json:"to_groupid"`
	}
	t := t_openid{openid, groupid}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//批量移动用户到分组
func (wx *Weixin) Move_groups(openid []string, groupid int) int {
	type t_openid struct {
		Openid     []string `json:"openid"`
		To_groupid int      `json:"to_groupid"`
	}
	t := t_openid{openid, groupid}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//删除用户分组
func (wx *Weixin) Del_groups(groupid int) int {
	type t_groups struct {
		Group ST_groups `json:"group"`
	}
	t := t_groups{ST_groups{groupid, ""}}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//创建用户标签
func (wx *Weixin) Create_tag(name string) (tagid int) {

	type tag struct {
		Tag ST_tag `json:"tag"`
	}
	t := tag{ST_tag{0, name}}
	d, _ := json.Marshal(t)
	log.Println(string(d))
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/tags/create?access_token="+wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	err = json.Unmarshal(res_body, &t)
	if err != nil {
		log.Println(err)
		return 0
	}
	return t.Tag.Tagid
}

//获取所有用户标签
func (wx *Weixin) Get_tag() (data []ST_tag, err error) {
	type tags struct {
		Tags []ST_tag `json:"tags"`
	}
	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/tags/get?access_token="+wx.Access_token, nil)
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	var t_data tags
	err = json.Unmarshal(res_body, &t_data)
	if err != nil {
		log.Println(err)
		return t_data.Tags, err
	}
	return t_data.Tags, nil
}

//更新用户标签名
func (wx *Weixin) Update_tag(tagid int, name string) int {
	type tags struct {
		Tag ST_tag `json:"tag"`
	}
	t := tags{ST_tag{tagid, name}}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/tags/update?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//删除用户标签
func (wx *Weixin) Del_tags(tagid int) int {
	type t_tags struct {
		Tag ST_tag `json:"tag"`
	}
	t := t_tags{ST_tag{tagid, ""}}
	d, _ := json.Marshal(t)
	log.Println(string(d))
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//获取用户所在标签
func (wx *Weixin) Get_user_tags(openid string) []int {
	type t_str struct {
		Openid string `json:"openid"`
	}
	t := t_str{openid}
	d, _ := json.Marshal(t)

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token="+wx.Access_token, bytes.NewReader(d))
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return nil
	}
	type t_groupid struct {
		Tagid_list []int `json:"tagid_list"`
	}
	var t_g t_groupid
	err = json.Unmarshal(res_body, &t_g)
	if err != nil {
		log.Println(err)
		return nil
	}
	return t_g.Tagid_list
}

//批量为用户打标签
func (wx *Weixin) Batch_tags(openid []string, tag_id int) int {
	type t_openid struct {
		Openid []string `json:"openid_list"`
		Tag_id int      `json:"tagid"`
	}
	t := t_openid{openid, tag_id}
	d, _ := json.Marshal(t)
	log.Println(string(d))
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}

	return 1
}

//批量取消用户标签
func (wx *Weixin) UnBatch_tags(openid []string, tag_id int) int {
	type t_openid struct {
		Openid []string `json:"openid_list"`
		Tag_id int      `json:"tagid"`
	}
	t := t_openid{openid, tag_id}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1
}

//设置备注名
func (wx *Weixin) Update_remark(openid, remark string) int {
	type t_remark struct {
		Openid string `json:"openid"`
		Remark string `json:"remark"`
	}
	t := t_remark{openid, remark}
	d, _ := json.Marshal(t)
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token="+wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}

	return 1
}
