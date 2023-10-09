package trans

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

/*
简体中文(中国)	zh-cn	繁体中文(台湾地区)	zh-tw
繁体中文(香港)	zh-hk	英语(香港)	en-hk
英语(美国)	en-us	英语(英国)	en-gb
英语(全球)	en-ww	英语(加拿大)	en-ca
英语(澳大利亚)	en-au	英语(爱尔兰)	en-ie
英语(芬兰)	en-fi	芬兰语(芬兰)	fi-fi
英语(丹麦)	en-dk	丹麦语(丹麦)	da-dk
英语(以色列)	en-il	希伯来语(以色列)	he-il
英语(南非)	en-za	英语(印度)	en-in
英语(挪威)	en-no	英语(新加坡)	en-sg
英语(新西兰)	en-nz	英语(印度尼西亚)	en-id
英语(菲律宾)	en-ph	英语(泰国)	en-th
英语(马来西亚)	en-my	英语(阿拉伯)	en-xa
韩文(韩国)	ko-kr	日语(日本)	ja-jp
荷兰语(荷兰)	nl-nl	荷兰语(比利时)	nl-be
葡萄牙语(葡萄牙)	pt-pt	葡萄牙语(巴西)	pt-br
法语(法国)	fr-fr	法语(卢森堡)	fr-lu
法语(瑞士)	fr-ch	法语(比利时)	fr-be
法语(加拿大)	fr-ca	西班牙语(拉丁美洲)	es-la
西班牙语(西班牙)	es-es	西班牙语(阿根廷)	es-ar
西班牙语(美国)	es-us	西班牙语(墨西哥)	es-mx
西班牙语(哥伦比亚)	es-co	西班牙语(波多黎各)	es-pr
德语(德国)	de-de	德语(奥地利)	de-at
德语(瑞士)	de-ch	俄语(俄罗斯)	ru-ru
意大利语(意大利)	it-it	希腊语(希腊)	el-gr
挪威语(挪威)	no-no	匈牙利语(匈牙利)	hu-hu
土耳其语(土耳其)	tr-tr	捷克语(捷克共和国)	cs-cz
斯洛文尼亚语	sl-sl	波兰语(波兰)	pl-pl
瑞典语(瑞典)	sv-se	西班牙语 (智利)	es-cl

go i18n 支持的语言代码
English: en
Spanish: es
French: fr
German: de
Italian: it
Portuguese: pt
Chinese (Simplified): zh-Hans
Chinese (Traditional): zh-Hant
Japanese: ja
Korean: ko
Russian: ru
Arabic: ar
Hindi: hi
Dutch: nl
Swedish: sv
Danish: da
Norwegian: no
Finnish: fi
Greek: el
Turkish: tr
Hebrew: he
Thai: th
Vietnamese: vi
Indonesian: id
Malay: ms
*/

// Tr 国际化
func Tr(c *gin.Context, transKey string, args ...map[string]string) string {
	data := map[string]string{}
	if len(args) > 0 {
		data = args[0]
	}
	return ginI18n.MustGetMessage(c, &i18n.LocalizeConfig{
		MessageID:    transKey,
		TemplateData: data,
	})
}
