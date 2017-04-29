package weixin

import (
	"net/http"
	"log"
	"encoding/json"
)

type SessionKey struct {
	Openid string `json:"openid"`
	Key string `json:"session_key"`
	Errcode int `json:"errcode"`
	Errmsg int `json:"errmsg"`
}

//小程序获取session_key
func (wx *Weixin) GetSessionKey(code string)(session_key SessionKey,err error){
	param := make(map[string]string)
	param["appid"] = wx.Wx_appid
	param["secret"] = wx.Wx_appsecret
	param["js_code"]=code
	param["grant_type"]="authorization_code"

	req, err := http.NewRequest("GET", Param("https://api.weixin.qq.com/sns/jscode2session", param), nil)
	res_body, err := requset(wx, req, 1)
	if err != nil {
		log.Println(err)
		return session_key, err
	}
	err = json.Unmarshal(res_body, &session_key)
	if err != nil {
		log.Println(err)
		return session_key, err
	}
	return session_key, nil
}
