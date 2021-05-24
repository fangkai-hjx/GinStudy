package controller

import (
	"GinStudy/pkg/result"
	"GinStudy/pkg/validCheck"
	"GinStudy/request"
	"GinStudy/service"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}


func (a *ArticleController) GetOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	articleOne, err := service.GetOneArticle(param.ID)
	if err != nil {
		result.Error(404, "数据查询错误")
	} else {
		result.Success(articleOne)
	}
	return
}
