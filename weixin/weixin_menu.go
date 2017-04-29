package weixin

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//按钮元素
type ST_button struct {
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	Key        string      `json:"key,omitempty"`
	Url        string      `json:"url,omitempty"`
	Media_id   string      `json:"media_id,omitempty"`
	Sub_button []ST_button `json:"sub_button,omitempty"`
}

//普通菜单
type ST_Menu struct {
	Button []ST_button `json:"button"`
}

//自定义菜单
type ST_cond_Menu struct {
	Button    []ST_button  `json:"button"`
	Matchrule ST_Matchrule `json:"matchrule"`
	Menuid    int       `json:"menuid,omitempty"`
}

//匹配规则
type ST_Matchrule struct {
	Tag_id               string`json:"tag_id,omitempty"`
	Group_id	string `json:"group_id,omitempty"`
	Sex                  string `json:"sex,omitempty"`
	Client_platform_type string `json:"client_platform_type,omitempty"`
	Country              string `json:"country,omitempty"`
	Province             string `json:"province,omitempty"`
	City                 string `json:"city,omitempty"`
	Language             string `json:"language,omitempty"`
}

//获取菜单
type ST_Menus struct {
	Menu     ST_Menu        `json:"menu"`
	CondMenu []ST_cond_Menu `json:"conditionalmenu"`
}

//获取菜单
func (wx *Weixin) Get_Menu() (data ST_Menus, err error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/menu/get", param), nil)
	if err != nil {
		log.Println(err)
		return data, err
	}
	resp, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, err
}

//创建普通菜单
func (wx *Weixin) Creat_Menu(menu ST_Menu) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	d, _ := json.Marshal(menu)
	d = bytes.Replace(d, []byte("\\u0026"), []byte("&"), -1)
	d = bytes.Replace(d, []byte("\\u003c"), []byte("<"), -1)
	d = bytes.Replace(d, []byte("\\u003e"), []byte(">"), -1)
	d = bytes.Replace(d, []byte("\\u003d"), []byte("="), -1)

	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/menu/create", param),
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

//创建自定义菜单
func (wx *Weixin) Creat_conditional_Menu(menu ST_cond_Menu) (string, error) {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	d, _ := json.Marshal(menu)
	d = bytes.Replace(d, []byte("\\u0026"), []byte("&"), -1)
	d = bytes.Replace(d, []byte("\\u003c"), []byte("<"), -1)
	d = bytes.Replace(d, []byte("\\u003e"), []byte(">"), -1)
	d = bytes.Replace(d, []byte("\\u003d"), []byte("="), -1)
	req, err := http.NewRequest("POST", Param("https://api.weixin.qq.com/cgi-bin/menu/addconditional", param),
		bytes.NewReader(d))
	if err != nil {
		log.Println(err)
		return "", err
	}
	res_body, err := requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return "", err
	}
	type ST_temp struct {
		Menuid int `json:"menuid"`
	}
	var temp ST_temp
	err = json.Unmarshal(res_body, &temp)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return strconv.Itoa(temp.Menuid), nil
}

//删除自定义菜单
func (wx *Weixin) Delete_conditional_Menu(menuid string) int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	type ST_temp struct {
		Menuid string `json:"menuid"`
	}
	temp := ST_temp{menuid}
	d, _ := json.Marshal(temp)
	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/menu/delconditional", param), bytes.NewReader(d))
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

//删除所有菜单
func (wx *Weixin) Delete_All_Menu() int {
	param := make(map[string]string)
	param["access_token"] = wx.Access_token

	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/cgi-bin/menu/delete", param), nil)
	if err != nil {
		log.Println(err)
		return 0
	}
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	//	log.Println(string(resp))
	return 1
}

//https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN
