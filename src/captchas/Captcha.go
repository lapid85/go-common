package captchas

import (
	"strings"

	"github.com/mojocn/base64Captcha"
)

// CaptchaSource 默认生成图片Source
var CaptchaSource = "0123456789"

// Captcha 生成验证码
type Captcha struct {
	Driver base64Captcha.Driver
	Store  base64Captcha.Store
}

// Capt 生成默认验主码
var Capt = func(platform string) *Captcha {
	return DefaultCaptcha(platform)
}

// GenerateCaptcha 生成验证码
func (ths *Captcha) GenerateCaptcha() (map[string]interface{}, error) {
	c := base64Captcha.NewCaptcha(ths.Driver, ths.Store)
	id, b64s, err := c.Generate()
	if err != nil {
		return nil, err
	}
	body := map[string]interface{}{"captcha": b64s, "captchaId": id}
	return body, nil
}

// Verify 校验验证码
func (ths *Captcha) Verify(id, verifyValue string) bool {
	verifyValue = strings.ToLower(verifyValue)
	return ths.Store.Verify(id, verifyValue, true)
}
