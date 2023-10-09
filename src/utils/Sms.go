package utils

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

// SendSMS 发送验证码短信
func SendSMS(phone string, code string, tmp int, channel string) error {
	if channel == "1" { // 首选短信通道
		return MessageChannel(phone, code, tmp)
	} else if channel == "2" { // 备用短信通道  未对接
		return MessageChannelSpare(phone, code, tmp)
	} else if channel == "3" { // 次用短信通道  未对接
		return MessageChannelSecond(phone, code, tmp)
	}
	return errors.New("暂无短信通道")
}

// MessageChannel 首选短信通道
func MessageChannel(phone string, code string, tmp int) error {
	/**
	253云通讯】 必须是备案的
	账号：YZM5612013
	状态： 正常
	密码：CkeXzI6DUJf9d8
	接口：http://smssh1.253.com/msg/send/json
	{
	"code":"0",
	"msgId":"17041010383624511",
	"time":"20170410103836",
	"errorMsg":""
	}
	*/
	//templateList := []string{
	//	"【天际】 您的验证码为：" + code + "，非本人操作请忽略本信息。",
	//	"【天际】" + code + " 号，密码输错3次，账户被锁定！",
	//	"【天际】" + code + " 号，输入应急码，账户被锁定！",
	//}
	templateList := []string{
		"【威尼斯】 您的验证码为：" + code + "，非本人操作请忽略本信息。",
		"【威尼斯】" + code + " 号，密码输错3次，账户被锁定！",
		"【威尼斯】" + code + " 号，输入应急码，账户被锁定！",
	}
	if len(templateList) <= tmp {
		return errors.New("模板不存在")
	}
	msg := templateList[tmp]

	message := url.QueryEscape(msg)
	var params string

	params = "&UserId=tj88888&UserPwd=qaztj1238!@A&SendPhone=" + phone + "&SendMessage=" + message + ""
	msgURL := "http://api.eait.cn/Sms/SmsHttp_U.aspx?Action=SendMessage" + params //"http://smssh1.253.com/msg/send/json"

	resp, err := HttpGet(msgURL)
	if err != nil {
		return errors.New("网络错误")
	}

	if resp == "" {
		return errors.New("发送失败")
	}
	if resp[0:1] == "1" {
		return nil
	}
	return errors.New("发送失败")
}

// MessageChannelSpare 备用短信通道
func MessageChannelSpare(phone string, code string, tmp int) error {
	return errors.New("暂无对接备用短信通道")
}

// MessageChannelSecond 次用短信通道
func MessageChannelSecond(phone string, code string, tmp int) error {
	return errors.New("暂无对接次用短信通道")
}

// SendEmail 发送邮件
func SendEmail(email, title, content string) error {
	accesskey := "5ca830e287b65f6374f3e1f3"
	secretkey := "eddf6f346eeb4ece92f7b7cf3ab99583"
	randInt := RandInt64(100000, 999999)
	randStr := strconv.Itoa(int(randInt))
	url := "https://live.kewail.com/directmail/v1/singleSendMail?accesskey=" +
		accesskey + "&random=" + randStr
	fromEmail := "mail@service.mnbkop.com"
	timeStamp := time.Now().Unix()
	timeStampStr := strconv.Itoa(int(timeStamp))

	data := "secretkey=" + secretkey + "&random=" + randStr + "&time=" +
		timeStampStr + "&fromEmail=" + fromEmail

	postMap := map[string]interface{}{
		"ext":         "",
		"replyEmail":  fromEmail,
		"fromAlias":   "验证码服务",
		"htmlBody":    content,
		"needToReply": true,
		"subject":     title,
		"clickTrace":  "0",
		"time":        timeStamp,
		"type":        0,
		"toEmail":     email,
		"fromEmail":   fromEmail,
		"sig":         SHA256(data),
	}
	postStr, _ := json.Marshal(postMap)
	resp, err := HttpPost(url, postStr, 0)
	if err != nil {
		return err
	}

	var result = struct {
		Result int    `json:"result"`
		ErrMsg string `json:"errmsg"`
	}{}
	_ = json.Unmarshal([]byte(resp), &result)
	if result.Result == 0 && result.ErrMsg == "OK" {
		return nil
	}

	return errors.New("发送失败")
}
