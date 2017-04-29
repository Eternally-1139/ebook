package weixin

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type hb_give struct {
	XMLName      xml.Name `xml:"xml"`
	Nonce_str    string   `xml:"nonce_str"`
	Sign         string   `xml:"sign"`
	Mch_billno   string   `xml:"mch_billno"`
	Mch_id       string   `xml:"mch_id"`
	Wxappid      string   `xml:"wxappid"`
	Send_name    string   `xml:"send_name"`
	Re_openid    string   `xml:"re_openid"`
	Total_amount int      `xml:"total_amount"`
	Total_num    int      `xml:"total_num"`
	Wishing      string   `xml:"wishing"`
	Client_ip    string   `xml:"client_ip"`
	Act_name     string   `xml:"act_name"`
	Remark       string   `xml:"remark"`
}

type ST_req_order struct {
	XMLName        xml.Name `xml:"xml"`
	Appid          string   `xml:"appid"`
	Mch_id         string   `xml:"mch_id"`
	Transaction_id string   `xml:"transaction_id,omitempty"`
	Out_trade_no   string   `xml:"out_trade_no,omitempty"`
	Nonce_str      string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
}

type Sendredpack_response struct {
	XMLName      xml.Name `xml:"xml"`
	Return_code  string   `xml:"return_code"`
	Return_msg   string   `xml:"return_msg"`
	Result_code  string   `xml:"result_code"`
	Err_code     string   `xml:"err_code"`
	Err_code_des string   `xml:"err_code_des"`
	Mch_billno   string   `xml:"mch_billno"`
	Mch_id       string   `xml:"mch_id"`
	Wxappid      string   `xml:"wxappid"`
	Re_openid    string   `xml:"re_openid"`
	Send_listid  string   `xml:"send_listid"`
}

type ST_Unifiedorder struct {
	XMLName          xml.Name `xml:"xml"`
	Appid            string   `xml:"appid"`
	Mch_id           string   `xml:"mch_id"`
	Device_info      string   `xml:"device_info,omitempty"`
	Nonce_str        string   `xml:"nonce_str"`
	Sign             string   `xml:"sign"`
	Body             string   `xml:"body"`
	Detail           string   `xml:"detail,omitempty"`
	Attach           string   `xml:"attach,omitempty"`
	Out_trade_no     string   `xml:"out_trade_no"`
	Fee_type         string   `xml:"fee_type,omitempty"`
	Total_fee        int      `xml:"total_fee"`
	Spbill_create_ip string   `xml:"spbill_create_ip"`
	Time_start       string   `xml:"time_start,omitempty"`
	Time_expire      string   `xml:"time_expire,omitempty"`
	Goods_tag        string   `xml:"goods_tag,omitempty"`
	Notify_url       string   `xml:"notify_url"`
	Trade_type       string   `xml:"trade_type"`
	Product_id       string   `xml:"product_id,omitempty"`
	Limit_pay        string   `xml:"limit_pay,omitempty"`
	Openid           string   `xml:"openid"`
}

type ST_res_Unified struct {
	XMLName      xml.Name `xml:"xml"`
	Return_code  string   `xml:"return_code"`
	Return_msg   string   `xml:"return_msg"`
	Appid        string   `xml:"appid"`
	Mch_id       string   `xml:"mch_id"`
	Device_info  string   `xml:"device_info"`
	Nonce_str    string   `xml:"nonce_str"`
	Sign         string   `xml:"sign"`
	Result_code  string   `xml:"result_code"`
	Err_code     string   `xml:"err_code"`
	Err_code_des string   `xml:"err_code_des"`
	Trade_type   string   `xml:"trade_type"`
	Prepay_id    string   `xml:"prepay_id"`
	Code_url     string   `xml:"code_url"`
}

type ST_Orderquery struct {
	Return_code          string `xml:"return_code"`
	Return_msg           string `xml:"return_msg"`
	Appid                string `xml:"appid"`
	Mch_id               string `xml:"mch_id"`
	Device_info          string `xml:"device_info"`
	Nonce_str            string `xml:"nonce_str"`
	Sign                 string `xml:"sign"`
	Result_code          string `xml:"result_code"`
	Err_code             string `xml:"err_code"`
	Err_code_des         string `xml:"err_code_des"`
	Openid               string `xml:"openid"`
	Is_subscribe         string `xml:"is_subscribe"`
	Trade_type           string `xml:"trade_type"`
	Bank_type            string `xml:"bank_type"`
	Total_fee            string `xml:"total_fee"`
	Settlement_total_fee int    `xml:"settlement_total_fee"`
	Fee_type             string `xml:"fee_type"`
	Cash_fee             int    `xml:"cash_fee"`
	Cash_fee_type        string `xml:"cash_fee_type"`
	Coupon_fee           int    `xml:"coupon_fee"`
	Coupon_count         int    `xml:"coupon_count"`
	Transaction_id       string `xml:"transaction_id"`
	Out_trade_no         string `xml:"out_trade_no"`
	Attach               string `xml:"attach"`
	Time_end             string `xml:"time_end"`
	Trade_state          string `xml:"trade_state"`
	Trade_state_desc     string `xml:"trade_state_desc"`
}

