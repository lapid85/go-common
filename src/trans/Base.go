package trans

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

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
