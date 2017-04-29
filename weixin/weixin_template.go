package weixin

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

//设置所属行业https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN
//获取设置的行业信息https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN
//获得模板IDhttps://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN
//获取模板列表https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN
//删除模板https://api,weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN
//

//      {
//           "touser":"OPENID",
//           "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
//           "url":"http://weixin.qq.com/download",
//           "data":{
//                   "first": {
//                       "value":"恭喜你购买成功！",
//                       "color":"#173177"
//                   },
//                   "keynote1":{
//                       "value":"巧克力",
//                       "color":"#173177"
//                   },
//                   "keynote2": {
//                       "value":"39.8元",
//                       "color":"#173177"
//                   },
//                   "keynote3": {
//                       "value":"2014年9月22日",
//                       "color":"#173177"
//                   },
//                   "remark":{
//                       "value":"欢迎再次购买！",
//                       "color":"#173177"
//                   }
//           }
//       }

type ST_template_data struct {
	Value string `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

type ST_template struct {
	Touser      string      `json:"touser"`
	Template_id string      `json:"template_id"`
	Url         string      `json:"url"`
	Data        interface{} `json:"data,json"`
}

//发送模板消息https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
func (wx *Weixin) Send_template(touser, template_id, url string, data map[string]ST_template_data) int {

	t := ST_template{touser, template_id, url, data}
	d, _ := json.Marshal(t)

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+
		wx.Access_token, bytes.NewReader(d))
	_, err = requset(wx, req, 0)
	if err != nil {
		log.Println(err)
		return 0
	}
	return 1

}