type ST_Close_order struct {
	Return_code  string `xml:"return_code"`
	Return_msg   string `xml:"return_msg"`
	Appid        string `xml:"appid"`
	Mch_id       string `xml:"mch_id"`
	Nonce_str    string `xml:"nonce_str"`
	Sign         string `xml:"sign"`
	Err_code     string `xml:"err_code"`
	Err_code_des string `xml:"err_code_des"`
}

//支付下单https://api.mch.weixin.qq.com/pay/unifiedorder
func (wx *Weixin) Unifiedorder(device_info, body, detail, attach, out_trade_no, fee_type string,
	total_fee int, spbill_create_ip, time_start, time_expire, goods_tag, notify_url,
	trade_type, product_id, limit_pay, openid string) (data ST_res_Unified, err error) {

	order := ST_Unifiedorder{xml.Name{}, wx.Wx_appid, wx.Option["mch_id"], device_info,
		Random_str(20, 3), "", body, detail, attach, out_trade_no, fee_type,
		total_fee, spbill_create_ip, time_start, time_expire, goods_tag,
		notify_url, trade_type, product_id, limit_pay, openid}

	order.Sign = Sign_v_md5(order, wx.Option["pay_key"])

	d, _ := xml.Marshal(order)

	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder",
		bytes.NewReader(d))

	res_body, err := requset(wx, req, -1)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = xml.Unmarshal(res_body, &data)
	if err != nil {
		return data, err
	}
	Sign := data.Sign
	data.Sign = ""
	if Sign_v_md5(data, wx.Option["pay_key"]) != Sign {
		return data, errors.New("签名错误")
	}
	return data, nil

}

//查询订单https://api.mch.weixin.qq.com/pay/orderquery
func (wx *Weixin) Orderquery(transaction_id, out_trade_no string) (data ST_Orderquery, err error) {

	query_order := ST_req_order{xml.Name{}, wx.Wx_appid, wx.Option["mch_id"],
		transaction_id, out_trade_no, Random_str(20, 3), ""}

	query_order.Sign = Sign_v_md5(query_order, wx.Option["pay_key"])

	d, _ := xml.Marshal(query_order)

	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/orderquery",
		bytes.NewReader(d))

	res_body, err := requset(wx, req, -1)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = xml.Unmarshal(res_body, &data)
	return data, err
}

//关闭订单https://api.mch.weixin.qq.com/pay/closeorder
func (wx *Weixin) Close_order(out_trade_no string) (data ST_Close_order, err error) {

	query_order := ST_req_order{xml.Name{}, wx.Wx_appid, wx.Option["mch_id"],
		"", out_trade_no, Random_str(20, 3), ""}

	query_order.Sign = Sign_v_md5(query_order, wx.Option["pay_key"])

	d, _ := xml.Marshal(query_order)

	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/closeorder",
		bytes.NewReader(d))

	res_body, err := requset(wx, req, -1)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = xml.Unmarshal(res_body, &data)
	return data, err
}

//申请退款https://api.mch.weixin.qq.com/secapi/pay/refund

//下载对账单https://api.mch.weixin.qq.com/pay/downloadbill

func (wx *Weixin) getTLSConfig() (*tls.Config, error) {
	if wx._tlsConfig != nil {
		return wx._tlsConfig, nil
	}
	cert, err := tls.LoadX509KeyPair(wx.Option["CertPath"], wx.Option["KeyPath"])
	if err != nil {
		return nil, err
	}

	caData, err := ioutil.ReadFile(wx.Option["CaPath"])
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	wx._tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return wx._tlsConfig, nil
}

func (wx *Weixin) HttpsPost(url string, xmlContent []byte, Content_Type string) (*http.Response, error) {
	tlsConfig, err := wx.getTLSConfig()
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(url,
		Content_Type,
		bytes.NewBuffer(xmlContent))
}

//发送红包
func (wx *Weixin) Send_hongbao(mch_billno, send_name, re_openid string, total_amount, total_num int,
	wishing, client_ip, act_name, remark string) (res_st Sendredpack_response, err error) {
	st_hb := hb_give{
		xml.Name{},
		Random_str(32, 3),
		"",
		mch_billno,
		wx.Option["mch_id"],
		wx.Wx_appid,
		send_name,
		re_openid,
		total_amount,
		total_num,
		wishing,
		client_ip,
		act_name,
		remark}

	st_hb.Sign = Sign_v_md5(st_hb, wx.Option["pay_key"])
	data, err := xml.MarshalIndent(&st_hb, "", " ")
	if err != nil {
		return res_st, err
	}
	res, err := wx.HttpsPost("https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", data, "text/xml")
	if err != nil {
		log.Println(res, err)
		return res_st, err
	}
	res_body := make([]byte, 1024)
	res.Body.Read(res_body)
	err = xml.Unmarshal(res_body, &res_st)
	if err != nil {
		return res_st, err
	}
	return res_st, nil
}
