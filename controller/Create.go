package controller

import (
	"common/model"

	"github.com/gin-gonic/gin"
)

// ActionCreate 创建
type ActionCreate struct {
	Model model.IModel
}

// Create 创建 - post
func (ths *ActionCreate) Create(c *gin.Context) {

}
