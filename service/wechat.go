package service

import (
	"net/http"
	"strings"
	"bytes"
	"fmt"
	"shopktl/models"
	"sort"
	"strconv"
	"crypto/sha1"
	"io"
	"reflect"
	"crypto/md5"
	"time"
	"math/rand"

	"io/ioutil"
	"encoding/json"
)

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}
type JsApiTicketResponse struct {
	JsApiTicket string `json:"ticket"`
	ExpiresIn   float64 `json:"expires_in"`
}

func  GetAccessToken(appid string,secret string) string{

	wx:=models.Weixin{Id:1}
	wx.Read()
	if (time.Now().Unix()-wx.GetAcTime.Unix())<6000{
		return wx.AccessToken
	}else{
		requestLine:="https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+appid+"&secret="+secret
		resp, err := http.Get(requestLine)

		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Println("发送get请求获取 openid 错误", err)
			return ""
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("发送get请求获取 atoken 读取返回body错误", err)
			return ""
		}

		if bytes.Contains(body, []byte("access_token")) {
			atr := AccessTokenResponse{}
			err = json.Unmarshal(body, &atr)
			if err != nil {
				fmt.Println("发送get请求获取 atoken 返回数据json解析错误", err)
				return ""
			}
			fmt.Println("AccessToken:",atr.AccessToken)


			wx.AccessToken=atr.AccessToken
			wx.ExpiresIn=atr.ExpiresIn
			wx.GetAcTime=time.Now()
			wx.Update()
			//createWxMenu(atr.AccessToken)
			defer resp.Body.Close()
			return atr.AccessToken

		} else {

		}

		return ""
	}


}
func PushWxMenuCreate(accessToken string, menuJsonBytes []byte) error {

	postReq, err := http.NewRequest("POST",
		strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/menu/create", "?access_token=", accessToken}, ""),
		bytes.NewReader(menuJsonBytes))

	if err != nil {
		fmt.Println("向微信发送菜单建立请求失败", err)
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("client向微信发送菜单建立请求失败", err)
		return err
	} else {
		fmt.Println("向微信发送菜单建立成功")
	}
	defer resp.Body.Close()

	return nil
}

func  createWxMenu(accessToken string) {


	menuStr := `{
            "button": [
            {
                "name": "合取官网",
                "type": "view",
                "url": "http://www.hnhqjk.com"
            },
            {

                "name":"公司名医",
                "type": "view",
                "url": "http://www.hnhqjk.com/small"
            },
            {
                "name": "合取动态",
                "type": "view",
                "url": "http://www.hnhqjk.com/dynamic"
            }
            ]
        }`

	//发送建立菜单的post请求
	PushWxMenuCreate(accessToken, []byte(menuStr))

}

func GetTicket(accessToken string)string{
	wx:=models.Weixin{Id:1}
	wx.Read()
	if (time.Now().Unix()-wx.GetTiTime.Unix())<6000{
		return wx.JsApiTicket
	}else{

		requestLine:="https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token="+accessToken+"&type=jsapi"
		resp, err := http.Get(requestLine)

		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Println("发送get请求获取 签名 错误", err)
			return ""
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("发送get请求获取 ticket 读取返回body错误", err)
			return ""
		}

		if bytes.Contains(body, []byte("ticket")) {
			tickstr := JsApiTicketResponse{}
			err = json.Unmarshal(body, &tickstr)
			if err != nil {
				fmt.Println("发送get请求获取 atoken 返回数据json解析错误", err)
				return ""
			}
			fmt.Println("JsApiTicket:",tickstr.JsApiTicket)

			wx.JsApiTicket=tickstr.JsApiTicket
			wx.GetTiTime=time.Now()
			wx.Update()
			defer resp.Body.Close()
			return tickstr.JsApiTicket

		} else {

		}

		return ""

	}
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

func Create_signature( Jsapi_ticket string, noncestr string, timestamp string, url string) string {
	var str1 string

	str1 = "jsapi_ticket=" + Jsapi_ticket + "&noncestr=" + noncestr + "&timestamp=" + timestamp + "&url=" + url
	return str2sha1(str1)
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
