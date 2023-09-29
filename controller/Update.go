package controller

import (
	"common/model"

	"github.com/gin-gonic/gin"
)

// ActionUpdate 更新
type ActionUpdate struct {
	Model model.IModel
}

// Update 更新 - put
func (ths *ActionUpdate) Update(c *gin.Context) {

}
